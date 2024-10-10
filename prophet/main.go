package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"

	"github.com/warden-protocol/wardenprotocol/go-client"
	"github.com/warden-protocol/wardenprotocol/prophet/internal/comet"
	"github.com/warden-protocol/wardenprotocol/prophet/internal/egress"
	"github.com/warden-protocol/wardenprotocol/prophet/internal/exec"
	"github.com/warden-protocol/wardenprotocol/prophet/internal/http"
	"github.com/warden-protocol/wardenprotocol/prophet/internal/ingress"

	_ "github.com/warden-protocol/wardenprotocol/prophet/internal/futures/echo"
	_ "github.com/warden-protocol/wardenprotocol/prophet/internal/futures/wardenai"
)

var (
	tendermintRpc = "http://127.0.0.1:26657"
	listenAddr    = "127.0.0.1:8088"
	wardenGrpc    = "127.0.0.1:9090"
)

func main() {
	initGlobalLogger()

	cmt, err := comet.New(tendermintRpc)
	if err != nil {
		panic(err)
	}

	valAddr, err := cmt.ValidatorAddress()
	if err != nil {
		panic(err)
	}
	log.Printf("connected to validator %s", valAddr)

	src := simpleFuturesSource{}
	proposalSrc := simpleProposalSource{
		ValidatorAddress: valAddr,
	}
	memorySink := egress.NewMemorySink()
	memoryProposalSink := egress.NewVoteMemorySink()

	if err := exec.Futures(src, memorySink); err != nil {
		panic(err)
	}

	if err := exec.Votes(proposalSrc, memoryProposalSink); err != nil {
		panic(err)
	}

	log.Println("serving on", listenAddr)
	server := http.NewServer(listenAddr, memorySink, memoryProposalSink)
	if err := server.Serve(); err != nil {
		panic(err)
	}
}

type simpleFuturesSource struct{}

func (f simpleFuturesSource) Fetch() <-chan ingress.Future {
	c, err := client.NewQueryClient(wardenGrpc, true)
	if err != nil {
		panic(err)
	}

	ch := make(chan ingress.Future)
	go func() {
		for {
			futures, err := c.PendingFutures(context.TODO(), &client.PageRequest{
				Limit: 10,
			})
			if err != nil {
				log.Printf("error fetching futures: %v", err)
				time.Sleep(6 * time.Second)
				continue
			}

			for _, f := range futures {
				ch <- ingress.Future{
					ID:      f.Id,
					Handler: f.Handler,
					Input:   f.Input,
				}
			}

			time.Sleep(6 * time.Second)
		}
	}()
	return ch
}

type simpleProposalSource struct {
	ValidatorAddress string
}

func (f simpleProposalSource) Fetch() <-chan ingress.FutureResult {
	c, err := client.NewQueryClient(wardenGrpc, true)
	if err != nil {
		panic(err)
	}

	ch := make(chan ingress.FutureResult)
	go func() {
		for {
			futures, err := c.FuturesPendingVote(context.TODO(), f.ValidatorAddress, &client.PageRequest{
				Limit: 10,
			})
			if err != nil {
				log.Printf("error fetching futures pending vote: %v", err)
				time.Sleep(6 * time.Second)
				continue
			}

			for _, f := range futures {
				ch <- ingress.FutureResult{
					Future: ingress.Future{
						ID:      f.Future.Id,
						Handler: f.Future.Handler,
						Input:   f.Future.Input,
					},
					Output: f.Result.Output,
				}
			}

			time.Sleep(6 * time.Second)
		}
	}()
	return ch
}

func initGlobalLogger() {
	slog.SetDefault(slog.New(
		tint.NewHandler(os.Stderr, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
		}),
	))
}

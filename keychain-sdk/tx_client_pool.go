package keychain

import (
	"context"
	"fmt"
	"log/slog"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	client "github.com/warden-protocol/wardenprotocol/go-client"
	"google.golang.org/grpc/connectivity"
)

type AppClient struct {
	grpcUrl      string
	grpcInsecure bool
	query        *client.QueryClient
	txClient     *client.TxClient
}

type ClientsPool struct {
	clients []*AppClient
	config  Config
}

func NewClientsPool(config Config) *ClientsPool {
	pool := ClientsPool{
		clients: make([]*AppClient, 0),
		config:  config,
	}

	return &pool
}

func (a *ClientsPool) initConnections(logger *slog.Logger) error {
	identity, err := client.NewIdentityFromSeed(a.config.Mnemonic)
	if err != nil {
		return fmt.Errorf("failed to create identity: %w", err)
	}

	for _, grpcUrl := range a.config.GRPCConfigs {
		appClient, err := a.initConnection(logger, grpcUrl, a.config.BasicConfig, identity)
		if err != nil {
			return err
		}

		a.clients = append(a.clients, appClient)
	}

	return nil
}

func (a *ClientsPool) initConnection(
	logger *slog.Logger,
	grpcNodeConfig GrpcNodeConfig,
	config BasicConfig,
	identity client.Identity) (*AppClient, error) {
	appClient := &AppClient{
		grpcUrl:      grpcNodeConfig.GRPCURL,
		grpcInsecure: grpcNodeConfig.GRPCInsecure,
	}

	logger.Info("connecting to Warden Protocol using gRPC", "url", grpcNodeConfig.GRPCURL, "insecure", grpcNodeConfig.GRPCInsecure)

	query, err := client.NewQueryClient(grpcNodeConfig.GRPCURL, grpcNodeConfig.GRPCInsecure)
	if err != nil {
		return nil, fmt.Errorf("failed to create query client: %w", err)
	}
	appClient.query = query

	conn := query.Conn()

	logger.Info("keychain writer identity", "address", identity.Address.String())

	appClient.txClient = client.NewTxClient(identity, config.ChainID, conn, query)

	return appClient, nil
}

func (a *ClientsPool) liveTxClient() (*client.TxClient, error) {
	for _, appClient := range a.clients {
		if state := appClient.query.Conn().GetState(); isOnline(state) {
			return appClient.txClient, nil
		}
	}

	return nil, fmt.Errorf("all node clients are down")
}

func isOnline(state connectivity.State) bool {
	return state == connectivity.Ready || state == connectivity.Idle
}

func (p *ClientsPool) BuildTx(
	ctx context.Context,
	gasLimit uint64,
	fees sdkTypes.Coins,
	msgers ...client.Msger) ([]byte, error) {
	liveClient, err := p.liveTxClient()
	if err != nil {
		return nil, fmt.Errorf("failed to aquire live client for BuildTx: %w", err)
	}

	return liveClient.BuildTx(ctx, gasLimit, fees, msgers...)
}

func (p *ClientsPool) SendWaitTx(ctx context.Context, txBytes []byte) error {
	liveClient, err := p.liveTxClient()
	if err != nil {
		return fmt.Errorf("failed to aquire live client for SendWaitTx: %w", err)
	}

	return liveClient.SendWaitTx(ctx, txBytes)
}

func (a *ClientsPool) ConnectionState() map[string]connectivity.State {
	statuses := make(map[string]connectivity.State)

	for _, appClient := range a.clients {
		state := appClient.query.Conn().GetState()
		statuses[appClient.grpcUrl] = state
	}

	return statuses
}
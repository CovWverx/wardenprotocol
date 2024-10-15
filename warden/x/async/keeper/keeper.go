package keeper

import (
	"context"
	"fmt"
	"time"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	types "github.com/warden-protocol/wardenprotocol/warden/x/async/types/v1beta1"
)

type (
	Keeper struct {
		cdc          codec.Codec
		storeService store.KVStoreService
		logger       log.Logger

		// the address capable of executing a MsgUpdateParams message. Typically, this
		// should be the x/gov module account.
		authority string

		futures *FutureKeeper
		votes   collections.Map[collections.Pair[uint64, []byte], int32]

		FuturesSource keeperFutureSource
		ResultsSource keeperResultsSource
	}
)

var (
	FutureSeqPrefix       = collections.NewPrefix(0)
	FuturesPrefix         = collections.NewPrefix(1)
	FutureByAddressPrefix = collections.NewPrefix(2)
	ResultsPrefix         = collections.NewPrefix(3)
	VotesPrefix           = collections.NewPrefix(4)
)

func NewKeeper(
	cdc codec.Codec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,
	selfValAddr sdk.ConsAddress,
) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	sb := collections.NewSchemaBuilder(storeService)

	futures := NewFutureKeeper(sb, cdc)

	votes := collections.NewMap(
		sb,
		VotesPrefix,
		"votes",
		collections.PairKeyCodec(collections.Uint64Key, collections.BytesKey),
		collections.Int32Value,
	)

	_, err := sb.Build()
	if err != nil {
		panic(fmt.Sprintf("failed to build schema: %s", err))
	}

	return Keeper{
		cdc:          cdc,
		storeService: storeService,
		authority:    authority,
		logger:       logger,

		futures: futures,
		votes:   votes,

		FuturesSource: NewFuturesSource(),
		ResultsSource: NewResultsSource(selfValAddr),
	}
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) getBlockTime(ctx context.Context) time.Time {
	return sdk.UnwrapSDKContext(ctx).HeaderInfo().Time
}

func (k Keeper) AddFutureResult(ctx context.Context, id uint64, submitter, output []byte) error {
	if err := k.futures.SetResult(ctx, types.FutureResult{
		Id:        id,
		Output:    output,
		Submitter: submitter,
	}); err != nil {
		return err
	}

	if err := k.SetFutureVote(ctx, id, submitter, types.FutureVoteType_VOTE_TYPE_VERIFIED); err != nil {
		return err
	}

	return nil
}

func (k Keeper) SetFutureVote(ctx context.Context, id uint64, voter []byte, vote types.FutureVoteType) error {
	return k.votes.Set(ctx, collections.Join(id, voter), int32(vote))
}

func (k Keeper) GetFutureVotes(ctx context.Context, futureId uint64) ([]types.FutureVote, error) {
	it, err := k.votes.Iterate(ctx, collections.NewPrefixedPairRange[uint64, []byte](futureId))
	if err != nil {
		return nil, err
	}
	defer it.Close()

	var votes []types.FutureVote
	for ; it.Valid(); it.Next() {
		k, err := it.Key()
		if err != nil {
			return nil, err
		}
		vote, err := it.Value()
		if err != nil {
			return nil, err
		}
		votes = append(votes, types.FutureVote{
			FutureId: futureId,
			Voter:    k.K2(),
			Vote:     types.FutureVoteType(vote),
		})
	}

	return votes, nil
}

package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) EndBlocker(ctx context.Context) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	blockHeight := sdkCtx.BlockHeight()
	params := k.GetParams(ctx)

	latestPruneHeight, err := k.ActionKeeper.GetLatestPruneHeight(ctx)
	if err != nil {
		return err
	}

	if (blockHeight - latestPruneHeight) < params.PruneCheckBlockFrequency {
		return nil
	}

	blockTime := k.getBlockTime(ctx)
	expiredActions, err := k.ActionKeeper.ExpiredActions(ctx, blockTime, params.MaxPendingTime, params.MaxCompletedTime)
	if err != nil {
		return err
	}

	for _, act := range expiredActions {
		if err := k.ActionKeeper.pruneAction(ctx, act, blockHeight); err != nil {
			return err
		}
	}

	return nil
}
package crisis

import (
	"context"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	"github.com/cosmos/cosmos-sdk/x/crisis/types"
)

// EndBlocker checks all registered invariants
func EndBlocker(ctx context.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, telemetry.Now(), telemetry.MetricKeyEndBlocker)

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if k.InvCheckPeriod() == 0 || sdkCtx.BlockHeight()%int64(k.InvCheckPeriod()) != 0 {
		// skip running the invariant check
		return
	}
	k.AssertInvariants(sdkCtx)
}

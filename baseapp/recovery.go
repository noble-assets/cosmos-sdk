package baseapp

import (
	"fmt"
	"runtime/debug"

	errorsmod "cosmossdk.io/errors"
	storetypes "cosmossdk.io/store/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// RecoveryHandler handles recovery() object.
// Return a non-nil error if recoveryObj was processed.
// Return nil if recoveryObj was not processed.
type RecoveryHandler func(recoveryObj any) error

// recoveryMiddleware is wrapper for RecoveryHandler to create chained recovery handling.
// returns (recoveryMiddleware, nil) if recoveryObj was not processed and should be passed to the next middleware in chain.
// returns (nil, error) if recoveryObj was processed and middleware chain processing should be stopped.
type recoveryMiddleware func(recoveryObj any) (recoveryMiddleware, error)

// processRecovery processes recoveryMiddleware chain for recovery() object.
// Chain processing stops on non-nil error or when chain is processed.
func processRecovery(recoveryObj any, middleware recoveryMiddleware) error {
	if middleware == nil {
		return nil
	}

	next, err := middleware(recoveryObj)
	if err != nil {
		return err
	}

	return processRecovery(recoveryObj, next)
}

// newRecoveryMiddleware creates a RecoveryHandler middleware.
func newRecoveryMiddleware(handler RecoveryHandler, next recoveryMiddleware) recoveryMiddleware {
	return func(recoveryObj any) (recoveryMiddleware, error) {
		if err := handler(recoveryObj); err != nil {
			return nil, err
		}

		return next, nil
	}
}

// newOutOfGasRecoveryMiddleware creates a standard OutOfGas recovery middleware for app.runTx method.
func newOutOfGasRecoveryMiddleware(gasWanted uint64, ctx sdk.Context, next recoveryMiddleware) recoveryMiddleware {
	handler := func(recoveryObj any) error {
		err, ok := recoveryObj.(storetypes.ErrorOutOfGas)
		if !ok {
			return nil
		}

		return errorsmod.Wrap(
			sdkerrors.ErrOutOfGas, fmt.Sprintf(
				"out of gas in location: %v; gasWanted: %d, gasUsed: %d",
				err.Descriptor, gasWanted, ctx.GasMeter().GasConsumed(),
			),
		)
	}

	return newRecoveryMiddleware(handler, next)
}

// newDefaultRecoveryMiddleware creates a default (last in chain) recovery middleware for app.runTx method.
func newDefaultRecoveryMiddleware() recoveryMiddleware {
	handler := func(recoveryObj any) error {
		return errorsmod.Wrap(
			sdkerrors.ErrPanic, fmt.Sprintf(
				"recovered: %v\nstack:\n%v", recoveryObj, string(debug.Stack()),
			),
		)
	}

	return newRecoveryMiddleware(handler, nil)
}

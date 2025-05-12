package types

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type SendActionFn func(ctx context.Context, fromAddr, toAddr sdk.AccAddress, amt sdk.Coins) error

func (SendActionFn) IsOnePerModuleType() {}

func (p SendActionFn) Then(second SendActionFn) SendActionFn {
	return ComposeSendActionFn(p, second)

}

func ComposeSendActionFn(postProcessing ...SendActionFn) SendActionFn {
	toRun := make([]SendActionFn, 0, len(postProcessing))
	for _, p := range postProcessing {
		if p != nil {
			toRun = append(toRun, p)
		}
	}

	switch len(toRun){
	case 0:
		return nil
	case 1:
		return toRun[0]
	default:
		return func(ctx context.Context, fromAddr, toAddr sdk.AccAddress, amt sdk.Coins) error {
			var err error
			for _, p := range toRun {
				err = p(ctx, fromAddr, toAddr, amt) 
				if err != nil {
					return err
				}
			}
			return nil
		}
	}
}

var _ SendActionFn = NoOpSendActionFn

func NoOpSendActionFn(_ context.Context, _, toAddr sdk.AccAddress, _ sdk.Coins) error {
	return nil
}

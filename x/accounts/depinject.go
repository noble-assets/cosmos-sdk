package accounts

import (
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/gogo/protobuf/codec"
)

func init() {

}

type AccountsInputs struct {
	depinject.In

	Cdc      codec.Codec
	Registry cdctypes.InterfaceRegistry
}

type AccountsOutputs struct {
	depinject.Out

	Module appmodule.AppModule
}

func ProvideModule(in AccountsInputs) AccountsOutputs {
	return AccountsOutputs{}
}

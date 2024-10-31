package accounts

import (
	"encoding/json"
	"fmt"

	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/errors"
	v1 "cosmossdk.io/x/accounts/v1"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

const (
	ModuleName       = "accounts"
	StoreKey         = "_" + ModuleName // unfortunately accounts collides with auth store key
	ConsensusVersion = 1
)

var (
	ModuleAccountAddress = authtypes.NewModuleAddress(ModuleName)
)

var (
	_ module.AppModuleBasic      = AppModuleBasic{}
	_ module.HasGenesis          = AppModule{}
	_ module.HasConsensusVersion = AppModule{}

	_ appmodule.AppModule   = AppModule{}
	_ appmodule.HasServices = AppModule{}
)

// AppModuleBasic defines the basic application module used by the accounts module.
type AppModuleBasic struct {
	cdc codec.Codec
}

// AppModule implements the sdk.AppModule interface
type AppModule struct {
	AppModuleBasic
	keeper Keeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(cdc codec.Codec, k Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{cdc: cdc},
		keeper:         k,
	}
}

// Name implements module.AppModuleBasic.
func (a AppModuleBasic) Name() string { return ModuleName }

// RegisterGRPCGatewayRoutes implements module.AppModuleBasic.
func (a AppModuleBasic) RegisterGRPCGatewayRoutes(client.Context, *runtime.ServeMux) {}

// RegisterInterfaces implements module.AppModuleBasic.
func (a AppModuleBasic) RegisterInterfaces(registry types.InterfaceRegistry) {
	msgservice.RegisterMsgServiceDesc(registry, v1.MsgServiceDesc())
}

// RegisterLegacyAminoCodec implements module.AppModuleBasic.
func (a AppModuleBasic) RegisterLegacyAminoCodec(*codec.LegacyAmino) {
	panic("unimplemented")
}

// RegisterServices implements appmodule.HasServices.
func (a AppModule) RegisterServices(grpc.ServiceRegistrar) error {
	panic("unimplemented")
}

// ConsensusVersion implements module.HasConsensusVersion.
func (a AppModule) ConsensusVersion() uint64 { return ConsensusVersion }

// DefaultGenesis implements module.HasGenesis.
func (a AppModule) DefaultGenesis(codec.JSONCodec) json.RawMessage {
	return a.cdc.MustMarshalJSON(&v1.GenesisState{})
}

// ExportGenesis implements module.HasGenesis.
func (a AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	gs, err := a.keeper.ExportState(ctx)
	if err != nil {
		panic(fmt.Sprintf("failed to export genesis state: %s", err))
	}
	return cdc.MustMarshalJSON(gs)
}

// InitGenesis implements module.HasGenesis.
func (a AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) {
	var genesisState v1.GenesisState
	cdc.MustUnmarshalJSON(data, &genesisState)
	if err := a.keeper.ImportState(ctx, &genesisState); err != nil {
		panic(fmt.Sprintf("failed to import genesis state: %s", err))
	}
}

// ValidateGenesis implements module.HasGenesis.
func (a AppModule) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
	var data v1.GenesisState
	if err := cdc.UnmarshalJSON(bz, &data); err != nil {
		return errors.Wrapf(err, "failed to unmarshal %s genesis state", ModuleName)
	}

	return nil // TODO: validate
}

// IsAppModule implements appmodule.AppModule.
func (a AppModule) IsAppModule() {}

// IsOnePerModuleType implements appmodule.AppModule.
func (a AppModule) IsOnePerModuleType() {}

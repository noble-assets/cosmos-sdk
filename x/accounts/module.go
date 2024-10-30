package accounts

import (
	"encoding/json"

	"cosmossdk.io/core/appmodule"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
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
}

// NewAppModule creates a new AppModule object
func NewAppModule(cdc codec.Codec) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{cdc: cdc},
	}
}

// Name implements module.AppModuleBasic.
func (a AppModuleBasic) Name() string { return ModuleName }

// RegisterGRPCGatewayRoutes implements module.AppModuleBasic.
func (a AppModuleBasic) RegisterGRPCGatewayRoutes(client.Context, *runtime.ServeMux) {
	panic("unimplemented")
}

// RegisterInterfaces implements module.AppModuleBasic.
func (a AppModuleBasic) RegisterInterfaces(types.InterfaceRegistry) {
	panic("unimplemented")
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
	panic("unimplemented")
}

// ExportGenesis implements module.HasGenesis.
func (a AppModule) ExportGenesis(sdk.Context, codec.JSONCodec) json.RawMessage {
	panic("unimplemented")
}

// InitGenesis implements module.HasGenesis.
func (a AppModule) InitGenesis(sdk.Context, codec.JSONCodec, json.RawMessage) {
	panic("unimplemented")
}

// ValidateGenesis implements module.HasGenesis.
func (a AppModule) ValidateGenesis(codec.JSONCodec, client.TxEncodingConfig, json.RawMessage) error {
	panic("unimplemented")
}

// IsAppModule implements appmodule.AppModule.
func (a AppModule) IsAppModule() {}

// IsOnePerModuleType implements appmodule.AppModule.
func (a AppModule) IsOnePerModuleType() {}

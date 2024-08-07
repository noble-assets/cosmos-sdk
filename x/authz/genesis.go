package authz

import (
	"fmt"

	gogoprotoany "github.com/cosmos/gogoproto/types/any"
)

// NewGenesisState creates new GenesisState object
func NewGenesisState(entries []GrantAuthorization) *GenesisState {
	return &GenesisState{
		Authorization: entries,
	}
}

// ValidateGenesis check the given genesis state has no integrity issues
func ValidateGenesis(data GenesisState) error {
	for i, a := range data.Authorization {
		if a.Grantee == "" {
			return fmt.Errorf("authorization: %d, missing grantee", i)
		}
		if a.Granter == "" {
			return fmt.Errorf("authorization: %d,missing granter", i)
		}

	}
	return nil
}

// DefaultGenesisState - Return a default genesis state
func DefaultGenesisState() *GenesisState {
	return &GenesisState{}
}

var _ gogoprotoany.UnpackInterfacesMessage = GenesisState{}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (data GenesisState) UnpackInterfaces(unpacker gogoprotoany.AnyUnpacker) error {
	for _, a := range data.Authorization {
		err := a.UnpackInterfaces(unpacker)
		if err != nil {
			return err
		}
	}
	return nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (msg GrantAuthorization) UnpackInterfaces(unpacker gogoprotoany.AnyUnpacker) error {
	var a Authorization
	return unpacker.UnpackAny(msg.Authorization, &a)
}

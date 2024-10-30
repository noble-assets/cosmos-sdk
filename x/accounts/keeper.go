package accounts

import (
	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	corestoretypes "cosmossdk.io/core/store"
	"cosmossdk.io/x/accounts/accountstd"
	"cosmossdk.io/x/accounts/internal/implementation"
	"github.com/cosmos/cosmos-sdk/codec"
)

var (
	// AccountTypeKeyPrefix is the prefix for the account type key.
	AccountTypeKeyPrefix = collections.NewPrefix(0)
	// AccountNumberKey is the key for the account number.
	AccountNumberKey = collections.NewPrefix(1)
	// AccountByNumber is the key for the accounts by number.
	AccountByNumber = collections.NewPrefix(2)
)

type Keeper struct {
	addressCodec address.Codec
	accounts     map[string]implementation.Implementation

	// Schema is the schema for the module.
	Schema collections.Schema
	// AccountNumber is the last global account number.
	AccountNumber collections.Sequence
	// AccountsByType maps account address to their implementation.
	AccountsByType collections.Map[[]byte, string]
	// AccountByNumber maps account number to their address.
	AccountByNumber collections.Map[[]byte, uint64]
	// AccountsState keeps track of the state of each account.
	// NOTE: this is only used for genesis import and export.
	// Account set and get their own state but this helps providing a nice mapping
	// between: (account number, account state key) => account state value.
	AccountsState collections.Map[collections.Pair[uint64, []byte], []byte]
}

func NewKeeper(cdc codec.Codec,
	storeService corestoretypes.KVStoreService,
	addressCodec address.Codec,
	accounts ...accountstd.AccountCreatorFunc,
) (Keeper, error) {
	sb := collections.NewSchemaBuilder(storeService)
	keeper := Keeper{
		addressCodec:    addressCodec,
		AccountNumber:   collections.NewSequence(sb, AccountNumberKey, "account_number"),
		AccountsByType:  collections.NewMap(sb, AccountTypeKeyPrefix, "accounts_by_type", collections.BytesKey, collections.StringValue),
		AccountByNumber: collections.NewMap(sb, AccountByNumber, "account_by_number", collections.BytesKey, collections.Uint64Value),
		AccountsState:   collections.NewMap(sb, implementation.AccountStatePrefix, "accounts_state", collections.PairKeyCodec(collections.Uint64Key, collections.BytesKey), collections.BytesValue),
	}
	schema, err := sb.Build()
	if err != nil {
		return Keeper{}, err
	}
	keeper.Schema = schema
	keeper.accounts, err = implementation.MakeAccountsMap(keeper.addressCodec, accounts)
	if err != nil {
		return Keeper{}, err
	}
	return keeper, nil
}

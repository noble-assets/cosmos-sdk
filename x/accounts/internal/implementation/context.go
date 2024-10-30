package implementation

import (
	"bytes"
	"context"
	"errors"

	"cosmossdk.io/core/store"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/runtime/protoiface"
)

var errUnauthorized = errors.New("unauthorized")

type contextKey struct{}

type contextValue struct {
	store             store.KVStore                           // store is the prefixed store for the account.
	sender            []byte                                  // sender is the address of the entity invoking the account action.
	whoami            []byte                                  // whoami is the address of the account being invoked.
	originalContext   context.Context                         // originalContext that was used to build the account context.
	getExpectedSender func(msg proto.Message) ([]byte, error) // getExpectedSender is a function that returns the expected sender for a given message.
	moduleExec        ModuleExecFunc                          // moduleExec is a function that executes a module message.
	moduleQuery       ModuleQueryFunc                         // moduleQuery is a function that queries a module.
}

type (
	ModuleExecFunc  = func(ctx context.Context, msg, msgResp protoiface.MessageV1) error
	ModuleQueryFunc = ModuleExecFunc
)

// OpenKVStore returns the prefixed store for the account given the context.
func OpenKVStore(ctx context.Context) store.KVStore {
	return ctx.Value(contextKey{}).(contextValue).store
}

// Sender returns the address of the entity invoking the account action.
func Sender(ctx context.Context) []byte {
	return ctx.Value(contextKey{}).(contextValue).sender
}

// Whoami returns the address of the account being invoked.
func Whoami(ctx context.Context) []byte {
	return ctx.Value(contextKey{}).(contextValue).whoami
}

// ExecModule can be used to execute a message towards a module.
func ExecModule[Resp any, RespProto ProtoMsg[Resp], Req any, ReqProto ProtoMsg[Req]](ctx context.Context, msg ReqProto) (RespProto, error) {
	// get sender
	v := ctx.Value(contextKey{}).(contextValue)
	// check sender
	expectedSender, err := v.getExpectedSender(msg)
	if err != nil {
		return nil, err
	}
	if !bytes.Equal(expectedSender, v.whoami) {
		return nil, errUnauthorized
	}

	// execute module, unwrapping the original context.
	resp := RespProto(new(Resp))
	err = v.moduleExec(v.originalContext, msg, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// QueryModule can be used by an account to execute a module query.
func QueryModule[Resp any, RespProto ProtoMsg[Resp], Req any, ReqProto ProtoMsg[Req]](ctx context.Context, req ReqProto) (RespProto, error) {
	// we do not need to check the sender in a query because it is not a state transition.
	// we also unwrap the original context.
	v := ctx.Value(contextKey{}).(contextValue)
	resp := RespProto(new(Resp))
	err := v.moduleQuery(v.originalContext, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

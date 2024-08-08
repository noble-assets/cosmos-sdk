// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: cosmos/evidence/v1beta1/query.proto

package evidencev1beta1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Query_Evidence_FullMethodName    = "/cosmos.evidence.v1beta1.Query/Evidence"
	Query_AllEvidence_FullMethodName = "/cosmos.evidence.v1beta1.Query/AllEvidence"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Query defines the gRPC querier service.
type QueryClient interface {
	// Evidence queries evidence based on evidence hash.
	Evidence(ctx context.Context, in *QueryEvidenceRequest, opts ...grpc.CallOption) (*QueryEvidenceResponse, error)
	// AllEvidence queries all evidence.
	AllEvidence(ctx context.Context, in *QueryAllEvidenceRequest, opts ...grpc.CallOption) (*QueryAllEvidenceResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Evidence(ctx context.Context, in *QueryEvidenceRequest, opts ...grpc.CallOption) (*QueryEvidenceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryEvidenceResponse)
	err := c.cc.Invoke(ctx, Query_Evidence_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AllEvidence(ctx context.Context, in *QueryAllEvidenceRequest, opts ...grpc.CallOption) (*QueryAllEvidenceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryAllEvidenceResponse)
	err := c.cc.Invoke(ctx, Query_AllEvidence_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility.
//
// Query defines the gRPC querier service.
type QueryServer interface {
	// Evidence queries evidence based on evidence hash.
	Evidence(context.Context, *QueryEvidenceRequest) (*QueryEvidenceResponse, error)
	// AllEvidence queries all evidence.
	AllEvidence(context.Context, *QueryAllEvidenceRequest) (*QueryAllEvidenceResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedQueryServer struct{}

func (UnimplementedQueryServer) Evidence(context.Context, *QueryEvidenceRequest) (*QueryEvidenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Evidence not implemented")
}
func (UnimplementedQueryServer) AllEvidence(context.Context, *QueryAllEvidenceRequest) (*QueryAllEvidenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllEvidence not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}
func (UnimplementedQueryServer) testEmbeddedByValue()               {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	// If the following call pancis, it indicates UnimplementedQueryServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Evidence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEvidenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Evidence(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Evidence_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Evidence(ctx, req.(*QueryEvidenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AllEvidence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllEvidenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AllEvidence(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_AllEvidence_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AllEvidence(ctx, req.(*QueryAllEvidenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.evidence.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Evidence",
			Handler:    _Query_Evidence_Handler,
		},
		{
			MethodName: "AllEvidence",
			Handler:    _Query_AllEvidence_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/evidence/v1beta1/query.proto",
}

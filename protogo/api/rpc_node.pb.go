// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/rpc_node.proto

package api

import (
	common "chainmaker.org/chainmaker-go/pb/protogo/common"
	config "chainmaker.org/chainmaker-go/pb/protogo/config"
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("api/rpc_node.proto", fileDescriptor_ba278e4b8f6bb771) }

var fileDescriptor_ba278e4b8f6bb771 = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x8b, 0xd3, 0x40,
	0x18, 0x86, 0x13, 0x0a, 0x8a, 0xe3, 0xa9, 0x63, 0xb5, 0x6d, 0x84, 0xa0, 0xbd, 0x28, 0x88, 0x89,
	0x28, 0x78, 0xf1, 0xd6, 0x0a, 0x5e, 0x6a, 0x85, 0x54, 0x3d, 0x88, 0x52, 0x26, 0x93, 0xaf, 0xe9,
	0xd0, 0x34, 0xdf, 0x38, 0x93, 0x54, 0x7f, 0xc6, 0xfe, 0xac, 0x3d, 0xf6, 0xb8, 0xc7, 0xa5, 0xfd,
	0x01, 0xfb, 0x17, 0x96, 0x26, 0x99, 0x34, 0xcb, 0x86, 0xdd, 0xe3, 0xbc, 0xcf, 0xfb, 0xbd, 0xf3,
	0xce, 0xf0, 0x11, 0xca, 0xa4, 0xf0, 0x95, 0xe4, 0x8b, 0x14, 0x23, 0xf0, 0xa4, 0xc2, 0x0c, 0x69,
	0x87, 0x49, 0xe1, 0xf4, 0x38, 0x6e, 0x36, 0x98, 0xfa, 0x0a, 0xfe, 0xe6, 0xa0, 0xb3, 0x12, 0x39,
	0x4f, 0x6a, 0x55, 0xe7, 0x89, 0x11, 0x87, 0x1c, 0xd3, 0xa5, 0x88, 0xfd, 0x04, 0x39, 0x4b, 0x16,
	0xe5, 0xa1, 0x42, 0xfd, 0x1a, 0xc5, 0x37, 0x81, 0x5b, 0x01, 0xbe, 0x62, 0x22, 0xdd, 0xb0, 0x35,
	0xa8, 0x85, 0x06, 0xb5, 0x05, 0x55, 0xf2, 0xf7, 0x57, 0x1d, 0xf2, 0x30, 0x90, 0x7c, 0x86, 0x11,
	0xd0, 0x8f, 0xe4, 0xf1, 0x1c, 0xd2, 0x28, 0x28, 0x9b, 0xd0, 0xae, 0x57, 0x96, 0xf0, 0xbe, 0xff,
	0xaf, 0x24, 0x87, 0x36, 0x25, 0x2d, 0x31, 0xd5, 0x30, 0xb2, 0xe8, 0x27, 0xf2, 0x68, 0x9e, 0x87,
	0x9a, 0x2b, 0x11, 0x42, 0xdb, 0x54, 0xdf, 0x48, 0xb5, 0x2b, 0x28, 0x9e, 0x35, 0xb2, 0xde, 0xd9,
	0x74, 0x46, 0xba, 0x3f, 0x64, 0xc4, 0x32, 0xf8, 0x0c, 0x61, 0x1e, 0x4f, 0x8a, 0xb6, 0xd4, 0xf1,
	0xaa, 0x47, 0x34, 0x44, 0x93, 0xf6, 0xbc, 0x95, 0xd5, 0x65, 0xbe, 0x91, 0x67, 0x01, 0x2c, 0x15,
	0xe8, 0xd5, 0x14, 0xe3, 0x29, 0x6c, 0x21, 0xd1, 0x55, 0xe8, 0xc0, 0x0c, 0xd6, 0xc0, 0x44, 0x0e,
	0x5b, 0x48, 0x1d, 0xf8, 0x87, 0xf4, 0xbe, 0x40, 0x36, 0x39, 0xfe, 0xdf, 0xd7, 0xe3, 0xff, 0xfd,
	0x04, 0xa5, 0x05, 0xa6, 0xf4, 0x85, 0x19, 0xba, 0x85, 0x4c, 0xec, 0xcb, 0x3b, 0x1c, 0x75, 0x3c,
	0x92, 0xc1, 0x64, 0x05, 0x7c, 0x3d, 0x83, 0x7f, 0xe3, 0x04, 0xf9, 0xba, 0xf0, 0x56, 0x8d, 0x5f,
	0x9d, 0x02, 0xda, 0x1d, 0xe6, 0xa6, 0xd7, 0xf7, 0x1b, 0xcd, 0x85, 0xe3, 0xdf, 0xe7, 0x7b, 0xd7,
	0xde, 0xed, 0x5d, 0xfb, 0x72, 0xef, 0xda, 0x67, 0x07, 0xd7, 0xda, 0x1d, 0x5c, 0xeb, 0xe2, 0xe0,
	0x5a, 0xe4, 0x29, 0xaa, 0xd8, 0x3b, 0x2d, 0x8a, 0x27, 0x43, 0x8f, 0x49, 0xf1, 0xeb, 0x4d, 0x43,
	0x42, 0xd5, 0x5c, 0xa5, 0xb7, 0x31, 0xfa, 0x32, 0xf4, 0x8b, 0x4d, 0x8a, 0xd1, 0x67, 0x52, 0x84,
	0x0f, 0x8a, 0xc3, 0x87, 0xeb, 0x00, 0x00, 0x00, 0xff, 0xff, 0xeb, 0xa6, 0x3f, 0x10, 0xf0, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RpcNodeClient is the client API for RpcNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RpcNodeClient interface {
	// processing transaction message requests
	SendRequest(ctx context.Context, in *common.TxRequest, opts ...grpc.CallOption) (*common.TxResponse, error)
	// processing requests for message subscription
	Subscribe(ctx context.Context, in *common.TxRequest, opts ...grpc.CallOption) (RpcNode_SubscribeClient, error)
	// update debug status (development debugging)
	UpdateDebugConfig(ctx context.Context, in *config.DebugConfigRequest, opts ...grpc.CallOption) (*config.DebugConfigResponse, error)
	// refreshLogLevelsConfig
	RefreshLogLevelsConfig(ctx context.Context, in *config.LogLevelsRequest, opts ...grpc.CallOption) (*config.LogLevelsResponse, error)
	// get chainmaker version
	GetChainMakerVersion(ctx context.Context, in *config.ChainMakerVersionRequest, opts ...grpc.CallOption) (*config.ChainMakerVersionResponse, error)
	// check chain configuration and load new chain dynamically
	CheckNewBlockChainConfig(ctx context.Context, in *config.CheckNewBlockChainConfigRequest, opts ...grpc.CallOption) (*config.CheckNewBlockChainConfigResponse, error)
}

type rpcNodeClient struct {
	cc *grpc.ClientConn
}

func NewRpcNodeClient(cc *grpc.ClientConn) RpcNodeClient {
	return &rpcNodeClient{cc}
}

func (c *rpcNodeClient) SendRequest(ctx context.Context, in *common.TxRequest, opts ...grpc.CallOption) (*common.TxResponse, error) {
	out := new(common.TxResponse)
	err := c.cc.Invoke(ctx, "/api.RpcNode/SendRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcNodeClient) Subscribe(ctx context.Context, in *common.TxRequest, opts ...grpc.CallOption) (RpcNode_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RpcNode_serviceDesc.Streams[0], "/api.RpcNode/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &rpcNodeSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RpcNode_SubscribeClient interface {
	Recv() (*common.SubscribeResult, error)
	grpc.ClientStream
}

type rpcNodeSubscribeClient struct {
	grpc.ClientStream
}

func (x *rpcNodeSubscribeClient) Recv() (*common.SubscribeResult, error) {
	m := new(common.SubscribeResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rpcNodeClient) UpdateDebugConfig(ctx context.Context, in *config.DebugConfigRequest, opts ...grpc.CallOption) (*config.DebugConfigResponse, error) {
	out := new(config.DebugConfigResponse)
	err := c.cc.Invoke(ctx, "/api.RpcNode/UpdateDebugConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcNodeClient) RefreshLogLevelsConfig(ctx context.Context, in *config.LogLevelsRequest, opts ...grpc.CallOption) (*config.LogLevelsResponse, error) {
	out := new(config.LogLevelsResponse)
	err := c.cc.Invoke(ctx, "/api.RpcNode/RefreshLogLevelsConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcNodeClient) GetChainMakerVersion(ctx context.Context, in *config.ChainMakerVersionRequest, opts ...grpc.CallOption) (*config.ChainMakerVersionResponse, error) {
	out := new(config.ChainMakerVersionResponse)
	err := c.cc.Invoke(ctx, "/api.RpcNode/GetChainMakerVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcNodeClient) CheckNewBlockChainConfig(ctx context.Context, in *config.CheckNewBlockChainConfigRequest, opts ...grpc.CallOption) (*config.CheckNewBlockChainConfigResponse, error) {
	out := new(config.CheckNewBlockChainConfigResponse)
	err := c.cc.Invoke(ctx, "/api.RpcNode/CheckNewBlockChainConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RpcNodeServer is the server API for RpcNode service.
type RpcNodeServer interface {
	// processing transaction message requests
	SendRequest(context.Context, *common.TxRequest) (*common.TxResponse, error)
	// processing requests for message subscription
	Subscribe(*common.TxRequest, RpcNode_SubscribeServer) error
	// update debug status (development debugging)
	UpdateDebugConfig(context.Context, *config.DebugConfigRequest) (*config.DebugConfigResponse, error)
	// refreshLogLevelsConfig
	RefreshLogLevelsConfig(context.Context, *config.LogLevelsRequest) (*config.LogLevelsResponse, error)
	// get chainmaker version
	GetChainMakerVersion(context.Context, *config.ChainMakerVersionRequest) (*config.ChainMakerVersionResponse, error)
	// check chain configuration and load new chain dynamically
	CheckNewBlockChainConfig(context.Context, *config.CheckNewBlockChainConfigRequest) (*config.CheckNewBlockChainConfigResponse, error)
}

// UnimplementedRpcNodeServer can be embedded to have forward compatible implementations.
type UnimplementedRpcNodeServer struct {
}

func (*UnimplementedRpcNodeServer) SendRequest(ctx context.Context, req *common.TxRequest) (*common.TxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendRequest not implemented")
}
func (*UnimplementedRpcNodeServer) Subscribe(req *common.TxRequest, srv RpcNode_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (*UnimplementedRpcNodeServer) UpdateDebugConfig(ctx context.Context, req *config.DebugConfigRequest) (*config.DebugConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDebugConfig not implemented")
}
func (*UnimplementedRpcNodeServer) RefreshLogLevelsConfig(ctx context.Context, req *config.LogLevelsRequest) (*config.LogLevelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshLogLevelsConfig not implemented")
}
func (*UnimplementedRpcNodeServer) GetChainMakerVersion(ctx context.Context, req *config.ChainMakerVersionRequest) (*config.ChainMakerVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChainMakerVersion not implemented")
}
func (*UnimplementedRpcNodeServer) CheckNewBlockChainConfig(ctx context.Context, req *config.CheckNewBlockChainConfigRequest) (*config.CheckNewBlockChainConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckNewBlockChainConfig not implemented")
}

func RegisterRpcNodeServer(s *grpc.Server, srv RpcNodeServer) {
	s.RegisterService(&_RpcNode_serviceDesc, srv)
}

func _RpcNode_SendRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.TxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcNodeServer).SendRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.RpcNode/SendRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcNodeServer).SendRequest(ctx, req.(*common.TxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcNode_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(common.TxRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RpcNodeServer).Subscribe(m, &rpcNodeSubscribeServer{stream})
}

type RpcNode_SubscribeServer interface {
	Send(*common.SubscribeResult) error
	grpc.ServerStream
}

type rpcNodeSubscribeServer struct {
	grpc.ServerStream
}

func (x *rpcNodeSubscribeServer) Send(m *common.SubscribeResult) error {
	return x.ServerStream.SendMsg(m)
}

func _RpcNode_UpdateDebugConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(config.DebugConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcNodeServer).UpdateDebugConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.RpcNode/UpdateDebugConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcNodeServer).UpdateDebugConfig(ctx, req.(*config.DebugConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcNode_RefreshLogLevelsConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(config.LogLevelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcNodeServer).RefreshLogLevelsConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.RpcNode/RefreshLogLevelsConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcNodeServer).RefreshLogLevelsConfig(ctx, req.(*config.LogLevelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcNode_GetChainMakerVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(config.ChainMakerVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcNodeServer).GetChainMakerVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.RpcNode/GetChainMakerVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcNodeServer).GetChainMakerVersion(ctx, req.(*config.ChainMakerVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcNode_CheckNewBlockChainConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(config.CheckNewBlockChainConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcNodeServer).CheckNewBlockChainConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.RpcNode/CheckNewBlockChainConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcNodeServer).CheckNewBlockChainConfig(ctx, req.(*config.CheckNewBlockChainConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RpcNode_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.RpcNode",
	HandlerType: (*RpcNodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendRequest",
			Handler:    _RpcNode_SendRequest_Handler,
		},
		{
			MethodName: "UpdateDebugConfig",
			Handler:    _RpcNode_UpdateDebugConfig_Handler,
		},
		{
			MethodName: "RefreshLogLevelsConfig",
			Handler:    _RpcNode_RefreshLogLevelsConfig_Handler,
		},
		{
			MethodName: "GetChainMakerVersion",
			Handler:    _RpcNode_GetChainMakerVersion_Handler,
		},
		{
			MethodName: "CheckNewBlockChainConfig",
			Handler:    _RpcNode_CheckNewBlockChainConfig_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _RpcNode_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/rpc_node.proto",
}

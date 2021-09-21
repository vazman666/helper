// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package gen

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SqlRequestClient is the client API for SqlRequest service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SqlRequestClient interface {
	StreamSql(ctx context.Context, in *Request, opts ...grpc.CallOption) (SqlRequest_StreamSqlClient, error)
}

type sqlRequestClient struct {
	cc grpc.ClientConnInterface
}

func NewSqlRequestClient(cc grpc.ClientConnInterface) SqlRequestClient {
	return &sqlRequestClient{cc}
}

func (c *sqlRequestClient) StreamSql(ctx context.Context, in *Request, opts ...grpc.CallOption) (SqlRequest_StreamSqlClient, error) {
	stream, err := c.cc.NewStream(ctx, &SqlRequest_ServiceDesc.Streams[0], "/fromsql.SqlRequest/StreamSql", opts...)
	if err != nil {
		return nil, err
	}
	x := &sqlRequestStreamSqlClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SqlRequest_StreamSqlClient interface {
	Recv() (*Answer, error)
	grpc.ClientStream
}

type sqlRequestStreamSqlClient struct {
	grpc.ClientStream
}

func (x *sqlRequestStreamSqlClient) Recv() (*Answer, error) {
	m := new(Answer)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SqlRequestServer is the server API for SqlRequest service.
// All implementations must embed UnimplementedSqlRequestServer
// for forward compatibility
type SqlRequestServer interface {
	StreamSql(*Request, SqlRequest_StreamSqlServer) error
	mustEmbedUnimplementedSqlRequestServer()
}

// UnimplementedSqlRequestServer must be embedded to have forward compatible implementations.
type UnimplementedSqlRequestServer struct {
}

func (UnimplementedSqlRequestServer) StreamSql(*Request, SqlRequest_StreamSqlServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamSql not implemented")
}
func (UnimplementedSqlRequestServer) mustEmbedUnimplementedSqlRequestServer() {}

// UnsafeSqlRequestServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SqlRequestServer will
// result in compilation errors.
type UnsafeSqlRequestServer interface {
	mustEmbedUnimplementedSqlRequestServer()
}

func RegisterSqlRequestServer(s grpc.ServiceRegistrar, srv SqlRequestServer) {
	s.RegisterService(&SqlRequest_ServiceDesc, srv)
}

func _SqlRequest_StreamSql_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SqlRequestServer).StreamSql(m, &sqlRequestStreamSqlServer{stream})
}

type SqlRequest_StreamSqlServer interface {
	Send(*Answer) error
	grpc.ServerStream
}

type sqlRequestStreamSqlServer struct {
	grpc.ServerStream
}

func (x *sqlRequestStreamSqlServer) Send(m *Answer) error {
	return x.ServerStream.SendMsg(m)
}

// SqlRequest_ServiceDesc is the grpc.ServiceDesc for SqlRequest service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SqlRequest_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fromsql.SqlRequest",
	HandlerType: (*SqlRequestServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamSql",
			Handler:       _SqlRequest_StreamSql_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "fromsql.proto",
}

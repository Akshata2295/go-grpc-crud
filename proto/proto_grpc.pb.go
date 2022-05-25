// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: proto/proto.proto

package __

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

// MovieServiceClient is the client API for MovieService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MovieServiceClient interface {
	CreateMovie(ctx context.Context, in *MovieRequest, opts ...grpc.CallOption) (*MovieResponse, error)
	GetMovie(ctx context.Context, in *MovieID, opts ...grpc.CallOption) (*MovieResponse, error)
	GetMovies(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*MovieListResponse, error)
	UpdateMovie(ctx context.Context, in *MovieRequest, opts ...grpc.CallOption) (*MovieResponse, error)
	DeleteMovie(ctx context.Context, in *MovieID, opts ...grpc.CallOption) (*Message, error)
}

type movieServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMovieServiceClient(cc grpc.ClientConnInterface) MovieServiceClient {
	return &movieServiceClient{cc}
}

func (c *movieServiceClient) CreateMovie(ctx context.Context, in *MovieRequest, opts ...grpc.CallOption) (*MovieResponse, error) {
	out := new(MovieResponse)
	err := c.cc.Invoke(ctx, "/proto.MovieService/CreateMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieServiceClient) GetMovie(ctx context.Context, in *MovieID, opts ...grpc.CallOption) (*MovieResponse, error) {
	out := new(MovieResponse)
	err := c.cc.Invoke(ctx, "/proto.MovieService/GetMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieServiceClient) GetMovies(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*MovieListResponse, error) {
	out := new(MovieListResponse)
	err := c.cc.Invoke(ctx, "/proto.MovieService/GetMovies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieServiceClient) UpdateMovie(ctx context.Context, in *MovieRequest, opts ...grpc.CallOption) (*MovieResponse, error) {
	out := new(MovieResponse)
	err := c.cc.Invoke(ctx, "/proto.MovieService/UpdateMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieServiceClient) DeleteMovie(ctx context.Context, in *MovieID, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/proto.MovieService/DeleteMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MovieServiceServer is the server API for MovieService service.
// All implementations must embed UnimplementedMovieServiceServer
// for forward compatibility
type MovieServiceServer interface {
	CreateMovie(context.Context, *MovieRequest) (*MovieResponse, error)
	GetMovie(context.Context, *MovieID) (*MovieResponse, error)
	GetMovies(context.Context, *EmptyRequest) (*MovieListResponse, error)
	UpdateMovie(context.Context, *MovieRequest) (*MovieResponse, error)
	DeleteMovie(context.Context, *MovieID) (*Message, error)
	mustEmbedUnimplementedMovieServiceServer()
}

// UnimplementedMovieServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMovieServiceServer struct {
}

func (UnimplementedMovieServiceServer) CreateMovie(context.Context, *MovieRequest) (*MovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMovie not implemented")
}
func (UnimplementedMovieServiceServer) GetMovie(context.Context, *MovieID) (*MovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovie not implemented")
}
func (UnimplementedMovieServiceServer) GetMovies(context.Context, *EmptyRequest) (*MovieListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovies not implemented")
}
func (UnimplementedMovieServiceServer) UpdateMovie(context.Context, *MovieRequest) (*MovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMovie not implemented")
}
func (UnimplementedMovieServiceServer) DeleteMovie(context.Context, *MovieID) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMovie not implemented")
}
func (UnimplementedMovieServiceServer) mustEmbedUnimplementedMovieServiceServer() {}

// UnsafeMovieServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MovieServiceServer will
// result in compilation errors.
type UnsafeMovieServiceServer interface {
	mustEmbedUnimplementedMovieServiceServer()
}

func RegisterMovieServiceServer(s grpc.ServiceRegistrar, srv MovieServiceServer) {
	s.RegisterService(&MovieService_ServiceDesc, srv)
}

func _MovieService_CreateMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServiceServer).CreateMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MovieService/CreateMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServiceServer).CreateMovie(ctx, req.(*MovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieService_GetMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MovieID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServiceServer).GetMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MovieService/GetMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServiceServer).GetMovie(ctx, req.(*MovieID))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieService_GetMovies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServiceServer).GetMovies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MovieService/GetMovies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServiceServer).GetMovies(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieService_UpdateMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServiceServer).UpdateMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MovieService/UpdateMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServiceServer).UpdateMovie(ctx, req.(*MovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieService_DeleteMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MovieID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServiceServer).DeleteMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MovieService/DeleteMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServiceServer).DeleteMovie(ctx, req.(*MovieID))
	}
	return interceptor(ctx, in, info, handler)
}

// MovieService_ServiceDesc is the grpc.ServiceDesc for MovieService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MovieService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MovieService",
	HandlerType: (*MovieServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMovie",
			Handler:    _MovieService_CreateMovie_Handler,
		},
		{
			MethodName: "GetMovie",
			Handler:    _MovieService_GetMovie_Handler,
		},
		{
			MethodName: "GetMovies",
			Handler:    _MovieService_GetMovies_Handler,
		},
		{
			MethodName: "UpdateMovie",
			Handler:    _MovieService_UpdateMovie_Handler,
		},
		{
			MethodName: "DeleteMovie",
			Handler:    _MovieService_DeleteMovie_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/proto.proto",
}

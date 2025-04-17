// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.2
// source: proto/movie_board_service/movie_board.proto

package movie_board_service

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
	MovieBoardService_GetMovies_FullMethodName    = "/movie_board.MovieBoardService/GetMovies"
	MovieBoardService_GetMovieById_FullMethodName = "/movie_board.MovieBoardService/GetMovieById"
	MovieBoardService_InsertMovie_FullMethodName  = "/movie_board.MovieBoardService/InsertMovie"
	MovieBoardService_UpdateMovie_FullMethodName  = "/movie_board.MovieBoardService/UpdateMovie"
	MovieBoardService_DeleteMovie_FullMethodName  = "/movie_board.MovieBoardService/DeleteMovie"
)

// MovieBoardServiceClient is the client API for MovieBoardService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MovieBoardServiceClient interface {
	GetMovies(ctx context.Context, in *GetMoviesRequest, opts ...grpc.CallOption) (*GetMoviesResponse, error)
	GetMovieById(ctx context.Context, in *GetMovieByIdRequest, opts ...grpc.CallOption) (*GetMovieByIdResponse, error)
	InsertMovie(ctx context.Context, in *InsertMovieRequest, opts ...grpc.CallOption) (*InsertMovieResponse, error)
	UpdateMovie(ctx context.Context, in *UpdateMovieRequest, opts ...grpc.CallOption) (*UpdateMovieResponse, error)
	DeleteMovie(ctx context.Context, in *DeleteMovieRequest, opts ...grpc.CallOption) (*DeleteMovieResponse, error)
}

type movieBoardServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMovieBoardServiceClient(cc grpc.ClientConnInterface) MovieBoardServiceClient {
	return &movieBoardServiceClient{cc}
}

func (c *movieBoardServiceClient) GetMovies(ctx context.Context, in *GetMoviesRequest, opts ...grpc.CallOption) (*GetMoviesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMoviesResponse)
	err := c.cc.Invoke(ctx, MovieBoardService_GetMovies_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieBoardServiceClient) GetMovieById(ctx context.Context, in *GetMovieByIdRequest, opts ...grpc.CallOption) (*GetMovieByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMovieByIdResponse)
	err := c.cc.Invoke(ctx, MovieBoardService_GetMovieById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieBoardServiceClient) InsertMovie(ctx context.Context, in *InsertMovieRequest, opts ...grpc.CallOption) (*InsertMovieResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InsertMovieResponse)
	err := c.cc.Invoke(ctx, MovieBoardService_InsertMovie_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieBoardServiceClient) UpdateMovie(ctx context.Context, in *UpdateMovieRequest, opts ...grpc.CallOption) (*UpdateMovieResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateMovieResponse)
	err := c.cc.Invoke(ctx, MovieBoardService_UpdateMovie_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieBoardServiceClient) DeleteMovie(ctx context.Context, in *DeleteMovieRequest, opts ...grpc.CallOption) (*DeleteMovieResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteMovieResponse)
	err := c.cc.Invoke(ctx, MovieBoardService_DeleteMovie_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MovieBoardServiceServer is the server API for MovieBoardService service.
// All implementations must embed UnimplementedMovieBoardServiceServer
// for forward compatibility.
type MovieBoardServiceServer interface {
	GetMovies(context.Context, *GetMoviesRequest) (*GetMoviesResponse, error)
	GetMovieById(context.Context, *GetMovieByIdRequest) (*GetMovieByIdResponse, error)
	InsertMovie(context.Context, *InsertMovieRequest) (*InsertMovieResponse, error)
	UpdateMovie(context.Context, *UpdateMovieRequest) (*UpdateMovieResponse, error)
	DeleteMovie(context.Context, *DeleteMovieRequest) (*DeleteMovieResponse, error)
	mustEmbedUnimplementedMovieBoardServiceServer()
}

// UnimplementedMovieBoardServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMovieBoardServiceServer struct{}

func (UnimplementedMovieBoardServiceServer) GetMovies(context.Context, *GetMoviesRequest) (*GetMoviesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovies not implemented")
}
func (UnimplementedMovieBoardServiceServer) GetMovieById(context.Context, *GetMovieByIdRequest) (*GetMovieByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovieById not implemented")
}
func (UnimplementedMovieBoardServiceServer) InsertMovie(context.Context, *InsertMovieRequest) (*InsertMovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertMovie not implemented")
}
func (UnimplementedMovieBoardServiceServer) UpdateMovie(context.Context, *UpdateMovieRequest) (*UpdateMovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMovie not implemented")
}
func (UnimplementedMovieBoardServiceServer) DeleteMovie(context.Context, *DeleteMovieRequest) (*DeleteMovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMovie not implemented")
}
func (UnimplementedMovieBoardServiceServer) mustEmbedUnimplementedMovieBoardServiceServer() {}
func (UnimplementedMovieBoardServiceServer) testEmbeddedByValue()                           {}

// UnsafeMovieBoardServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MovieBoardServiceServer will
// result in compilation errors.
type UnsafeMovieBoardServiceServer interface {
	mustEmbedUnimplementedMovieBoardServiceServer()
}

func RegisterMovieBoardServiceServer(s grpc.ServiceRegistrar, srv MovieBoardServiceServer) {
	// If the following call pancis, it indicates UnimplementedMovieBoardServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MovieBoardService_ServiceDesc, srv)
}

func _MovieBoardService_GetMovies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMoviesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieBoardServiceServer).GetMovies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MovieBoardService_GetMovies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieBoardServiceServer).GetMovies(ctx, req.(*GetMoviesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieBoardService_GetMovieById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMovieByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieBoardServiceServer).GetMovieById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MovieBoardService_GetMovieById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieBoardServiceServer).GetMovieById(ctx, req.(*GetMovieByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieBoardService_InsertMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieBoardServiceServer).InsertMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MovieBoardService_InsertMovie_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieBoardServiceServer).InsertMovie(ctx, req.(*InsertMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieBoardService_UpdateMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieBoardServiceServer).UpdateMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MovieBoardService_UpdateMovie_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieBoardServiceServer).UpdateMovie(ctx, req.(*UpdateMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieBoardService_DeleteMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieBoardServiceServer).DeleteMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MovieBoardService_DeleteMovie_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieBoardServiceServer).DeleteMovie(ctx, req.(*DeleteMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MovieBoardService_ServiceDesc is the grpc.ServiceDesc for MovieBoardService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MovieBoardService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "movie_board.MovieBoardService",
	HandlerType: (*MovieBoardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMovies",
			Handler:    _MovieBoardService_GetMovies_Handler,
		},
		{
			MethodName: "GetMovieById",
			Handler:    _MovieBoardService_GetMovieById_Handler,
		},
		{
			MethodName: "InsertMovie",
			Handler:    _MovieBoardService_InsertMovie_Handler,
		},
		{
			MethodName: "UpdateMovie",
			Handler:    _MovieBoardService_UpdateMovie_Handler,
		},
		{
			MethodName: "DeleteMovie",
			Handler:    _MovieBoardService_DeleteMovie_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/movie_board_service/movie_board.proto",
}

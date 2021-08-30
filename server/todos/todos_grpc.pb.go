// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package todos

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

// TodoAppClient is the client API for TodoApp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodoAppClient interface {
	CreateTodo(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*Todo, error)
	GetAllTodos(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Todos, error)
}

type todoAppClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoAppClient(cc grpc.ClientConnInterface) TodoAppClient {
	return &todoAppClient{cc}
}

func (c *todoAppClient) CreateTodo(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*Todo, error) {
	out := new(Todo)
	err := c.cc.Invoke(ctx, "/todos.TodoApp/CreateTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoAppClient) GetAllTodos(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Todos, error) {
	out := new(Todos)
	err := c.cc.Invoke(ctx, "/todos.TodoApp/GetAllTodos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoAppServer is the server API for TodoApp service.
// All implementations must embed UnimplementedTodoAppServer
// for forward compatibility
type TodoAppServer interface {
	CreateTodo(context.Context, *Todo) (*Todo, error)
	GetAllTodos(context.Context, *Empty) (*Todos, error)
	mustEmbedUnimplementedTodoAppServer()
}

// UnimplementedTodoAppServer must be embedded to have forward compatible implementations.
type UnimplementedTodoAppServer struct {
}

func (UnimplementedTodoAppServer) CreateTodo(context.Context, *Todo) (*Todo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTodo not implemented")
}
func (UnimplementedTodoAppServer) GetAllTodos(context.Context, *Empty) (*Todos, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTodos not implemented")
}
func (UnimplementedTodoAppServer) mustEmbedUnimplementedTodoAppServer() {}

// UnsafeTodoAppServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodoAppServer will
// result in compilation errors.
type UnsafeTodoAppServer interface {
	mustEmbedUnimplementedTodoAppServer()
}

func RegisterTodoAppServer(s grpc.ServiceRegistrar, srv TodoAppServer) {
	s.RegisterService(&TodoApp_ServiceDesc, srv)
}

func _TodoApp_CreateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Todo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoAppServer).CreateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todos.TodoApp/CreateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoAppServer).CreateTodo(ctx, req.(*Todo))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoApp_GetAllTodos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoAppServer).GetAllTodos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todos.TodoApp/GetAllTodos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoAppServer).GetAllTodos(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// TodoApp_ServiceDesc is the grpc.ServiceDesc for TodoApp service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TodoApp_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "todos.TodoApp",
	HandlerType: (*TodoAppServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTodo",
			Handler:    _TodoApp_CreateTodo_Handler,
		},
		{
			MethodName: "GetAllTodos",
			Handler:    _TodoApp_GetAllTodos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todos.proto",
}

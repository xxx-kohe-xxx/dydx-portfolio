/**
 * @fileoverview gRPC-Web generated client stub for todos
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as todos_pb from './todos_pb';


export class TodoAppClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoCreateTodo = new grpcWeb.AbstractClientBase.MethodInfo(
    todos_pb.Todo,
    (request: todos_pb.Todo) => {
      return request.serializeBinary();
    },
    todos_pb.Todo.deserializeBinary
  );

  createTodo(
    request: todos_pb.Todo,
    metadata: grpcWeb.Metadata | null): Promise<todos_pb.Todo>;

  createTodo(
    request: todos_pb.Todo,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: todos_pb.Todo) => void): grpcWeb.ClientReadableStream<todos_pb.Todo>;

  createTodo(
    request: todos_pb.Todo,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: todos_pb.Todo) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/todos.TodoApp/CreateTodo',
        request,
        metadata || {},
        this.methodInfoCreateTodo,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/todos.TodoApp/CreateTodo',
    request,
    metadata || {},
    this.methodInfoCreateTodo);
  }

  methodInfoGetAllTodos = new grpcWeb.AbstractClientBase.MethodInfo(
    todos_pb.Todos,
    (request: todos_pb.Empty) => {
      return request.serializeBinary();
    },
    todos_pb.Todos.deserializeBinary
  );

  getAllTodos(
    request: todos_pb.Empty,
    metadata: grpcWeb.Metadata | null): Promise<todos_pb.Todos>;

  getAllTodos(
    request: todos_pb.Empty,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: todos_pb.Todos) => void): grpcWeb.ClientReadableStream<todos_pb.Todos>;

  getAllTodos(
    request: todos_pb.Empty,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: todos_pb.Todos) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/todos.TodoApp/GetAllTodos',
        request,
        metadata || {},
        this.methodInfoGetAllTodos,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/todos.TodoApp/GetAllTodos',
    request,
    metadata || {},
    this.methodInfoGetAllTodos);
  }

}


#!/bin/sh

set -xe

SERVER_OUTPUT_DIR="server/todos"
CLIENT_OUTPUT_DIR="client/src/todos"

# コード生成
protoc --proto_path=proto todos.proto \
  --go_out=${SERVER_OUTPUT_DIR} --go_opt=paths=source_relative \
  --go-grpc_out=${SERVER_OUTPUT_DIR} --go-grpc_opt=paths=source_relative \
  --js_out=import_style=commonjs:${CLIENT_OUTPUT_DIR} \
  --grpc-web_out=import_style=typescript,mode=grpcwebtext:${CLIENT_OUTPUT_DIR}
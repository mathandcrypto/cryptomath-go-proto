//go:generate protoc --proto_path=auth --go_out=./auth --go_opt=paths=source_relative --go-grpc_out=./auth --go-grpc_opt=paths=source_relative auth.proto

package proto
//go:generate protoc --proto_path=auth --go_out=./auth --go_opt=paths=source_relative --go-grpc_out=./auth --go-grpc_opt=paths=source_relative auth.proto
//go:generate protoc --proto_path=captcha --go_out=./captcha --go_opt=paths=source_relative --go-grpc_out=./captcha --go-grpc_opt=paths=source_relative captcha.proto

package generate
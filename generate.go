//go:generate protoc --proto_path=auth --go_out=./auth --go_opt=paths=source_relative --go-grpc_out=./auth --go-grpc_opt=paths=source_relative auth.proto
//go:generate protoc --proto_path=captcha --go_out=./captcha --go_opt=paths=source_relative --go-grpc_out=./captcha --go-grpc_opt=paths=source_relative captcha.proto
//go:generate protoc --proto_path=user --go_out=./user --go_opt=paths=source_relative --go-grpc_out=./user --go-grpc_opt=paths=source_relative user.proto

package generate
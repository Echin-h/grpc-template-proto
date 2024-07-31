protoc --go_out=gen/ --go_opt=paths=source_relative --go-grpc_out=gen/ --go-grpc_opt=paths=source_relative --grpc-gateway_out=gen/ --grpc-gateway_opt=paths=source_relative proto/hello/v1/hello.proto

:: protoc is the Protocol Buffers compiler

:: --go_out=../../../gen: Generate Go code in the gen directory
:: --go_opt=paths=source_relative: Use source-relative paths for generated Go code

:: --go-grpc_out=../../../gen: Generate gRPC code in the gen directory
:: --go-grpc_opt=paths=source_relative: Use source-relative paths for generated gRPC code

:: hello.proto: The Protocol Buffers file to compile
:: hello.proto means the creation for the gen directory as 'gen/hello.pb.go' and 'gen/hello_grpc.pb.go'
:: but if 'hello.proto' is replaced as 'proto/hello/v1/hello.proto' then the gen directory will created as the gen/proto/hello/v1 directory and the generated files will be placed in the gen/proto/hello/v1 directory
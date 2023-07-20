proto:
	protoc                                  \
		--go_out=.                          \
		--go_opt=paths=source_relative      \
    	--go-grpc_out=.                     \
		--go-grpc_opt=paths=source_relative \
    	helloworld/helloworld.proto

server:
	go run helloworld_server/main.go

client:
	go run helloworld_client/main.go
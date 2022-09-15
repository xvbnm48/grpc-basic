create:
	protoc --proto_path=proto proto/*.proto --go_out=gen/
	protoc --proto_path=proto proto/*.proto --go-grpc_out=gen/

clear:
	rm -rf gen/proto/*.go

client_side:
	go run client/client.go
all: proto main

main:
	go build .

proto:
	protoc --proto_path $(shell pwd)/protobuf --go_out=plugins=grpc:protobuf cvInterface.proto
	protoc --proto_path $(shell pwd)/protobuf --go_out=plugins=grpc:protobuf rvInterface.proto
	protoc --proto_path $(shell pwd)/protobuf --go_out=plugins=grpc:protobuf token.proto
	protoc --proto_path $(shell pwd)/protobuf --go_out=plugins=grpc:protobuf dbRequests.proto
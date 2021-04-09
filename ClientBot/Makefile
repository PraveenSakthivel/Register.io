all: proto

proto:
	protoc --proto_path $(shell pwd)/protobuf --go_out=plugins=grpc:protobuf cvInterface.proto
	protoc --proto_path $(shell pwd)/protobuf --go_out=plugins=grpc:protobuf rvInterface.proto
	protoc --proto_path $(shell pwd)/protobuf --go_out=plugins=grpc:protobuf token.proto
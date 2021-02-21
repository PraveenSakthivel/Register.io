all: proto

proto:	
		protoc --proto_path $(shell pwd)/protobuf --go_out=plugins=grpc:protobuf cvInterface.proto
		go build protobuf/cvinterface.pb.go
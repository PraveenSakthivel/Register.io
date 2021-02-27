all: proto prereq

proto:	
		protoc --proto_path $(shell pwd)/protobuf --go_out=plugins=grpc:protobuf cvInterface.proto
		go build protobuf/cvinterface.pb.go

prereq:
		cp ../PrereqValidation/protobuf/prereqInterface.pb.go prereq/prereqInterface.pb.go
		go build prereq/prereqInterface.pb.go

.PHONY: prereq
all: proto server test

server:	
		go build database/data.go
		go build prereqServer.go

proto:	
		protoc --proto_path $(shell pwd)/protobuf --go_out=plugins=grpc:protobuf prereqInterface.proto
		go build protobuf/prereqInterface.pb.go

test: 	
		go build -o tests/tests tests/tester.go

clean:	
		rm prereqServer
		rm tests/tests
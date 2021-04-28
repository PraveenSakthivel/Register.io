
all: proto server test

server:	
		go build rv.go

proto:	
		protoc --proto_path $(shell pwd)/protobuf --go_out=plugins=grpc:protobuf rvInterface.proto
		go build protobuf/rvinterface.pb.go

test: 	
		go build -o tests/tests tests/tester.go

clean:	
		rm rv
		rm tests/tests
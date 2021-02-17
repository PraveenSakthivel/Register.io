
all: proto server test

server:	
		go build rv.go

proto:	
		protoc --go_out=plugins=grpc:validation rvInterface.proto
		go build validation/rvinterface.pb.go

test: 	
		go build -o tests/tests tests/tester.go

clean:	
		rm rv
		rm tests/tests
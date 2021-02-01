
all: proto server test

server:	
		go build rv.go

proto:	
		protoc --go_out=. protobuf/*.proto
		go build protobuf/rvinterface.pb.go

test: 	
		go build -o tests/tests tests/tester.go

clean:	
		rm rv
		rm tests/tests
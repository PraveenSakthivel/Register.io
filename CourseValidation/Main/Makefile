all: proto prereq server test

server:
		go build database/data.go
		go build classTiming/classTiming.go
		go build cv.go

proto:	
		protoc --proto_path $(shell pwd)/protobuf --go_out=plugins=grpc:protobuf cvInterface.proto
		go build protobuf/cvinterface.pb.go

prereq:
		cp ../PrereqValidation/protobuf/prereqInterface.pb.go prereq/prereqInterface.pb.go
		go build prereq/prereqInterface.pb.go

test: 	
		go build -o tests/tests tests/tester.go
		go build -o spntests/tests spntests/spntester.go

clean:	
		rm cv
		rm tests/tests
		rm spntests/tests

.PHONY: prereq
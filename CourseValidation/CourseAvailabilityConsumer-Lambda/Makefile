all: consumer

consumer:
	go build consumer.go

proto:	
	protoc --proto_path $(shell pwd)/protobuf --go_out=plugins=grpc:protobuf consumerState.proto
	go build protobuf/consumerState.pb.go

docker:
	docker build -t consumer-lambda .

clean:
	rm consumer
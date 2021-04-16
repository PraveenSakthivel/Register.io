all: consumer

consumer:
	go build consumer.go

docker:
	docker build -t consumer .

clean:
	rm consumer
FROM golang:alpine AS build
WORKDIR /src/registerio/cv/consumer-lambda/

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY database/data.go database/data.go
COPY consumer.go . 
COPY secrets/secret.go secrets/secret.go
COPY protobuf/consumerState.pb.go protobuf/consumerState.pb.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" .

FROM scratch
WORKDIR /usr/bin/
COPY --from=build /src/registerio/cv/consumer-lambda/consumer-lambda /go/bin/consumer
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/go/bin/consumer"]
CMD ["14147"]
FROM golang:alpine AS build
WORKDIR /src/registerio/cv/consumer/

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY database/data.go database/data.go
COPY consumer.go . 
COPY secrets/secret.go secrets/secret.go
RUN go build .

FROM alpine
WORKDIR /usr/bin
COPY --from=build /src/registerio/cv/consumer/consumer /go/bin/consumer

CMD ["/go/bin/consumer"]
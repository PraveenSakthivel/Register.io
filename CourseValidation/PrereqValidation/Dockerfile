FROM golang:alpine AS build
WORKDIR /src/registerio/cv/preqreq/

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY protobuf/prereqInterface.pb.go protobuf/prereqInterface.pb.go
COPY database/data.go database/data.go
COPY secrets/secret.go secrets/secret.go
COPY prereqServer.go . 
WORKDIR /src/registerio/cv/preqreq/protobuf
RUN go build prereqInterface.pb.go
WORKDIR /src/registerio/cv/preqreq/secrets
RUN go build secret.go
WORKDIR /src/registerio/cv/preqreq/database
RUN go build data.go
WORKDIR /src/registerio/cv/preqreq
RUN go build prereqServer.go

FROM alpine
WORKDIR /usr/bin
COPY --from=build /src/registerio/cv/preqreq/prereqServer /go/bin/prereqServer

EXPOSE 8080

CMD ["/go/bin/prereqServer"]
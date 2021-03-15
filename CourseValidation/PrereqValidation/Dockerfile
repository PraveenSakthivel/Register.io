FROM golang:alpine AS build
WORKDIR /src/registerio/cv/preqreq/

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY protobuf/prereqInterface.pb.go protobuf/prereqInterface.pb.go
COPY database/data.go database/data.go
COPY prereqServer.go . 
WORKDIR /src/registerio/cv/preqreq/protobuf
RUN go build prereqInterface.pb.go
WORKDIR /src/registerio/cv/preqreq/database
RUN go build data.go
WORKDIR /src/registerio/cv/preqreq
RUN go build prereqServer.go

FROM alpine
WORKDIR /usr/bin
COPY --from=build /src/registerio/cv/preqreq/prereqServer /go/bin/prereqServer
ENV AWS_ACCESS_KEY_ID=AKIA53XJENGN4KQU3KFO
ENV AWS_SECRET_ACCESS_KEY=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
ENV AWS_REGION=us-east-1

EXPOSE 8081

CMD ["/go/bin/prereqServer"]
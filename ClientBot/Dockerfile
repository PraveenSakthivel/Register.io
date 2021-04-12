FROM golang:alpine AS build
WORKDIR /src/registerio/clientBot
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY bot/bot.go bot/bot.go
COPY Protobuf protobuf
COPY botMaster.go botMaster.go
WORKDIR /src/registerio/clientBot/protobuf/cvInterface
RUN go build cvInterface.pb.go
WORKDIR /src/registerio/clientBot/protobuf/rvInterface
RUN go build rvInterface.pb.go
WORKDIR /src/registerio/clientBot/protobuf/dbRequests
RUN go build dbRequests.pb.go
WORKDIR /src/registerio/clientBot/protobuf/token
RUN go build token.pb.go
WORKDIR /src/registerio/clientBot/bot
RUN go build bot.go
WORKDIR /src/registerio/clientBot/
RUN go build .
## Our start command which kicks off
## our newly created binary executable
FROM alpine
COPY --from=build /src/registerio/clientBot/clientBot /go/bin/clientBot
CMD ["/go/bin/clientBot"]
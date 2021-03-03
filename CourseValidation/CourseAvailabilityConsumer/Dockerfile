FROM golang:alpine AS build
WORKDIR /src/registerio/cv/consumer/

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY database/data.go database/data.go
COPY consumer.go . 
WORKDIR /src/registerio/cv/consumer/database
RUN go build data.go
WORKDIR /src/registerio/cv/consumer
RUN go build consumer.go

FROM alpine
WORKDIR /usr/bin
COPY --from=build /src/registerio/cv/consumer/consumer /go/bin/consumer
ENV AWS_ACCESS_KEY_ID=AKIA53XJENGN4KQU3KFO
ENV AWS_SECRET_ACCESS_KEY=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
ENV AWS_REGION=us-east-1

CMD ["/go/bin/consumer"]
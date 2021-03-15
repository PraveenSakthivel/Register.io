FROM golang:alpine AS build
WORKDIR /src/registerio/cv/main/

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY prereq/prereqInterface.pb.go prereq/prereqInterface.pb.go
COPY protobuf/cvInterface.pb.go protobuf/cvInterface.pb.go
COPY cv.go . 
WORKDIR /src/registerio/cv/main/prereq
RUN go build prereqInterface.pb.go
WORKDIR /src/registerio/cv/main/protobuf
RUN go build cvInterface.pb.go
WORKDIR /src/registerio/cv/main
RUN go build cv.go

FROM alpine
WORKDIR /usr/bin
COPY --from=build /src/registerio/cv/main/cv /go/bin/cv
ENV AWS_ACCESS_KEY_ID=AKIA53XJENGN4KQU3KFO
ENV AWS_SECRET_ACCESS_KEY=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
ENV AWS_REGION=us-east-1

EXPOSE 8080

CMD ["/go/bin/cv"]
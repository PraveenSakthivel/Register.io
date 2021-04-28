FROM golang:alpine AS build
WORKDIR /src/registerio/cv/main/

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY prereq prereq
COPY protobuf/cvInterface.pb.go protobuf/cvInterface.pb.go
COPY secrets/secret.go secrets/secret.go
COPY database/data.go database/data.go
COPY classTiming/classTiming.go classtiming/classtiming.go
COPY cv.go . 
WORKDIR /src/registerio/cv/main/prereq
RUN go build .
WORKDIR /src/registerio/cv/main/protobuf
RUN go build cvInterface.pb.go
WORKDIR /src/registerio/cv/main/secrets
RUN go build secret.go
WORKDIR /src/registerio/cv/main/classtiming
RUN go build classtiming.go
WORKDIR /src/registerio/cv/main/database
RUN go build data.go
WORKDIR /src/registerio/cv/main
RUN go build cv.go

FROM alpine
WORKDIR /usr/bin
COPY --from=build /src/registerio/cv/main/cv /go/bin/cv

EXPOSE 8080

CMD ["/go/bin/cv"]
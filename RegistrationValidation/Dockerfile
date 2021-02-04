FROM golang:alpine AS build
WORKDIR /src/registerio/rv

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY protobuf/rvInterface.pb.go protobuf/rvInterface.pb.go
COPY rv.go . 
WORKDIR /src/registerio/rv/protobuf
RUN go build rvInterface.pb.go
WORKDIR /src/registerio/rv
RUN go build rv.go

FROM alpine
WORKDIR /usr/bin
COPY --from=build /src/registerio/rv/rv /go/bin/rv

EXPOSE 8080

CMD ["/go/bin/rv"]






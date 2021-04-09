FROM golang:alpine AS build
WORKDIR /src/registerio/db

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY databse/data.go database/data.go
COPY secrets/secret.go secrets/secret.go
COPY models/soc.go models/soc.go
COPY pb/dbRequests.go pb/dbRequests.go
COPY pb/dbRequests_grpc.pb.go pb/dbRequests_grpc.pb.go
COPY pb/dbRequests.pb.go pb/dbRequests.pb.go
COPY main.go . 
WORKDIR /src/registerio/db/pb
RUN go build dbRequests.pb.go dbRequests_grpc.pb.go dbRequests.go
WORKDIR /src/registerio/db
RUN go build main.go

FROM alpine
WORKDIR /usr/bin
COPY --from=build /src/registerio/db/main /go/bin/db

ENV PORT=8080

EXPOSE 8080

CMD ["/go/bin/db"]






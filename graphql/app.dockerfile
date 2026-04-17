FROM golang:1.25-alpine AS build
RUN apk --no-cache add ca-certificates git
WORKDIR /go/src/go-grpc-elk-postgres-microservice
COPY go.mod go.sum ./
RUN go mod download
COPY account account
COPY product product
COPY order order
COPY graphql graphql
RUN CGO_ENABLED=0 go build -o /go/bin/app ./graphql

FROM alpine:3.21
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=build /go/bin .
EXPOSE 8080
CMD ["app"]

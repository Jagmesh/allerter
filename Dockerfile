FROM golang:1.23-alpine as builder
WORKDIR /build
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o /main main.go

FROM alpine:3
COPY --from=builder main /bin/main
ENTRYPOINT ["/bin/main"]
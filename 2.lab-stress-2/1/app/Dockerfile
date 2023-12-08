FROM golang:1.21.1-alpine3.18 as builder
WORKDIR /build
COPY . .
RUN go build -o application main.go

FROM alpine:3.18
RUN apk add --no-cache bash
COPY --from=builder /build/application /bin/application
WORKDIR /bin
ENTRYPOINT ["/bin/application"]

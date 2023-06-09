FROM golang:alpine AS builder

WORKDIR /build
COPY go.mod ./
COPY *.go ./

RUN go build -o nas-perf-test

FROM alpine:latest
WORKDIR /app
COPY --from=builder /build/nas-perf-test ./
RUN mkdir -p /nas-perf-test-folder

ENTRYPOINT ["/app/nas-perf-test"]

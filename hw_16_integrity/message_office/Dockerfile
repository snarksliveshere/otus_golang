FROM golang:1.13.1 as builder
LABEL maintainer="snarksliveshere"
RUN mkdir -p /opt/msg_service
WORKDIR /opt/msg_service
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /opt/service/msg_service

# Release
FROM alpine:latest
COPY --from=builder /opt/service/msg_service /bin/msg_service
ENTRYPOINT ["/bin/msg_service"]
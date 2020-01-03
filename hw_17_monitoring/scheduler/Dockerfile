FROM golang:1.13.1 as builder
LABEL maintainer="snarksliveshere"
RUN mkdir -p /opt/sched_service
WORKDIR /opt/sched_service
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /opt/service/sched_service

# Release
FROM alpine:latest
COPY --from=builder /opt/service/sched_service /bin/sched_service
ENTRYPOINT ["/bin/sched_service"]
FROM golang:1.13.1 as builder
LABEL maintainer="snarksliveshere"
RUN mkdir -p /opt/reg_service
WORKDIR /opt/reg_service
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /opt/service/reg_service

# Release
FROM alpine:latest
COPY --from=builder /opt/service/reg_service /bin/reg_service
COPY wait.sh /bin/wait.sh
ENV LOG_LEVEL = "error"
RUN ["/bin/wait.sh"]
#ENTRYPOINT ["/bin/reg_service"]

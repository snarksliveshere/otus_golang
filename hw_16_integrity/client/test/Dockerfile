FROM golang:1.13.1

RUN mkdir -p /opt/integration_tests
WORKDIR /opt/integration_tests

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
WORKDIR /opt/integration_tests/web
CMD ["go", "test"]


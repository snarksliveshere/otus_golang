FROM golang:latest as builder
LABEL maintainer="snarksliveshere"
RUN mkdir -p /opt/reg_service
WORKDIR /opt/reg_service
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /opt/service/reg_service
#ENV DB_USER="md"
#ENV DB_PASSWORD="secret"
#ENV DB_NAME="md_calendar"
#ENV DB_HOST=172.30.99.51
#ENV DB_PORT=5432
#ENV LOG_LEVEL="info"
#ENV WEB_PORT=8080
#ENV GRPC_PORT=50054
#ENV LISTEN_IP="localhost"
#RUN cd ./migrations && go run main.go
RUN CGO_ENABLED=0 cd ./migrations && go build -o /opt/service/migrations
#RUN cd /opt/service && ./migrations

# Release
FROM alpine:latest
COPY --from=builder /opt/service/migrations /bin/migrations
COPY --from=builder /opt/service/reg_service /bin/reg_service
RUN apk update && apk upgrade && apk add bash
#RUN cd /bin && ./migrations
#RUN chmod +x /bin/migrations
#RUN cd /bin && ./migrations up
#ENTRYPOINT ["/bin/migrations"]
EXPOSE 8888
CMD ["/bin/migrations"]
ENTRYPOINT ["/bin/reg_service"]
#cd $(CURDIR)/message_office && go build

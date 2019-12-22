FROM golang:alpine
RUN apk update && apk upgrade && apk add bash && apk add curl
RUN mkdir -p /opt/reg_service
WORKDIR /opt/reg_service
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /opt/service/reg_service
RUN CGO_ENABLED=0 GOOS=linux cd ./migrations && go build -o /opt/service/migrations

#EXPOSE 8888
#CMD ["/bin/migrations"]

#ENTRYPOINT ["/bin/reg_service"]


ENTRYPOINT ["/opt/service/reg_service"]
CMD ["/opt/service/migrations", "up"]

#RUN cd /bin && ./migrations
#RUN chmod +x /bin/migrations
#RUN cd /bin && ./migrations up
#ENTRYPOINT ["/bin/migrations"]

#cd $(CURDIR)/message_office && go build

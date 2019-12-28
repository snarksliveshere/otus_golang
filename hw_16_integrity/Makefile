up-all: docker-up migrations-up
up: docker-up

test-all: build-server run-server test

test:
	cd $(CURDIR)/client/grpc && go test

rabbit-up:
	docker run -d --name rb -p 15672:15672
docker-up:
	docker-compose up -d
down: docker-down

build-worker:
	cd $(CURDIR)/message_office && go build

run-worker:
	cd $(CURDIR)/message_office && ./message_office

init: build-worker build-server build-scheduler

migrations-up:
	cd $(CURDIR)/server/ && go run migrations/main.go

migrations-down:
	cd $(CURDIR)/server/migrations && ./migrations down

build-server:
	cd $(CURDIR)/server && go build

run-server:
	cd $(CURDIR)/server && ./server

build-scheduler:
	cd $(CURDIR)/scheduler && go build

run-scheduler:
	cd $(CURDIR)/scheduler && ./scheduler

docker-down:
	docker-compose down --remove-orphans

docker-down-clear:
	docker-compose down -v --remove-orphans

docker-pull:
	docker-compose pull

docker-build:
	docker-compose build

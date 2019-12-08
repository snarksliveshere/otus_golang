up: docker-up rabbit-up

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

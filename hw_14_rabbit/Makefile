up: docker-up rabbit-up
rabbit-up:
	docker run -d --name rb -p 15672:15672
docker-up:
	docker-compose up -d
down: docker-down
init: docker-down-clear docker-pull docker-build docker-up






docker-down:
	docker-compose down --remove-orphans

docker-down-clear:
	docker-compose down -v --remove-orphans

docker-pull:
	docker-compose pull

docker-build:
	docker-compose build

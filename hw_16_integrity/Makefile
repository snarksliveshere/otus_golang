.PHONY: up up-test down restart test

up-test:
	docker-compose -f docker-compose.test.yml up --build -d ;

up:
	docker-compose up -d --build

rebuild-test:
	docker-compose -f docker-compose.test.yml up --build --no-deps -d client_server

restart-test:
	docker-compose -f docker-compose.test.yml restart client_server

down:
	docker-compose down

restart: down up

restart-test: down up-test

test:
	set -e ;\
	docker-compose -f docker-compose.test.yml up --build -d ;\
	test_status_code=0 ;\
	docker-compose -f docker-compose.test.yml run client_server go test || test_status_code=$$? ;\
	docker-compose -f docker-compose.test.yml down ;\
	exit $$test_status_code ;\
	echo $$test_status_code ;

example:
	docker-compose -f docker-compose.test.yml up --build -d
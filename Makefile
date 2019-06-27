default:
	@echo "=============Building Local Service============="
	docker build -f Service/Dockerfile -t ocsen-golang-service .
	docker build -f Consumer/Dockerfile -t ocsen-golang-consumer .

build: default
	@echo "=============Build And Starting Service Locally============="
	docker-compose up -d

up:
	@echo "=============Starting Service Locally============="
	docker-compose up -d

logs:
	docker-compose logs -f

down:
	docker-compose down

test:
	go test -v -cover ./...

clean: down
	@echo "=============Cleaning Up============="
	rm -f ocsen-service
	docker system prune -f
	docker volume prune -f

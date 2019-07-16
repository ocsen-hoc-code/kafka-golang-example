ocsen-kafka:
	@echo "=============Building OcSen-Kafka Image============="
	docker build -f KafkaDocker/Dockerfile -t ocsen-golang-kafka .
	@echo "-----------Building Ocsen Kafka Completed-----------"
default:
	@echo "=============Building Local Service============="
	@echo "-------------Building Ocsen Service-------------"
	docker build -f Service/Dockerfile -t ocsen-golang-service .
	@echo "--------Building Ocsen Service Completed--------"
	@echo "-------------Building Ocsen Consumer------------"
	docker build -f Consumer/Dockerfile -t ocsen-golang-consumer .
	@echo "--------Building Ocsen Consumer Completed-------"

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
	rm -f ocsen-golang-service
	rm -f ocsen-golang-consumer
	docker rmi ocsen-golang-service
	docker rmi ocsen-golang-consumer
	docker system prune -f
	docker volume prune -f

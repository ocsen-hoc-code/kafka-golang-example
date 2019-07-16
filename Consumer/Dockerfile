FROM ocsen-golang-kafka:latest

ADD . /src/consumer
WORKDIR /src/consumer

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
ENTRYPOINT CompileDaemon -log-prefix=false -build="go build -mod=vendor main.go" -command="./main"

EXPOSE 8080

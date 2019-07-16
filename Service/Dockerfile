FROM golang:1.12

ADD . /src/consumer
WORKDIR /src/consumer

ENV CGO_ENABLED=0
RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
ENTRYPOINT CompileDaemon -log-prefix=false -build="go build -mod=vendor main.go" -command="./main"

EXPOSE 8080

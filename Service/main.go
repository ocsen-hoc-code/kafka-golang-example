package main

import (
	"fmt"
	kafkaService "github.com/ocsen-hoc-code/kafka-go/kafka-service"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type Person struct {
	Id   int
	Name string
}

func main() {
	fmt.Printf("Service Starting!\n")

	p := Person{
		Id:   222,
		Name: "Binh Minh",
	}

	service := kafkaService.NewService(kafka.ConfigMap{
		"bootstrap.servers": "172.80.0.3:9092",
		"group.id":          "myGroup",
		"sasl.username":     "ocsen_broker",
		"sasl.password":     "ocsen_broker_password",
		"auto.offset.reset": "earliest",
	})
	done := make(chan bool)
	go service.Producer.Report(done)
	service.Producer.SendMessage("ocsen-hoc-code", p)
	<-done
	service.Producer.Close()
}

package main

import (
	"fmt"
	kafkaService "github.com/ocsen-hoc-code/kafka-go/kafka-service"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"time"
)

func myConsumer(msg *kafka.Message, args ...interface{}) {
	fmt.Printf("My Comsumer: %s: %s\n", msg.TopicPartition, string(msg.Value))
}

func main() {
	service := kafkaService.NewService(kafka.ConfigMap{
		"bootstrap.servers": "172.80.0.3:9092",
		"group.id":          "myGroup",
		"sasl.username":     "ocsen_broker",
		"sasl.password":     "ocsen_broker_password",
		"auto.offset.reset": "earliest",
	})

	if !service.TopicExisted("ocsen-hoc-code", 50000) {
		service.CreateTopic(nil, "ocsen-hoc-code", 0, 1, time.Duration(10)*time.Second)
	}

	service.Consumer.SubscribeTopics([]string{"ocsen-hoc-code"}, nil)
	service.Consumer.ConsumerWorker(myConsumer)
	service.Consumer.Close()
}

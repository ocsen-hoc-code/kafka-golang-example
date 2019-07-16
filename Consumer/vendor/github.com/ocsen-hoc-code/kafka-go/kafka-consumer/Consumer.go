package Consumer

import (
	"encoding/json"
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type Consumer struct {
	kafkaConsumer *kafka.Consumer
}
type worker func(msg *kafka.Message, args ...interface{})

func NewConsumer(kafkaConsumer *kafka.Consumer) *Consumer {
	consumer := Consumer{
		kafkaConsumer: kafkaConsumer,
	}

	return &consumer
}

func ConvertByteToObject(b []byte, object interface{}) error {
	return json.Unmarshal(b, object)
}

func (this *Consumer) SubscribeTopics(topics []string, rebalanceCb kafka.RebalanceCb) error {
	return this.kafkaConsumer.SubscribeTopics(topics, rebalanceCb)
}

func (this *Consumer) ConsumerWorker(w worker, args ...interface{}) {
	for {
		msg, err := this.kafkaConsumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			w(msg, args)
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
func (this *Consumer) Close() {
	this.kafkaConsumer.Close()
}

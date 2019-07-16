package Producer

import (
	"encoding/json"
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type Producer struct {
	kafkaProducer *kafka.Producer
}

func NewProducer(kafkaProducer *kafka.Producer) *Producer {
	producer := Producer{
		kafkaProducer: kafkaProducer,
	}

	return &producer
}

func ConvertObjectToByte(object interface{}) ([]byte, error) {
	b, err := json.Marshal(object)
	return b, err
}

func (this *Producer) Report(done chan bool) {
	defer close(done)
	for e := range this.kafkaProducer.Events() {
		switch ev := e.(type) {
		case *kafka.Message:
			m := ev
			if m.TopicPartition.Error != nil {
				fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
			} else {
				fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
					*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
			}
			return

		default:
			fmt.Printf("Ignored event: %s\n", ev)
		}
	}
}

func (this *Producer) SendMessage(topic string, data interface{}) (kafka.Event, error) {
	eventDelivery := make(chan kafka.Event)
	b, err := ConvertObjectToByte(data)
	if nil != err {
		fmt.Printf("Can't convert data to byte")
		return nil, err
	}

	err = this.kafkaProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          b,
	}, eventDelivery)

	return <-eventDelivery, err
}

func (this *Producer) Close() {
	this.kafkaProducer.Close()
}

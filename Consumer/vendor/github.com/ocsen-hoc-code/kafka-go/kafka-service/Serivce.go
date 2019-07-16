package Service

import (
	"context"
	"fmt"
	"time"

	ocsenConsumer "github.com/ocsen-hoc-code/kafka-go/kafka-consumer"
	ocsenProducer "github.com/ocsen-hoc-code/kafka-go/kafka-producer"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type KafkaService struct {
	KafkaConfig *kafka.ConfigMap
	KafkaAdmin  *kafka.AdminClient
	Producer    *ocsenProducer.Producer
	Consumer    *ocsenConsumer.Consumer
}

func NewService(kafkaConfig kafka.ConfigMap) *KafkaService {
	kafkaService := new(KafkaService)
	kafkaService.KafkaConfig = &kafkaConfig

	if c, err := createConsumer(kafkaService.KafkaConfig); nil == err {
		kafkaService.Consumer = ocsenConsumer.NewConsumer(c)
	} else {
		fmt.Println("Error: %s", err.Error())
	}

	if p, err := createProducer(kafkaService.KafkaConfig); nil == err {
		kafkaService.Producer = ocsenProducer.NewProducer(p)
	} else {
		fmt.Println("Error: %s", err.Error())
	}

	if adminClient, err := AdminClientInstance(kafkaService.KafkaConfig); nil == err {
		kafkaService.KafkaAdmin = adminClient
	} else {
		fmt.Println("Error: %s", err.Error())
	}

	return kafkaService
}

func AdminClientInstance(kafkaConfig *kafka.ConfigMap) (*kafka.AdminClient, error) {
	return kafka.NewAdminClient(kafkaConfig)
}

func (this *KafkaService) TopicExisted(topic string, timeOut int) bool {
	metaTopic, _ := this.KafkaAdmin.GetMetadata(&topic, false, timeOut)
	return nil != metaTopic
}

func (this *KafkaService) CreateTopic(ctx context.Context, topic string, numParts int, replicationFactor int, maxDur time.Duration) bool {
	_, err := this.KafkaAdmin.CreateTopics(
		ctx,
		[]kafka.TopicSpecification{{
			Topic:             topic,
			NumPartitions:     numParts,
			ReplicationFactor: replicationFactor}},
		kafka.SetAdminOperationTimeout(maxDur))
	if err != nil {
		fmt.Printf("Failed to create topic: %v\n", err)
		return false
	}
	return true
}

func (this *KafkaService) DeleteTopic(ctx context.Context, topic []string) bool {
	_, err := this.KafkaAdmin.DeleteTopics(ctx, topic)
	if err != nil {
		fmt.Printf("Failed to delete topic: %v\n", err)
		return false
	}
	return true
}

func createConsumer(config *kafka.ConfigMap) (*kafka.Consumer, error) {
	return kafka.NewConsumer(config)
}

func createProducer(config *kafka.ConfigMap) (*kafka.Producer, error) {
	return kafka.NewProducer(config)
}

func (this *KafkaService) CreateConsumer() (*kafka.Consumer, error) {
	return createConsumer(this.KafkaConfig)
}

func (this *KafkaService) CreateProducer() (*kafka.Producer, error) {
	return createProducer(this.KafkaConfig)
}

package KafkaClient

import (
	"context"
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"time"
)
type KafkaClient struct{
	KafkaConfig kafka.ConfigMap
	KafkaAdmin kafka.AdminClient
}

func NewKafkaClient(kafkaConfig kafka.ConfigMap) *KafkaClient {
	kafkaClient := new(KafkaClient)
	kafkaClient.KafkaConfig = kafkaConfig

	if adminClient, err := AdminClientInstance(kafkaClient.KafkaConfig); nil == err {
		kafkaClient.KafkaAdmin = *adminClient
	}else {

		fmt.Println("Error: %s" , err.Error())
	}

	return kafkaClient
}

func AdminClientInstance(kafkaConfig kafka.ConfigMap) (*kafka.AdminClient, error){
	return kafka.NewAdminClient(&kafkaConfig)
}

func (this *KafkaClient)TopicExisted(topic string, timeOut int) bool {
	metaTopic, _:=	this.KafkaAdmin.GetMetadata(&topic,false, timeOut)
	return nil != metaTopic
}

func (this *KafkaClient)CreateTopic(ctx context.Context,topic string, numParts int,  replicationFactor int, maxDur time.Duration) bool{
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

func (this *KafkaClient)DeleteTopic(ctx context.Context,topic[] string) bool{
	_, err := this.KafkaAdmin.DeleteTopics(ctx, topic)
	if err != nil {
		fmt.Printf("Failed to delete topic: %v\n", err)
		return false
	}
	return true
}

func (this *KafkaClient)CreateConsumer() (*kafka.Consumer, error){
	return kafka.NewConsumer(&this.KafkaConfig)
}

func (this *KafkaClient)CreateProducer() (*kafka.Producer, error){
	return kafka.NewProducer(&this.KafkaConfig)
}
package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	produce([]byte("message from confluent kafka producer"), "my-topic")
	fmt.Println("Message Sent")
}

func produce(msg []byte, topic string) {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	}
	kafkaProducer, err := kafka.NewProducer(configMap)
	if err != nil {
		panic(err)
	}
	defer kafkaProducer.Close()
	kafkaProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          msg,
	}, nil)
}

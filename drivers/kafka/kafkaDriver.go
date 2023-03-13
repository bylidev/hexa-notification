package drivers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/igloar96/hexa-notification/core/useCases"
	"github.com/segmentio/kafka-go"
)

type KafkaDriver struct {
	host                      string
	port                      uint16
	topic                     string
	createNotificationUseCase useCases.CreateNotificationUseCase
}

func NewKafkaDriver(host string, port uint16, topic string, createNotificationUseCase *useCases.CreateNotificationUseCase) *KafkaDriver {
	return &KafkaDriver{host: host, port: port, topic: topic, createNotificationUseCase: *createNotificationUseCase}
}

func (s *KafkaDriver) Excecute() {
	brokerAddresses := []string{fmt.Sprintf("%s:%d", s.host, s.port)}
	topic := s.topic
	partition := 0
	readerConfig := kafka.ReaderConfig{
		Brokers:         brokerAddresses,
		Topic:           topic,
		Partition:       partition,
		MinBytes:        10e3,
		MaxBytes:        10e6,
		MaxWait:         1 * time.Second,
		ReadLagInterval: -1,
	}
	reader := kafka.NewReader(readerConfig)
	defer reader.Close()

	// start reading messages
	for {
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error while reading message: %v", err)
			continue
		}
		fmt.Printf("Received message: key=%s value=%s\n", string(message.Key), string(message.Value))
	}
}

package adapters

import (
	"github.com/igloar96/hexa-notification/core/domain"
	"github.com/segmentio/kafka-go"
)

type KafkaMessageAdapter struct {
}

func NewKafkaMessageAdapter() *KafkaMessageAdapter {
	return &KafkaMessageAdapter{}
}

func (s *KafkaMessageAdapter) GetMessage(k *kafka.Message) (*domain.Message, error) {
	return &domain.Message{Text: string(k.Value)}, nil
}

package adapters

import (
	"github.com/igloar96/hexa-notification/core/domain"
	"github.com/segmentio/kafka-go"
)

type KafkaMessageAdapter struct {
	kmsg *kafka.Message
}

func NewKafkaMessageAdapter(kmsg *kafka.Message) *KafkaMessageAdapter {
	return &KafkaMessageAdapter{kmsg: kmsg}
}

func (s *KafkaMessageAdapter) GetMessage() (*domain.Message, error) {
	return &domain.Message{Text: string(s.kmsg.Value)}, nil
}

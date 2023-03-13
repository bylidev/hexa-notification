package ports

import (
	"github.com/igloar96/hexa-notification/core/domain"
)

type KafkaMessageAdapter struct {
}

func (s *KafkaMessageAdapter) GetMessage() (*domain.Message, error) {
	return nil, nil
}

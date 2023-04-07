package adapters

import (
	"github.com/igloar96/hexa-notification/core/domain"
)

type GreetingMessageAdapter struct {
	Text string
}

func NewGreetingMessageAdapter(text string) *GreetingMessageAdapter {
	return &GreetingMessageAdapter{Text: text}
}

func (s *GreetingMessageAdapter) GetMessage() (*domain.Message, error) {
	return &domain.Message{Text: s.Text}, nil
}

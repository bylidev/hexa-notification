package driven

import (
	"fmt"
	"time"

	"github.com/igloar96/hexa-notification/core/domain"
)

type ConsoleAdapter struct {
}

func NewConsoleAdapter() *ConsoleAdapter {
	return &ConsoleAdapter{}
}

func (s *ConsoleAdapter) Notificate(message *domain.Message) (*domain.Notification, error) {
	fmt.Printf("[ConsoleAdapter] message: %s", message.Text)
	return &domain.Notification{
		SendedAt: time.Now(), Details: make([]byte, 0)}, nil
}

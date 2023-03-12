package ports

import (
	"github.com/igloar96/hexa-notification/core/domain"
)

type MessageAdapter interface {
	GetMessage() (*domain.Message, error)
}

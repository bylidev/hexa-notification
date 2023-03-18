package ports

import (
	"github.com/igloar96/hexa-notification/core/domain"
)

type NotificationDriverAdapter interface {
	GetMessage() (*domain.Message, error)
}

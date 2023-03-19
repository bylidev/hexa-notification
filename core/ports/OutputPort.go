package ports

import (
	"github.com/igloar96/hexa-notification/core/domain"
)

type OutputPort interface {
	Notificate(message *domain.Message) (*domain.Notification, error)
}

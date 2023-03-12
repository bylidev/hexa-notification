package ports

import (
	"github.com/igloar96/hexa-notification/core/domain"
)

type NotificationAdapter interface {
	Notificate(message *domain.Message) (*domain.Notification, error)
}

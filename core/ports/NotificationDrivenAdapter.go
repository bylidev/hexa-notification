package ports

import (
	"github.com/igloar96/hexa-notification/core/domain"
)

type NotificationDrivenAdapter interface {
	Notificate(message *domain.Message) (*domain.Notification, error)
}

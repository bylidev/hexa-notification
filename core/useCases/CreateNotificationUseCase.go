package useCases

import "github.com/igloar96/hexa-notification/core/ports"

// Command pattern
type CreateNotificationUseCase interface {
	Excecute(driver ports.NotificationDriverAdapter) []error
}

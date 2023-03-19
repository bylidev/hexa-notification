package useCases

import "github.com/igloar96/hexa-notification/core/ports"

// Command pattern
type CreateNotificationUseCase interface {
	Execute(driver ports.InputPort) []error
}

package useCases

import "github.com/igloar96/hexa-notification/core/domain"

// Command pattern
type CreateNotificationUseCase interface {
	Excecute(message *domain.Message) error
}

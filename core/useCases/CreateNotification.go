package useCases

import (
	domain "github.com/igloar96/hexa-notification/core/domain"
	outputPort "github.com/igloar96/hexa-notification/core/ports/outbound"
)

type CreateNotification struct {
	notificationAdapter outputPort.NotificationAdapter
}

func NewCreateNotification(outboundAdapter *outputPort.NotificationAdapter) *CreateNotification {
	return &CreateNotification{
		notificationAdapter: *outboundAdapter,
	}
}

func (s *CreateNotification) Excecute(message *domain.Message) error {
	_, err := s.notificationAdapter.Notificate(message)

	return err
}

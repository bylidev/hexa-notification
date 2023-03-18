package useCases

import (
	"errors"

	domain "github.com/igloar96/hexa-notification/core/domain"
	outputPort "github.com/igloar96/hexa-notification/core/ports"
)

type CreateNotification struct {
	notificationAdapterList *[]outputPort.NotificationDrivenAdapter
}

func NewCreateNotification(outboundAdapters *[]outputPort.NotificationDrivenAdapter) *CreateNotification {
	return &CreateNotification{
		notificationAdapterList: outboundAdapters,
	}
}

func (s *CreateNotification) Excecute(message *domain.Message) []error {
	var errList []error

	if message.Text == "" {
		return append(errList, errors.New("text is required"))
	}

	for _, adapter := range *s.notificationAdapterList {
		_, err := adapter.Notificate(message)
		if err != nil {
			errList = append(errList, err)
		}
	}

	return errList
}

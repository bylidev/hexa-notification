package test

import (
	"errors"
	"testing"
	"time"

	"github.com/igloar96/hexa-notification/core/domain"
	"github.com/igloar96/hexa-notification/core/ports"
	"github.com/igloar96/hexa-notification/core/useCases"
)

type MockNotifierAdapter struct {
	ErrorMsg string
}

func (n *MockNotifierAdapter) Notificate(message *domain.Message) (*domain.Notification, error) {
	if n.ErrorMsg != "" {
		return nil, errors.New(n.ErrorMsg)
	}
	return &domain.Notification{
		SendedAt: time.Now(), Details: make([]byte, 0)}, nil
}

type MockMessageAdapter struct {
	Text string
}

func (n *MockMessageAdapter) GetMessage() (*domain.Message, error) {
	return &domain.Message{Text: n.Text}, nil
}

func TestCreateNotification(t *testing.T) {
	t.Run("CreateNotification", func(t *testing.T) {
		t.Log("CreateNotification expected to implements CreateNotificationUseCase")
		var _ useCases.CreateNotificationUseCase = (*useCases.CreateNotification)(nil)
	})
}

func TestCreateNotificationExecute(t *testing.T) {

	t.Run("TestCreateNotificationExecute_1", func(t *testing.T) {
		t.Log("Expected to adapt Message request body correctly.")
		//arrange
		msg := &MockMessageAdapter{Text: "byli.dev !"}
		var notificationOutputPort []ports.OutputPort
		notificationOutputPort = append(notificationOutputPort, &MockNotifierAdapter{})
		useCase := useCases.NewCreateNotification(&notificationOutputPort)

		//act
		err := useCase.Execute(msg)
		//assert

		if len(err) >= 1 {
			t.Errorf("Expected to adapt Message request body correctly but got error: %s", err)
		}
	})
	t.Run("TestCreateNotificationExecute_2", func(t *testing.T) {
		t.Log("Expected to return error if Message text is empty.")
		//arrange
		msg := &MockMessageAdapter{Text: ""}
		var notificationOutputPort []ports.OutputPort
		notificationOutputPort = append(notificationOutputPort, &MockNotifierAdapter{})
		useCase := useCases.NewCreateNotification(&notificationOutputPort)

		//act
		e := useCase.Execute(msg)
		//assert
		if len(e) == 0 || e[0].Error() != "text is required" {
			t.Errorf("Expected to return error if Message text is empty.")
		}

	})
	t.Run("TestCreateNotificationExecute_3", func(t *testing.T) {
		t.Log("Expected to return error if output adapter has an error.")
		//arrange
		msg := &MockMessageAdapter{Text: "byli.dev!"}
		var notificationOutputPort []ports.OutputPort
		mock := &MockNotifierAdapter{ErrorMsg: "Error inesperado"}
		notificationOutputPort = append(notificationOutputPort, mock)
		useCase := useCases.NewCreateNotification(&notificationOutputPort)

		//act
		e := useCase.Execute(msg)
		//assert
		if len(e) == 0 || e[0].Error() != mock.ErrorMsg {
			t.Errorf("Expected to return error if Message text is empty.")
		}

	})
}

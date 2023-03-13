package useCase

import (
	"errors"
	"testing"
	"time"

	"github.com/igloar96/hexa-notification/core/domain"
	"github.com/igloar96/hexa-notification/core/useCases"
)

type MockNotifier struct {
	ErrorMsg string
}

func (n *MockNotifier) Notificate(message *domain.Message) (*domain.Notification, error) {
	if n.ErrorMsg != "" {
		return nil, errors.New(n.ErrorMsg)
	}
	return &domain.Notification{
		SendedAt: time.Now(), Details: make([]byte, 0)}, nil
}

func TestCreateNotification(t *testing.T) {
	t.Run("CreateNotification", func(t *testing.T) {
		t.Log("CreateNotification expected to implements CreateNotificationUseCase")
		var _ useCases.CreateNotificationUseCase = (*useCases.CreateNotification)(nil)
	})
}

func TestCreateNotificationExcecute(t *testing.T) {

	t.Run("TestCreateNotificationExcecute_1", func(t *testing.T) {
		t.Log("Expected to adapt Message request body correctly.")
		//arrange
		msg := &domain.Message{Text: "Hi"}
		useCase := useCases.NewCreateNotification(&MockNotifier{})

		//act
		err := useCase.Excecute(msg)
		//assert

		if err != nil {
			t.Errorf("Expected to adapt Message request body correctly but got error: %s", err)
		}
	})
	t.Run("TestCreateNotificationExcecute_2", func(t *testing.T) {
		t.Log("Expected to return error if Message text is empty.")
		//arrange
		msg := &domain.Message{Text: ""}
		useCase := useCases.NewCreateNotification(&MockNotifier{})

		//act
		err := useCase.Excecute(msg)
		//assert

		if err == nil || err.Error() != "text is required" {
			t.Errorf("Expected to return error if Message text is empty.")
		}
	})
	t.Run("TestCreateNotificationExcecute_3", func(t *testing.T) {
		t.Log("Expected to return error if output adapter has an error.")
		//arrange
		msg := &domain.Message{Text: "hi !"}
		mock := &MockNotifier{ErrorMsg: "Error inesperado"}
		useCase := useCases.NewCreateNotification(mock)

		//act
		err := useCase.Excecute(msg)
		//assert

		if err == nil || err.Error() != mock.ErrorMsg {
			t.Errorf("Expected to return error if Message text is empty.")
		}
	})
}

package ports

import (
	"time"

	"fmt"
	"io"
	"net/http"

	"github.com/igloar96/hexa-notification/core/domain"
)

type TelegramNotificationAdapter struct {
	apiKey string
	chatId string
}

func NewTelegramNotificationAdapter(apiKey string, chatId string) *TelegramNotificationAdapter {
	return &TelegramNotificationAdapter{apiKey: apiKey, chatId: chatId}
}

func (s *TelegramNotificationAdapter) Notificate(message *domain.Message) (*domain.Notification, error) {
	var err error
	resp, httpError := http.Get(fmt.Sprintf(`https://api.telegram.org/bot%s/sendMessage?chat_id=%s&parse_mode=html&text=%s`,
		s.apiKey, s.chatId, message.Text))

	if err != nil {
		err = httpError
	}

	defer resp.Body.Close()

	body, ioError := io.ReadAll(resp.Body)

	if err != nil {
		err = ioError
	}

	return &domain.Notification{SendedAt: time.Now(), Details: body}, err
}

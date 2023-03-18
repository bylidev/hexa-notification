package handlers

import (
	"github.com/gin-gonic/gin"
	useCases "github.com/igloar96/hexa-notification/core/useCases"
	"github.com/igloar96/hexa-notification/drivers/gin/adapters"
)

type NotificationHandler struct {
	createNotificationUseCase useCases.CreateNotificationUseCase
}

func NewNotificationHandler(createNotificationUseCase *useCases.CreateNotificationUseCase) *NotificationHandler {
	return &NotificationHandler{createNotificationUseCase: *createNotificationUseCase}
}

func (s *NotificationHandler) Create(ctx *gin.Context) {
	var msg, _ = adapters.NewGinContextAdapter(ctx).GetMessage()
	s.createNotificationUseCase.Excecute(msg)
}

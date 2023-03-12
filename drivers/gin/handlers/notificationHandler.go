package handlers

import (
	"github.com/gin-gonic/gin"
	inbound "github.com/igloar96/hexa-notification/core/ports/inbound"
	useCases "github.com/igloar96/hexa-notification/core/useCases"
)

type NotificationHandler struct {
	createNotificationUseCase useCases.CreateNotificationUseCase
}

func NewNotificationHandler(createNotificationUseCase *useCases.CreateNotificationUseCase) *NotificationHandler {
	return &NotificationHandler{createNotificationUseCase: *createNotificationUseCase}
}

func (s *NotificationHandler) Create(ctx *gin.Context) {
	var msg, _ = inbound.NewGinContextAdapter(ctx).GetMessage()
	s.createNotificationUseCase.Excecute(msg)
}

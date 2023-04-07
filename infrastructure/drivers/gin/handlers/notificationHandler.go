package handlers

import (
	"github.com/gin-gonic/gin"
	useCases "github.com/igloar96/hexa-notification/core/useCases"
	"github.com/igloar96/hexa-notification/infrastructure/drivers/gin/adapters"
)

type NotificationHandler struct {
	createNotificationUseCase useCases.UseCase
}

func NewNotificationHandler(createNotificationUseCase *useCases.UseCase) *NotificationHandler {
	return &NotificationHandler{createNotificationUseCase: *createNotificationUseCase}
}

func (s *NotificationHandler) Create(ctx *gin.Context) {
	s.createNotificationUseCase.Execute(adapters.NewGinContextAdapter(ctx))
}

package ports

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/igloar96/hexa-notification/core/domain"
)

type GinContextAdapter struct {
	ctx *gin.Context
}

func NewGinContextAdapter(ctx *gin.Context) *GinContextAdapter {
	return &GinContextAdapter{ctx: ctx}
}

func (s *GinContextAdapter) GetMessage() (*domain.Message, error) {
	var message domain.Message

	if err := s.ctx.ShouldBindJSON(&message); err != nil {
		return nil, err
	}

	if message.Text == "" {
		return nil, errors.New("text cannot be empty")
	}

	return &message, nil
}

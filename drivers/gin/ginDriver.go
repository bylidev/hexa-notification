package drivers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/igloar96/hexa-notification/core/useCases"
	"github.com/igloar96/hexa-notification/drivers/gin/handlers"
)

type GinDriver struct {
	host                      string
	port                      uint16
	createNotificationUseCase useCases.CreateNotificationUseCase
}

func NewGinDriver(host string, port uint16, createNotificationUseCase *useCases.CreateNotificationUseCase) *GinDriver {
	return &GinDriver{host: host, port: port, createNotificationUseCase: *createNotificationUseCase}
}

func (s *GinDriver) Excecute() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.POST("/", handlers.NewNotificationHandler(&s.createNotificationUseCase).Create)
	router.Run(fmt.Sprintf("%s:%d", s.host, s.port))
}

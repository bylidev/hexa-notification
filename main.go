package main

import (
	ports "github.com/igloar96/hexa-notification/core/ports/outbound"
	useCases "github.com/igloar96/hexa-notification/core/useCases"
	ginDriver "github.com/igloar96/hexa-notification/drivers/gin"
)

func main() {
	var notificationOutputPort ports.NotificationAdapter = ports.NewTelegramNotificationAdapter(`5780441884:AAETjXLPww4DNwm-jV5P-Lmq1H1PaDad1mc`, `5711869588`)
	//var notificationOutputPort ports.NotificationAdapter = ports.NewConsoleAdapter()
	var createNotificationUseCase useCases.CreateNotificationUseCase = useCases.NewCreateNotification(&notificationOutputPort)

	//Drivers
	ginDriver.NewGinDriver(`localhost`, 8080, &createNotificationUseCase).Excecute()
}

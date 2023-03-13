package main

import (
	ports "github.com/igloar96/hexa-notification/core/ports/outbound"
	useCases "github.com/igloar96/hexa-notification/core/useCases"
	ginDriver "github.com/igloar96/hexa-notification/drivers/gin"
	kafkaDriver "github.com/igloar96/hexa-notification/drivers/kafka"
)

func main() {
	var notificationOutputPort ports.NotificationAdapter = ports.NewTelegramNotificationAdapter(``, `5711869588`)
	var createNotificationUseCase useCases.CreateNotificationUseCase = useCases.NewCreateNotification(*&notificationOutputPort)

	//Drivers
	ginDriver.NewGinDriver(`localhost`, 8080, &createNotificationUseCase).Excecute()
	kafkaDriver.NewKafkaDriver(`localhost`, 9092, `dev.byli.events.notifications`, &createNotificationUseCase).Excecute()
}

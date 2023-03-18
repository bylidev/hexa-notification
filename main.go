package main

import (
	"flag"
	"fmt"
	"sync"
	"time"

	"github.com/igloar96/hexa-notification/core/domain"
	"github.com/igloar96/hexa-notification/core/ports"
	useCases "github.com/igloar96/hexa-notification/core/useCases"
	"github.com/igloar96/hexa-notification/driven"
	"github.com/igloar96/hexa-notification/drivers/gin"
	"github.com/igloar96/hexa-notification/drivers/kafka"
)

var (
	TELEGRAM_API_KEY = flag.String("TELEGRAM-API-KEY", "", "TELEGRAM_API_KEY")
	TELEGRAM_CHAT_ID = flag.String("TELEGRAM-CHAT-ID", "", "TELEGRAM_CHAT_ID")
	HTTP_SERVER_HOST = flag.String("HTTP-SERVER-HOST", "localhost", "HTTP_SERVER_HOST")
	HTTP_SERVER_PORT = flag.Uint("HTTP-SERVER-PORT", 0, "port")
	KAFKA_HOST       = flag.String("KAFKA-HOST", "localhost", "KAFKA_HOST")
	KAFKA_PORT       = flag.Uint("KAFKA-PORT", 9092, "KAFKA_PORT")
	KAFKA_TOPIC      = flag.String("KAFKA-TOPIC", "dev.byli.events.notifications", "KAFKA_TOPIC")
)

func main() {
	flag.Parse()

	//goroutines
	var wg sync.WaitGroup
	wg.Add(2)

	//out ports
	var outputPorts = []ports.NotificationDrivenAdapter{
		//`5780441884:AAFq99d4bDkmplRw88nopii_prC1He_Pwig`, `5711869588`),
		driven.NewTelegramNotificationAdapter(*TELEGRAM_API_KEY, *TELEGRAM_CHAT_ID),
		driven.NewConsoleAdapter(),
	}
	//use-case
	var createNotificationUseCase useCases.CreateNotificationUseCase = useCases.NewCreateNotification(&outputPorts)

	//in ports
	go func() {
		gin.NewGinDriver(*HTTP_SERVER_HOST, uint16(*HTTP_SERVER_PORT), &createNotificationUseCase).Excecute()
	}()
	go func() {
		kafka.NewKafkaDriver(*KAFKA_HOST, uint16(*KAFKA_PORT), *KAFKA_TOPIC, &createNotificationUseCase).Excecute()
	}()

	//Greeting
	createNotificationUseCase.Excecute(&domain.Message{
		Text: fmt.Sprintf("Hello! I woke up at %s on %s. You can trigger me via http api on port: %d", time.Now().Format("3:04 PM"), time.Now().Format("January 2, 2006"), *HTTP_SERVER_PORT),
	})

	wg.Wait()
}

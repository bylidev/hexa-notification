
# Hexa Notification

  

## Description

  

Hexa Notification is a project focused on providing a customizable and easy-to-use notification system. The solution is based on a modular and scalable architecture, allowing for easy integration into different applications and platforms. This project is brought to you by [byli.dev](https://byli.dev), where you can find more information about software architectures and design patterns.

  

## Features

  

- Modular and scalable architecture

- Customizable notifications

- Easy integration into different applications and platforms

**Support for multiple notification channels:**

-  **Input ports**: 
- - HTTP support using the Gin framework 
- - Kafka support using the kafka-go library
-  **Output ports**:
- -  Telegram
- -  Console 

  

## Requirements

  

- Go

  

## Installation and Configuration

  

1. Clone the repository to your local machine:

  

`git clone https://github.com/igloar96/hexa-notification.git`

  

2. Navigate to the project directory:

  

`cd hexa-notification`

  

3. Compile the project using Go:

  

`go build`

  

4. Configure environment variables in a configuration file, if necessary.

5. Run the compiled binary:

> Replace <TELEGRAM_API_KEY>, <TELEGRAM_CHAT_ID>, <HTTP_SERVER_HOST>,
> <HTTP_SERVER_PORT>, <KAFKA_HOST>, <KAFKA_PORT>, and <KAFKA_TOPIC> with
> your actual values.

```
./hexa-notification -TELEGRAM-API-KEY=<TELEGRAM_API_KEY> -TELEGRAM-CHAT-ID=<TELEGRAM_CHAT_ID> -HTTP-SERVER-HOST=<HTTP_SERVER_HOST> -HTTP-SERVER-PORT=<HTTP_SERVER_PORT> -KAFKA-HOST=<KAFKA_HOST> -KAFKA-PORT=<KAFKA_PORT> -KAFKA-TOPIC=<KAFKA_TOPIC>

```

**Example**

```

go run main.go -TELEGRAM-API-KEY=1234567890:ABCDEFGHIJKLMNOPQRSTUVXYZ -TELEGRAM-CHAT-ID=0987654321 -HTTP-SERVER-HOST=localhost -HTTP-SERVER-PORT=8080 -KAFKA-HOST=localhost -KAFKA-PORT=9092 -KAFKA-TOPIC=dev.byli.events.notifications

  

```

  

The project should now be up and running and ready to be integrated into your application.

  

## Contributions

  

Contributions are always welcome. Please make sure to read the contribution guidelines before submitting a pull request.

  

## License

  

This project is under the MIT License. See the `LICENSE` file for more details.
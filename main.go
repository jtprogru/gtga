package main

import (
	"flag"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	// Получение токена бота и ID канала из переменных окружения
	botToken := os.Getenv("GTGA_BOT_TOKEN")
	channelID := os.Getenv("GTGA_CHANNEL_ID")

	// Проверка наличия токена бота и ID канала
	if botToken == "" || channelID == "" {
		log.Fatal("Не заданы GTGA_BOT_TOKEN и/или GTGA_CHANNEL_ID в переменных окружения")
	}

	// Парсинг флага командной строки --msg для получения сообщения
	messageFlag := flag.String("msg", "", "Сообщение для отправки в канал")
	flag.Parse()

	// Проверка, был ли передан флаг --msg
	if *messageFlag == "" {
		flag.Usage()
		return
	}

	// Создание клиента Telegram бота
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatalf("Ошибка при создании бота: %s", err)
	}

	// Создание сообщения
	message := tgbotapi.NewMessageToChannel(channelID, *messageFlag)

	// Отправка сообщения
	_, err = bot.Send(message)
	if err != nil {
		log.Fatalf("Ошибка при отправке сообщения: %s", err)
	}

	log.Println("Сообщение успешно отправлено в канал")
}

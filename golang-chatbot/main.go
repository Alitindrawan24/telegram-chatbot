package main

import (
	"errors"
	"log"
	"os"

	telegrambot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	token := os.Getenv("TELEGRAM_CHATBOT_TOKEN")

	// Validate the token is not empty
	if token == "" {
		log.Panic(errors.New("telegram bot token is required"))
	}

	// Initialize the TelegramBot
	bot, err := telegrambot.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := telegrambot.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			msg := telegrambot.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}

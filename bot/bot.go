package main

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	token := os.Getenv("TELEGRAM_TOKEN")
	bot, _ := tgbotapi.NewBotAPI(token)

    // Создаем канал входящих апдейтов
    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60

    updates := bot.GetUpdatesChan(u)

    for update := range updates {
        if update.Message == nil { // ignore any non-Message updates
            continue
        }
		button1 := tgbotapi.NewInlineKeyboardButtonData("Кнопка 1", "DATA_1")
		keyboard := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(button1),)
        msg := tgbotapi.NewMessage(update.Message.Chat.ID, "выберите: ")
		msg.ReplyMarkup = keyboard
        bot.Send(msg)
    }
}
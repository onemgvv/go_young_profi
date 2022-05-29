package telegram

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	exampleError = errors.New("example custom error")
)

func (b *Bot) handleError(chatID int64, err error) {
	msg := tgbotapi.NewMessage(chatID, "Неопознанная ошибка")
	switch err {
	case exampleError:
		msg.Text = "Это тестовая кастомная ошибка!"
		b.bot.Send(msg)
	default:
		b.bot.Send(msg)
	}
}

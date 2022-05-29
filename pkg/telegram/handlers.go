package telegram

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const (
	commandStart = "start"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	default:
		return b.handleUnknownCommands(message)
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	b.bot.Send(msg)
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Инициализация окружения...\nЗапуск системы!")
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleUnknownCommands(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Неопознанная команда")
	_, err := b.bot.Send(msg)
	return err
}

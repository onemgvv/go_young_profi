package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	bot *tgbotapi.BotAPI
)

func telegramInit() {
	var err error

	bot, err = tgbotapi.NewBotAPI("")
	if err != nil {
		logrus.Panic(err)
	}

	bot.Debug = true
	logrus.Printf("Bot: %s authorized", bot.Self.UserName)
}

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("error config initialization: %s", err.Error())
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		logrus.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		_, err := bot.Send(msg)
		if err == nil {
			logrus.Error(err.Error())
		}
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/onemgvv/go_young_profi/pkg/telegram"
	"github.com/spf13/viper"
	"log"
)

var (
	bot *tgbotapi.BotAPI
)

func telegramInit(token string) {
	var err error

	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true
	log.Printf("Bot: %s authorized", bot.Self.UserName)

	tgBot := telegram.NewBot(bot)
	if err = tgBot.Start(); err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error configs initialization: %s", err.Error())
	}
	var token = viper.GetString("token")

	telegramInit(token)
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
)

var bot *tgbotapi.BotAPI

func telegramInit(token string) {
	var err error

	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		logrus.Panic(err)
	}

	bot.Debug = true
	logrus.Printf("Bot: %s authorized", bot.Self.UserName)

	url := os.Getenv("BASE_URL") + bot.Token
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(url))

	if err != nil {
		logrus.Panic(err)
	}
}

func webhookHandler(c *gin.Context) {
	defer c.Request.Body.Close()

	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logrus.Error(err.Error())
		return
	}

	var update tgbotapi.Update
	err = json.Unmarshal(bytes, &update)
	if err != nil {
		logrus.Error(err.Error())
		return
	}
}

func main() {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	var token = viper.GetString("token")

	if err := initConfig(); err != nil {
		logrus.Fatalf("error config initialization: %s", err.Error())
	}

	// Init telegram bot
	telegramInit(token)

	// Init gin router
	router := gin.New()
	// Use logger middleware
	router.Use(gin.Logger())
	// Set server endpoint
	router.POST("/"+bot.Token, webhookHandler)

	err := router.Run(":", port)
	if err != nil {
		logrus.Panic(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

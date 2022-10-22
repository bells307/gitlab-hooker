package main

import (
	"log"

	tm "github.com/and3rson/telemux/v2"
	http "github.com/bells307/gitlab-hooker/internal/controller/http"
	tg "github.com/bells307/gitlab-hooker/internal/controller/tg"
	"github.com/bells307/gitlab-hooker/internal/service"
	"github.com/bells307/gitlab-hooker/internal/utils"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
)

type Config struct {
	Listen                  string `mapstructure:"listen"`
	utils.TelegramBotConfig `mapstructure:"telegram"`
}

func loadConfig(path string) (config Config, err error) {
	viper.SetDefault("listen", "0.0.0.0:8888")

	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func main() {
	log.Printf("loading configuration ...")
	config, err := loadConfig(".")
	if err != nil {
		log.Fatal("can't load configuration:", err)
	}

	api, err := tgbotapi.NewBotAPI(config.TelegramBotConfig.Token)
	if err != nil {
		log.Fatal("can't create telegram bot api:", err)
	}

	if config.TelegramBotConfig.Debug {
		api.Debug = true
	}

	telegramService := service.NewTelegramService(api, config.TelegramBotConfig.Chats)
	updateHandler := tg.NewUpdateHandler(telegramService)
	mux := tm.NewMux()
	updateHandler.Register(api, mux)

	hookService := service.NewHookService(telegramService)

	router := gin.Default()
	hookHandler := http.NewHookHandler(hookService)
	hookHandler.Register(router)

	log.Println("starting telegram bot ...")
	go utils.RunTelegramBot(api, mux)
	log.Printf("start listening on %s ...", config.Listen)
	router.Run(config.Listen)
}

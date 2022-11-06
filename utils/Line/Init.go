package Line

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/spf13/viper"
	"time"
)

var bot *linebot.Client

func InitBot() *linebot.Client {
	bot, _ = linebot.New(viper.GetViper().GetString("channel_secret"), viper.GetViper().GetString("channel_access_token"))
	return bot
}

func GetBot() *linebot.Client {
	for {
		if bot != nil {
			return bot
		}
		if err := InitBot(); err != nil {
			time.Sleep(5 * time.Second)
		}
	}
}

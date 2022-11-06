package line

import (
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/spf13/viper"
	"net/http"
	"project/models"
	"project/utils/Line"
)

func PushMessage(c *gin.Context) {
	toUserID := viper.GetViper().GetString("userid")
	_, err := Line.GetBot().PushMessage(toUserID, linebot.NewTextMessage("hello")).Do()
	if err != nil {
		c.String(http.StatusInternalServerError, "Something went wrong...")
		return
	}
	c.String(http.StatusOK, "ok")
}
func StoreMessage(c *gin.Context) {
	events, err := Line.GetBot().ParseRequest(c.Request)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			UserID := event.Source.UserID
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				messageModel := &models.LineMessage{}
				err = messageModel.InsertMessage(models.LineMessage{
					UserID:  UserID,
					ID:      message.ID,
					Message: message.Text,
				})
				if err != nil {
					c.String(http.StatusInternalServerError, err.Error())
					return
				}
			}
		}
	}
	c.String(http.StatusOK, "ok")
}

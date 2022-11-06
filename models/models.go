package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"project/utils/Database"
)

type LineMessage struct {
	ID      string `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID  string `json:"userID,omitempty"`
	Message string `json:"message,omitempty"`
}

func (m LineMessage) InsertMessage(pack LineMessage) (err error) {
	_, err = Database.GetCollection().InsertOne(context.Background(), bson.D{{"userid", pack.UserID}, {"id", pack.ID},
		{"message",
			pack.Message}})
	if err != nil {
		return err
	}
	return err
}

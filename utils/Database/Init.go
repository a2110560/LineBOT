package Database

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var client *mongo.Client
var messages *mongo.Collection

func InitDatabase() error {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(viper.GetViper().GetString("mongo_DSN")))
	if err != nil {
		fmt.Println(err.Error())
	}
	messages = client.Database("line").Collection("messages")

	return err
}

func GetDB() *mongo.Client {
	for {
		if client != nil {
			return client
		}
		if err := InitDatabase(); err != nil {
			time.Sleep(5 * time.Second)
		}
	}
}
func GetCollection() *mongo.Collection {
	for {
		if messages != nil {
			return messages
		}
		if err := InitDatabase(); err != nil {
			time.Sleep(5 * time.Second)
		}
	}
}

package Mongo

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"project/models"
	"project/utils/Database"
	"strconv"
	"time"
)

func GetUserMessage(c *gin.Context) {
	var data []models.LineMessage
	userID := c.Param("userid")
	limitQuery := c.Query("limit")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	limit, _ := strconv.ParseInt(limitQuery, 10, 64)
	opts := options.Find().SetLimit(limit)
	filter := bson.D{{"userid", userID}}
	cur, err := Database.GetCollection().Find(ctx, filter, opts)

	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	if err = cur.All(ctx, &data); err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, data)
}

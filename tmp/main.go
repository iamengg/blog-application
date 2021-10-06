package main

import (
	"blog-application/global"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	global.DB.Collection("user").InsertOne(context.Background(), bson.M{"name": "test"})
}

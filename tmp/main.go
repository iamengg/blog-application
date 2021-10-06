package main

import (
	"context"
	"github.com/iamengg/blog-application/global"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	global.DB.Collection("user").InsertOne(context.Background(), bson.M{"name": "test"})
}

package utils;

import (
	"time"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
);

func GetCollection(Name string) (*mongo.Collection, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"));

	if (err != nil) {
		return nil, err;
	}
	
	database := client.Database("fundse");

	collection := database.Collection(Name);

	return collection, nil;
}

func GetContext(timeout time.Duration) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout);

	return ctx, cancel;
}
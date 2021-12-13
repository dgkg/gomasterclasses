package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	conn    *mongo.Client
	colUser *mongo.Collection
}

func New() *MongoDB {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	localCredential := options.Credential{
		//AuthMechanism: "MONGODB-AWS",
		Username: "root",
		Password: "example",
	}
	client, err := mongo.Connect(
		context.Background(),
		options.Client().SetAuth(localCredential),
		options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		panic(err)
	}
	log.Print("success connected to Mongodb")
	return &MongoDB{
		conn:    client,
		colUser: client.Database("main").Collection("users"),
	}
}

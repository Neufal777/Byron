package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB() (*mongo.Client, context.Context, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://byron:Black_nebula000@byron.dqvrs.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))

	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	return client, ctx, err

}

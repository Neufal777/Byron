package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB() (*mongo.Client, context.Context, error) {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://byron:Black_nebula000@byroncluster.dqvrs.mongodb.net/byroncluster?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
		log.Println("CONEXION ERROOOOR! 1")
	}

	return client, ctx, err
}

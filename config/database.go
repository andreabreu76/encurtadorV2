package config

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

func ConnectMongo() (*mongo.Client, error) {
	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%d", MongoUser, MongoPwd, MongoServer, MongoPort)
	ctx := context.TODO()
	session, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("Error connecting to mongo: ", err)
	}

	defer func(session *mongo.Client, ctx context.Context) {
		err := session.Disconnect(ctx)
		if err != nil {
			log.Fatal("Error disconnecting from mongo: ", err)
		}
	}(session, ctx)

	fmt.Printf("%T\n", session)
	if err := session.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("Error pinging mongo: ", err)
	}

	log.Printf("Connected to mongo at %s", mongoURI)
	return session, nil
}

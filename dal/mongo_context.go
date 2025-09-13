package dal

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoContext struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func NewMongoClient(ctx context.Context, uri, dbName string) (*MongoContext, error) {
	clientOpts := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	ctxPing, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := client.Ping(ctxPing, nil); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB:: %w", err)
	}

	log.Println("Connected to Mongo")

	return &MongoContext{
		Client:   client,
		Database: client.Database(dbName),
	}, nil
}

func (mc *MongoContext) Disconnect(ctx context.Context) error {
	if mc.Client != nil {
		return mc.Client.Disconnect(ctx)
	}

	return nil
}

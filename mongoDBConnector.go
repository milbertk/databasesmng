package databasesmng

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDBConnector struct {
	Client   *mongo.Client
	Database *mongo.Database
}

// NewMongoDBConnector connects to a MongoDB instance
func NewMongoDBConnector(uri string, dbname string) (*MongoDBConnector, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Ping the database to verify the connection
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	db := client.Database(dbname)

	return &MongoDBConnector{
		Client:   client,
		Database: db,
	}, nil
}

// Close disconnects from MongoDB
func (m *MongoDBConnector) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return m.Client.Disconnect(ctx)
}

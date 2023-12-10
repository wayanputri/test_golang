package db

import (
	"context"
	"log"
	"test_backend/app/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MgoConn ...
func MgoConn(cfg *config.AppConfig) *mongo.Client {

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(cfg.DBURL).SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

// MgoCollection call a collection with passing client value
func MgoCollection(cfg *config.AppConfig, client *mongo.Client) *mongo.Collection {
	return client.Database(cfg.DBName).Collection(cfg.DBCollection)
}

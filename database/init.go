package database

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	dbOnce   sync.Once
	dbClient *mongo.Client
)

// GetDBClient returns a singleton MongoDB client
func GetDBClient() *mongo.Client {
	dbOnce.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		mongoURI := os.Getenv("mongo_db_connection_string")
		if mongoURI == "" {
			log.Fatal("MONGODB_URI is not set in the environment")
		}

		clientOptions := options.Client().ApplyURI(mongoURI)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatal("Error connecting to MongoDB: ", err)
		}

		err = client.Ping(ctx, nil)
		if err != nil {
			log.Fatal("Couldn't ping MongoDB: ", err)
		}

		dbClient = client
	})

	return dbClient
}

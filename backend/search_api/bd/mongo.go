package bd

import (
	"context"
	"os"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoClient *mongo.Client
	Database    *mongo.Database
)

func InitMongoDB() error {
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017/"
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Errorf("Error connecting to MongoDB: %v", err)
		return err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Errorf("Error pinging MongoDB: %v", err)
		return err
	}

	dbName := os.Getenv("DATABASE_NAME")
	if dbName == "" {
		dbName = "search_db"
	}

	MongoClient = client
	Database = client.Database(dbName)

	log.Info("MongoDB connection initialized successfully")
	return nil
}

func CloseConnection() {
	if MongoClient != nil {
		MongoClient.Disconnect(context.Background())
		log.Info("MongoDB connection closed")
	}
}

func GetDatabase() *mongo.Database {
	return Database
}

package bd

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectMongo() error {
	// Nombre del contenedor de Mongo como host en Docker Compose
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		uri = "mongodb://mongo_hoteles:27017" // <-- aquí cambiamos localhost por mongo_hoteles
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	// Ping para asegurarnos de que la conexión funciona
	if err := client.Ping(ctx, nil); err != nil {
		return err
	}

	MongoClient = client
	log.Println("MongoDB conectado correctamente")
	return nil
}

func GetCollection(dbName, colName string) *mongo.Collection {
	return MongoClient.Database(dbName).Collection(colName)
}

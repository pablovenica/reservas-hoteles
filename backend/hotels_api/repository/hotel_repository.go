package repository

import (
	"context"
	"time"

	"hotels_api/bd"
	"hotels_api/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	log "github.com/sirupsen/logrus"
)

var (
	collectionName = "hoteles"
	dbName         = "hoteles_db"
)

func GetHotels() []domain.Hotel {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := bd.GetCollection(dbName, collectionName)

	filter := bson.M{}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		log.Error("Error al listar hoteles:", err)
		return nil
	}
	defer cur.Close(ctx)

	var hotels []domain.Hotel
	if err := cur.All(ctx, &hotels); err != nil {
		log.Error("Error decoding hoteles:", err)
		return nil
	}

	return hotels
}

func GetHotelByIdHex(idHex string) (domain.Hotel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := bd.GetCollection(dbName, collectionName)

	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		return domain.Hotel{}, err
	}

	var hotel domain.Hotel
	if err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&hotel); err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.Hotel{}, nil
		}
		return domain.Hotel{}, err
	}

	return hotel, nil
}

func InsertHotel(h domain.Hotel) (domain.Hotel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := bd.GetCollection(dbName, collectionName)

	res, err := collection.InsertOne(ctx, h)
	if err != nil {
		return domain.Hotel{}, err
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		h.ID = oid
	}
	return h, nil
}

func UpdateHotelByHex(idHex string, updated domain.Hotel) (domain.Hotel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := bd.GetCollection(dbName, collectionName)

	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		return domain.Hotel{}, err
	}

	update := bson.M{
		"$set": bson.M{
			"nombre":      updated.Nombre,
			"imagen":      updated.Imagen,
			"descripcion": updated.Descripcion,
			"provincia":   updated.Provincia,
			"direccion":   updated.Direccion,
			"precio":      updated.Precio,
		},
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var result domain.Hotel

	if err := collection.FindOneAndUpdate(ctx, bson.M{"_id": id}, update, opts).Decode(&result); err != nil {
		return domain.Hotel{}, err
	}

	return result, nil
}

func DeleteHotelByHex(idHex string) (domain.Hotel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := bd.GetCollection(dbName, collectionName)

	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		return domain.Hotel{}, err
	}

	opts := options.FindOneAndDelete()
	var result domain.Hotel

	if err := collection.FindOneAndDelete(ctx, bson.M{"_id": id}, opts).Decode(&result); err != nil {
		return domain.Hotel{}, err
	}

	return result, nil
}

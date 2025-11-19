package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Search struct {
	ID        string             `bson:"_id"`
	UserID    string             `bson:"user_id"`
	HotelName string             `bson:"hotel_name"`
	City      string             `bson:"city"`
	CheckIn   string             `bson:"check_in"`
	CheckOut  string             `bson:"check_out"`
	Guests    int                `bson:"guests"`
	Timestamp time.Time          `bson:"timestamp"`
	Status    string             `bson:"status"`
	MongoID   primitive.ObjectID `bson:"_id,omitempty"`
}

type SearchResult struct {
	HotelID   string  `json:"hotel_id"`
	HotelName string  `json:"hotel_name"`
	City      string  `json:"city"`
	Price     float64 `json:"price"`
	Rating    float64 `json:"rating"`
}

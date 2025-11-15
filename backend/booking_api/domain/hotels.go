package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Hotel struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Titulo   string             `bson:"titulo" json:"titulo"`
	Nivel    int                `bson:"nivel" json:"nivel"`
	Estado   bool               `bson:"estado" json:"estado"`
	Duracion string             `bson:"duracion" json:"duracion"`
	Precio   string             `bson:"precio" json:"precio"`
}

type Hotels []Hotel

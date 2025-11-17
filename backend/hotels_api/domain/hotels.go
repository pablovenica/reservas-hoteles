package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Hotel struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre      string             `bson:"nombre"`
	Imagen		string 			   `bson:"imagen"`
	Descripcion string 		 	   `bson:"descripcion"`
	Provincia   string			   `bson:provincia`
	Direccion   string             `bson:"duracion"`
	Precio      float64            `bson:"precio"`
}

type Hotels []Hotel

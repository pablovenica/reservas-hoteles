package repository

import (
	"context"
	"search_api/bd"
	"search_api/domain"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const CollectionName = "searches"

// SaveSearch guarda una búsqueda en MongoDB
func SaveSearch(ctx context.Context, search *domain.Search) error {
	collection := bd.GetDatabase().Collection(CollectionName)

	result, err := collection.InsertOne(ctx, search)
	if err != nil {
		log.Errorf("Error saving search: %v", err)
		return err
	}

	log.Infof("Search saved with ID: %v", result.InsertedID)
	return nil
}

// GetSearchByID obtiene una búsqueda por ID
func GetSearchByID(ctx context.Context, searchID string) (*domain.Search, error) {
	collection := bd.GetDatabase().Collection(CollectionName)

	var search domain.Search
	err := collection.FindOne(ctx, bson.M{"_id": searchID}).Decode(&search)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Warnf("Search not found: %s", searchID)
			return nil, nil
		}
		log.Errorf("Error getting search: %v", err)
		return nil, err
	}

	return &search, nil
}

// GetSearchesByUserID obtiene todas las búsquedas de un usuario
func GetSearchesByUserID(ctx context.Context, userID string) ([]domain.Search, error) {
	collection := bd.GetDatabase().Collection(CollectionName)

	cursor, err := collection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		log.Errorf("Error getting searches: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var searches []domain.Search
	err = cursor.All(ctx, &searches)
	if err != nil {
		log.Errorf("Error decoding searches: %v", err)
		return nil, err
	}

	return searches, nil
}

// DeleteSearch elimina una búsqueda
func DeleteSearch(ctx context.Context, searchID string) error {
	collection := bd.GetDatabase().Collection(CollectionName)

	result, err := collection.DeleteOne(ctx, bson.M{"_id": searchID})
	if err != nil {
		log.Errorf("Error deleting search: %v", err)
		return err
	}

	if result.DeletedCount == 0 {
		log.Warnf("Search not found for deletion: %s", searchID)
		return nil
	}

	log.Infof("Search deleted: %s", searchID)
	return nil
}

// UpdateSearchStatus actualiza el estado de una búsqueda
func UpdateSearchStatus(ctx context.Context, searchID, status string) error {
	collection := bd.GetDatabase().Collection(CollectionName)

	_, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": searchID},
		bson.M{"$set": bson.M{"status": status}},
	)

	if err != nil {
		log.Errorf("Error updating search status: %v", err)
		return err
	}

	return nil
}

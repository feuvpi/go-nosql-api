package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/your-username/mongodb-api/internal/models"
)

type ListingsDatabase struct {
	collection *mongo.Collection
}

func NewListingsDatabase(db *mongo.Database) *ListingsDatabase {
	return &ListingsDatabase{
		collection: db.Collection("listings"),
	}
}

func (db *ListingsDatabase) GetAll() ([]models.Listing, error) {
	cursor, err := db.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var listings []models.Listing
	for cursor.Next(context.TODO()) {
		var listing models.Listing
		if err := cursor.Decode(&listing); err != nil {
			return nil, err
		}
		listings = append(listings, listing)
	}
	return listings, nil
}

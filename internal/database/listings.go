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

func (db *ListingsDatabase) GetByID(id primitive.ObjectID) (models.Listing, error) {
	filter := bson.M{"_id": id}
	var listing models.Listing
	err := db.collection.FindOne(context.TODO(), filter).Decode(&listing)
	if err != nil {
		return models.Listing{}, err
	}
	return listing, nil
}

func (db *ListingsDatabase) Insert(listing models.Listing) (primitive.ObjectID, error) {
	result, err := db.collection.InsertOne(context.TODO(), listing)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func (db *ListingsDatabase) Update(id primitive.ObjectID, listing models.Listing) (int64, error) {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": listing}
	result, err := db.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return 0, err
	}
	return result.ModifiedCount, nil
}

func (db *ListingsDatabase) Delete(id primitive.ObjectID) (int64, error) {
	filter := bson.M{"_id": id}
	result, err := db.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}

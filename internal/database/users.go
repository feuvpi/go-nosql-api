package database 

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/your-username/mongodb-api/models"
)

type UserDatabase struct { 
	collection *mongo.Collection
}

func NewUsersDatabase(db *mongo.Database) *UsersDatabase {
	return &UsersDatabase{
		collection: db.Collection("users"),
	}
}

func (db *UsersDatabase) GetAll() ([]models.User, error) {
	cursor, err := db.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var users []models.User
	for cursor.Next(context.TODO()) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (db *UsersDatabase) Insert(user models.User) (primitive.ObjectID, error) {
	result, err := db.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func (db *UsersDatabase) Update(id primitive.ObjectID, user models.User) (int64, error) {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": user}
	result, err := db.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return 0, err
	}
	return result.ModifiedCount, nil
}

func (db *UsersDatabase) Delete(id primitive.ObjectID) (int64, error) {
	filter := bson.M{"_id": id}
	result, err := db.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}

func (db *UsersDatabase) GetByID(id primitive.ObjectID) (models.User, error) {
	filter := bson.M{"_id": id}
	var user models.User
	err := db.collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
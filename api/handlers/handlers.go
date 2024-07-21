package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/your-username/mongodb-api/models"
)

var db *mongo.Database

func getListings(w http.ResponseWriter, r *http.Request) {
	collection := db.Collection("listings")

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	var listings []models.Listing
	for cursor.Next(context.TODO()) {
		var listing models.Listing
		err := cursor.Decode(&listing)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		listings = append(listings, listing)
	}

	json.NewEncoder(w).Encode(listings)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	collection := db.Collection("users")

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	var users []models.User
	for cursor.Next(context.TODO()) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)
}

func insertListing(w http.ResponseWriter, r *http.Request) {
	collection := db.Collection("listings")

	var listing models.Listing
	err := json.NewDecoder(r.Body).Decode(&listing)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := collection.InsertOne(context.TODO(), listing)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result.InsertedID)
}

func updateListing(w http.ResponseWriter, r *http.Request) {
	collection := db.Collection("listings")

	idParam := mux.Vars(r)["id"]
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var listing models.Listing
	err = json.NewDecoder(r.Body).Decode(&listing)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": listing}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result.ModifiedCount)
}

func deleteListing(w http.ResponseWriter, r *http.Request) {
	collection := db.Collection("listings")

	idParam := mux.Vars(r)["id"]
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result.DeletedCount)
}

func getListingByID(w http.ResponseWriter, r *http.Request) {
	collection := db.Collection("listings")

	idParam := mux.Vars(r)["id"]
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	filter := bson.M{"_id": id}
	var listing models.Listing
	err = collection.FindOne(context.TODO(), filter).Decode(&listing)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(listing)
}

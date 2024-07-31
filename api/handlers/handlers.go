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
	"github.com/go-playground/validator/v10"
)

var listingsService *services.ListingsService
var usersService *services.UsersService

func getListings(w http.ResponseWriter, r *http.Request) {
	listings, err := listingsService.GetAllListings()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "No listings found", http.StatusNotFound)
		} else {
			http.Error(w, "Error fetching listings", http.StatusInternalServerError)
		}
		return
	}
	json.NewEncoder(w).Encode(listings)
}
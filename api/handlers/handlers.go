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

var listingsService *services.ListingsService
var usersService *services.UsersService

func getListings(w http.ResponseWriter, r *http.Request) {
	listings, err := listingsService.GetAllListings()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(listings)
}
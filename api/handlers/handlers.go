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

type QueryParams struct {
    Page     int `validate:"min=1"`
    PageSize int `validate:"min=1,max=100"`
}

func getPaginatedListings(w http.ResponseWriter, r *http.Request) {
	params := QueryParams{
        Page:     r.URL.Query().Get("page"),
        PageSize: r.URL.Query().Get("pageSize"),
    }
	
	if err := validator.New().Struct(params); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	page, _ := strconv.Atoi(params.Page)
    pageSize, _ := strconv.Atoi(params.PageSize)

	listings, err := listingsService.GetPaginatedListings(page, pageSize)
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

func getListings(w http.ResponseWriter, r *http.Request) {
	params := QueryParams{
        Page:     r.URL.Query().Get("page"),
        PageSize: r.URL.Query().Get("pageSize"),
    }

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
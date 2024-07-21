package services

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/your-username/mongodb-api/internal/database"
	"github.com/your-username/mongodb-api/internal/models"
)

type ListingsService struct {
	db *database.ListingsDatabase
}

func NewListingsService(db *database.ListingsDatabase) *ListingsService {
	return &ListingsService{db: db}
}



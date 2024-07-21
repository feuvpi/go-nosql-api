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

func (s *ListingsService) GetAllListings() ([]models.Listing, error) {
	return s.db.GetAll()
}

func (s *ListingsService) GetListingByID(id string) (*models.Listing, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return s.db.GetByID(objectID)
}

func (s *ListingsService) CreateListing(listing models.Listing) (interface{}, error) {
	return s.db.Insert(listing)
}

func (s *ListingsService) UpdateListing(id string, listing models.Listing) (int64, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}
	return s.db.Update(objectID, listing)
}

func (s *ListingsService) DeleteListing(id string) (int64, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}
	return s.db.Delete(objectID)
}
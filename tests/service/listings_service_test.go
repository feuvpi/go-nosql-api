package services_test

import (
    "context"
    "testing"
    "time"

    "github.com/stretchr/testify/assert"
    "github.com/your-username/mongodb-api/database"
    "github.com/your-username/mongodb-api/models"
    "github.com/your-username/mongodb-api/services"
    "github.com/your-username/mongodb-api/tests"
)

// -- createTestListing creates a test listing for use in tests.
func createTestListing(db *mongo.Database) (*mongo.InsertOneResult, error) {
    listingsDB := database.NewListingsDatabase(db)

    listing := models.Listing{
        Title:       "Test Listing",
        Description: "Description",
        Category:    "Category",
        State:       "State",
        City:        "City",
        Zipcode:     "Zipcode",
        Coordinates: "Coordinates",
        Date:        time.Now(),
        ImagesUrls:  []string{"http://example.com/image.jpg"},
    }

    return listingsDB.collection.InsertOne(context.Background(), listing)
}

func TestListingsService_GetAll(t *testing.T) {
    db, teardown := tests.SetupTestDB(t)
    defer teardown()

    listingsDB := database.NewListingsDatabase(db)
    listingsService := services.NewListingsService(listingsDB)

    _, err := createTestListing(db)
    if err != nil {
        t.Fatal(err)
    }

    listings, err := listingsService.GetAll()
    assert.NoError(t, err)
    assert.NotEmpty(t, listings)
}

func TestListingsService_GetByID(t *testing.T) {
    db, teardown := tests.SetupTestDB(t)
    defer teardown()

    listingsDB := database.NewListingsDatabase(db)
    listingsService := services.NewListingsService(listingsDB)

    result, err := createTestListing(db)
    if err != nil {
        t.Fatal(err)
    }

    id := result.InsertedID.(primitive.ObjectID)
    listing, err := listingsService.GetByID(id)
    assert.NoError(t, err)
    assert.Equal(t, "Test Listing", listing.Title)
}

func TestListingsService_Insert(t *testing.T) {
    db, teardown := tests.SetupTestDB(t)
    defer teardown()

    listingsDB := database.NewListingsDatabase(db)
    listingsService := services.NewListingsService(listingsDB)

    listing := models.Listing{
        Title:       "New Listing",
        Description: "New Description",
        Category:    "New Category",
        State:       "New State",
        City:        "New City",
        Zipcode:     "New Zipcode",
        Coordinates: "New Coordinates",
        Date:        time.Now(),
        ImagesUrls:  []string{"http://example.com/new-image.jpg"},
    }

    id, err := listingsService.Insert(listing)
    assert.NoError(t, err)

    retrieved, err := listingsDB.collection.FindOne(context.Background(), bson.M{"_id": id}).DecodeBytes()
    assert.NoError(t, err)
    assert.Equal(t, listing.Title, retrieved.Lookup("title").StringValue())
}
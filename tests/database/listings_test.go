package database

import (
    "context"
    "testing"
    "time"

    "github.com/stretchr/testify/assert"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"

    "github.com/your-username/mongodb-api/models"
)

// createTestListing creates a test listing for use in tests.
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

func TestInsertListing(t *testing.T) {
    db, teardown := SetupTestDB(t)
    defer teardown()

    listingsDB := database.NewListingsDatabase(db)

    result, err := createTestListing(db)
    assert.NoError(t, err)
    assert.NotNil(t, result.InsertedID)

    filter := bson.M{"_id": result.InsertedID}
    var retrieved models.Listing
    err = listingsDB.collection.FindOne(context.Background(), filter).Decode(&retrieved)
    assert.NoError(t, err)
    assert.Equal(t, "Test Listing", retrieved.Title)
}

func TestGetAllListings(t *testing.T) {
    db, teardown := SetupTestDB(t)
    defer teardown()

    listingsDB := database.NewListingsDatabase(db)

    _, err := createTestListing(db)
    if err != nil {
        t.Fatal(err)
    }

    listings, err := listingsDB.GetAll()
    assert.NoError(t, err)
    assert.NotEmpty(t, listings)
}
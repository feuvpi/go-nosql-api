package tests

import (
    "context"
    "testing"
    "time"

    "github.com/stretchr/testify/assert"
    "go.mongodb.org/mongo-driver/bson"
    "github.com/your-username/mongodb-api/database"
    "github.com/your-username/mongodb-api/models"
)

// -- createTestUser creates a test user for use in tests.
func createTestUser(db *mongo.Database) (*mongo.InsertOneResult, error) {
    usersDB := database.NewUsersDatabase(db)

    user := models.User{
        Name:         "Test User",
        Email:        "test@example.com",
        MobileNumber: "1234567890",
        Password:     "hashedpassword",
        PasswordSalt: "somesalt",
    }

    return usersDB.collection.InsertOne(context.Background(), user)
}

func TestInsertUser(t *testing.T) {
    db, teardown := SetupTestDB(t)
    defer teardown()

    usersDB := database.NewUsersDatabase(db)

    result, err := createTestUser(db)
    assert.NoError(t, err)
    assert.NotNil(t, result.InsertedID)

    filter := bson.M{"_id": result.InsertedID}
    var retrieved models.User
    err = usersDB.collection.FindOne(context.Background(), filter).Decode(&retrieved)
    assert.NoError(t, err)
    assert.Equal(t, "Test User", retrieved.Name)
}

func TestGetAllUsers(t *testing.T) {
    db, teardown := SetupTestDB(t)
    defer teardown()

    usersDB := database.NewUsersDatabase(db)

    _, err := createTestUser(db)
    if err != nil {
        t.Fatal(err)
    }

    users, err := usersDB.GetAll()
    assert.NoError(t, err)
    assert.NotEmpty(t, users)
}
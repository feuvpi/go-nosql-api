package services_test

import (
    "context"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/your-username/mongodb-api/database"
    "github.com/your-username/mongodb-api/models"
    "github.com/your-username/mongodb-api/services"
    "github.com/your-username/mongodb-api/tests"
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

func TestUsersService_GetAll(t *testing.T) {
    db, teardown := tests.SetupTestDB(t)
    defer teardown()

    usersDB := database.NewUsersDatabase(db)
    usersService := services.NewUsersService(usersDB)

    _, err := createTestUser(db)
    if err != nil {
        t.Fatal(err)
    }

    users, err := usersService.GetAll()
    assert.NoError(t, err)
    assert.NotEmpty(t, users)
}

func TestUsersService_GetByID(t *testing.T) {
    db, teardown := tests.SetupTestDB(t)
    defer teardown()

    usersDB := database.NewUsersDatabase(db)
    usersService := services.NewUsersService(usersDB)

    result, err := createTestUser(db)
    if err != nil {
        t.Fatal(err)
    }

    id := result.InsertedID.(primitive.ObjectID)
    user, err := usersService.GetByID(id)
    assert.NoError(t, err)
    assert.Equal(t, "Test User", user.Name)
}

func TestUsersService_Insert(t *testing.T) {
    db, teardown := tests.SetupTestDB(t)
    defer teardown()

    usersDB := database.NewUsersDatabase(db)
    usersService := services.NewUsersService(usersDB)

    user := models.User{
        Name:         "New User",
        Email:        "new@example.com",
        MobileNumber: "0987654321",
        Password:     "newhashedpassword",
        PasswordSalt: "newsalt",
    }

    id, err := usersService.Insert(user)
    assert.NoError(t, err)

    retrieved, err := usersDB.collection.FindOne(context.Background(), bson.M{"_id": id}).DecodeBytes()
    assert.NoError(t, err)
    assert.Equal(t, user.Name, retrieved.Lookup("name").StringValue())
}

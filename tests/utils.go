package tests

import (
    "context"
    "testing"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// -- SetupTestDB sets up a test database and returns the database instance along with a teardown function.
func SetupTestDB(t *testing.T) (*mongo.Database, func()) {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        t.Fatal(err)
    }

    db := client.Database("testdb")
    return db, func() {
        client.Disconnect(context.Background())
    }
}
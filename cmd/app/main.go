import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/feuvpi/go-nosql-api/handlers"
	"github.com/feuvpi/go-nosql-api/database"
)

// var db *mongo.Database

func main() {
	client, db, err := database.Connect("mongodb://localhost:27017", "mydb")

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())

	// -- setup routes
	router := mux.NewRouter()
	router.HandleFunc("/listings", getListings).Methods("GET")
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/listings", insertListing).Methods("POST")
	router.HandleFunc("/listings/{id}", updateListing).Methods("PUT")
	router.HandleFunc("/listings/{id}", deleteListing).Methods("DELETE")
	router.HandleFunc("/listings/{id}", getListingByID).Methods("GET")

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}
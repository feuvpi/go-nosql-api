import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	db = client.Database("mydb")

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
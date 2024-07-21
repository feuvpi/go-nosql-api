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
	"github.com/feuvpi/go-nosql-api/internal"
)

var listingsService *services.ListingsService
var usersService *services.UsersService

func main() {
	// -- Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
	 -- Load configuration from environment variables
	cfg := config.LoadFromEnv()

	// -- connect to database
	client, db, err := database.Connect(cfg.MongoDBURI, cfg.MongoDBName)

	if err != nil {
		log.Fatal(err)
	}

	// defer client.Disconnect(context.Background())

	listingsDB := database.NewListingsDatabase(db)
	usersDB := database.NewUsersDatabase(db)

	listingsService = services.NewListingsService(listingsDB)
	usersService = services.NewUsersService(usersDB)

	handlers.SetServices(listingsService, usersService)

	// Setup routes
    router := mux.NewRouter()
    router.HandleFunc("/listings", handlers.GetListings).Methods("GET")
    router.HandleFunc("/users", handlers.GetUsers).Methods("GET")
    router.HandleFunc("/listings", handlers.InsertListing).Methods("POST")
    router.HandleFunc("/listings/{id}", handlers.UpdateListing).Methods("PUT")
    router.HandleFunc("/listings/{id}", handlers.DeleteListing).Methods("DELETE")
    router.HandleFunc("/listings/{id}", handlers.GetListingByID).Methods("GET")
    router.HandleFunc("/users", handlers.InsertUser).Methods("POST")
    router.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
    router.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")
    router.HandleFunc("/users/{id}", handlers.GetUserByID).Methods("GET")

	log.Println("Starting server at port", cfg.Port)
    if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
        log.Fatal(err)
    }

}
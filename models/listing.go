package models

import (
	"time"
)

type Listing struct {
	ID          string    `bson:"_id,omitempty"`
	Title       string    `bson:"title"`
	Description string    `bson:"description"`
	Category    string    `bson:"category"`
	State       string    `bson:"state"`
	City        string    `bson:"city"`
	Zipcode     string    `bson:"zipcode"`
	Coordinates string    `bson:"coordinates"`
	Date        time.Time `bson:"date"`
	ImagesUrls  []string  `bson:"imagesUrls"`
}

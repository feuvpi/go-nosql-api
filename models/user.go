package models

type User struct {
	ID           string `bson:"_id,omitempty"`
	Name         string `bson:"name"`
	Email        string `bson:"email"`
	MobileNumber string `bson:"mobileNumber"`
	Password     string `bson:"password"`
	PasswordSalt string `bson:"passwordSalt"`
}

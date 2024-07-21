package services

import (
	"github.com/your-username/mongodb-api/database"
	"github.com/your-username/mongodb-api/models"
)

type UsersService struct {
	db *database.UsersDatabase
}

func NewUsersService(db *database.UsersDatabase) *UsersService {
	return &UsersService{db: db}
}

func (s *UsersService) GetAllUsers() ([]models.User, error) {
	return s.db.GetAll()
}

func (s *UsersService) GetUserByID(id string) (models.User, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return s.db.GetByID(oid)

func (s *UsersService) CreateUser(user models.User) (primitive.ObjectID, error) {
	return s.db.Insert(user)
}

func (s *UsersService) UpdateUser(id string, user models.User) (int64, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}
	return s.db.Update(oid, user)
}

func (s *UsersService) DeleteUser(id string) (int64, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}
	return s.db.Delete(oid)
}
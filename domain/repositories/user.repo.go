package repositories

import (
	"errors"
	"hulk/go-webservice/domain/entities"
	"hulk/go-webservice/infrastructure/persist"
)

type UserRepository struct{}

func (m *UserRepository) GetUsers() []entities.User {
	var users []entities.User
	persist.DB.Find(&users)
	return users
}

func (m *UserRepository) FindUserById(userId string) (user entities.User, err error) {
	r := persist.DB.First(&user, userId)
	if r.RowsAffected > 0 {
		err = errors.New("Not found")
		return
	}
	return
}

func (m *UserRepository) FindUserByEmail(email string) (user entities.User, err error) {
	r := persist.DB.Where("email = ?", email).First(&user)
	if r.RowsAffected > 0 {
		err = errors.New("Not found")
		return
	}
	return
}

func (m *UserRepository) IsEmailExisted(email string) bool {
	var duplicated entities.User
	r := persist.DB.Where("email = ?", email).First(&duplicated)
	return r.RowsAffected > 0
}

func (m *UserRepository) IsUserExist(userID string) bool {
	var user entities.User
	r := persist.DB.First(&user, userID)
	return r.RowsAffected > 0
}

func (m *UserRepository) CreateUser(user entities.User) (entities.User, error) {
	persist.DB.Create(&user)
	return user, nil
}

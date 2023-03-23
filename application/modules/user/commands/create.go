package commands

import (
	"errors"
	"hulk/go-webservice/common"
	"hulk/go-webservice/domain/entities"
	"hulk/go-webservice/domain/repositories"
	"time"
)

type CreateUserDto struct {
	FirstName   string    `form:"first_name" json:"first_name" xml:"first_name"  binding:"required"`
	LastName    string    `form:"last_name" json:"last_name" xml:"last_name"  binding:"required"`
	Email       string    `form:"email" json:"email" xml:"email"  binding:"required"`
	Password    string    `form:"password" json:"password" xml:"password" binding:"required"`
	Hobby       string    `form:"hobby" json:"hobby" xml:"hobby"`
	PhoneNumber string    `form:"phone_number" json:"phone_number" xml:"phone_number" binding:"required"`
	Address     string    `form:"address" json:"address" xml:"address"`
	Dob         time.Time `form:"dob" json:"dob" xml:"dob"`
}

func CreateUserCommand(request CreateUserDto) (entities.User, error) {
	var user entities.User

	userRepo := repositories.UserRepository{}

	if userRepo.IsEmailExisted(request.Email) {
		return user, errors.New("Email already has been taken")
	}

	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.Email = request.Email
	user.Address = request.Address
	user.Hobby = request.Hobby
	user.PhoneNumber = request.PhoneNumber
	user.Dob = request.Dob
	passwordHashed, err := common.HashPassword(request.Password)
	if err != nil {
		return user, err
	}
	user.Password = passwordHashed

	userCreated, err := userRepo.CreateUser(user)
	if err != nil {
		return user, err
	}
	return userCreated, nil
}

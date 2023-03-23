package commands

import (
	"errors"
	"hulk/go-webservice/common"
	"hulk/go-webservice/domain/entities"
	"hulk/go-webservice/domain/repositories"
)

type LoginRequestDto struct {
	Email    string `form:"email" json:"email" xml:"email"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password"  binding:"required"`
}

func LoginCommand(request LoginRequestDto) (string, error) {
	var user entities.User

	userRepo := repositories.UserRepository{}

	user, err := userRepo.FindUserByEmail(request.Email)
	if err != nil {
		return "", err
	}

	if err := common.CheckPassword(request.Password, user.Password); err != nil {
		return "", errors.New("Password does not match")
	}

	token, err := common.GenerateJWT(common.UserClaim{Username: user.FirstName + "" + user.LastName, Id: user.ID})

	if err != nil {
		return "", errors.New("Unable generate token")
	}

	return token, nil
}

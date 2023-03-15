package queries

import (
	"hulk/go-webservice/domain/entities"
	"hulk/go-webservice/domain/repositories"
)

func GetAllUserQuery() []entities.User {
	userRepo := repositories.UserRepository{}
	return userRepo.GetUsers()
}

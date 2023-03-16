package commands

import (
	"errors"
	"hulk/go-webservice/application/services"
	"hulk/go-webservice/common"
	"hulk/go-webservice/domain/entities"
	"hulk/go-webservice/domain/repositories"
	"hulk/go-webservice/infrastructure/persist"
	"mime/multipart"
)

func UpdateUserAvatarCommand(userID string, file *multipart.FileHeader) (entities.User, error) {
	var user entities.User
	userRepo := repositories.UserRepository{}
	if userRepo.IsUserExist(userID) {
		return user, errors.New("User not found")
	}

	fileUploadedPath, err := services.UploadAzBlob(file)
	if err != nil {
		return user, err
	}
	common.RemoveFile(file.Filename)
	user.Avatar = fileUploadedPath
	persist.DB.Save(&user)
	return user, nil
}

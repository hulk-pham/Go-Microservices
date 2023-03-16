package handler

import (
	"fmt"
	"hulk/go-webservice/domain/entities"
	"hulk/go-webservice/presentation/rpc/pb"
)

var TIME_FORMAT = "2006-01-02 15:04"

func ParseUser(user entities.User) *pb.User {
	return &pb.User{
		Id:          fmt.Sprint(user.ID),
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		Password:    user.Password,
		Hobby:       user.Hobby,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		Dob:         user.Dob.Format(TIME_FORMAT),
		Avatar:      user.Avatar,
		CreatedAt:   user.CreatedAt.Format(TIME_FORMAT),
		UpdatedAt:   user.UpdatedAt.Format(TIME_FORMAT),
		DeletedAt:   user.DeletedAt.Time.Format(TIME_FORMAT),
	}
}

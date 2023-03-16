package handler

import (
	"context"
	"hulk/go-webservice/application/modules/user/commands"
	"hulk/go-webservice/application/modules/user/queries"
	"hulk/go-webservice/presentation/rpc/pb"
	"time"
)

func (s *Server) CreateUser(_ context.Context, input *pb.CreateUserRequest) (*pb.UserResponse, error) {

	var request commands.CreateUserDto

	dob, err := time.Parse(TIME_FORMAT, input.Dob)

	request = commands.CreateUserDto{
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Email:       input.Email,
		Password:    input.Password,
		Hobby:       input.Hobby,
		PhoneNumber: input.PhoneNumber,
		Address:     input.Address,
		Dob:         dob,
	}

	user, err := commands.CreateUserCommand(request)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{User: ParseUser(user)}, nil
}

func (s *Server) GetUsers(req *pb.GetUsersRequest, stream pb.UserService_GetUsersServer) error {
	users := queries.GetAllUserQuery()
	for _, user := range users {
		stream.Send(ParseUser(user))
	}
	return nil
}

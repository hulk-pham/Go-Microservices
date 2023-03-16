package handler

import (
	"hulk/go-webservice/presentation/rpc/pb"
)

type Server struct {
	pb.UnimplementedGreeterServer
	pb.UnimplementedUserServiceServer
}

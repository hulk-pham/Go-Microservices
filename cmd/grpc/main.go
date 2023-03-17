package main

import (
	"fmt"
	"log"
	"net"

	"hulk/go-webservice/domain/entities"
	"hulk/go-webservice/infrastructure/config"
	"hulk/go-webservice/infrastructure/persist"
	"hulk/go-webservice/presentation/rpc/handler"
	"hulk/go-webservice/presentation/rpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config := config.AppConfig()

	lis, err := net.Listen("tcp", ":"+config.RpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	persist.InitDB()
	s := grpc.NewServer()

	persist.DB.AutoMigrate(&entities.User{})
	pb.RegisterGreeterServer(s, &handler.Server{})
	pb.RegisterUserServiceServer(s, &handler.Server{})

	fmt.Println("Run RPC server at :" + config.RpcPort)

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

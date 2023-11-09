package main

import (
	"fmt"
	"log"
	"net"
	"os"

	pbUser "github.com/0x726f6f6b6965/mono-service/protos/user/v1"
	"github.com/0x726f6f6b6965/mono-service/services/user-service/internal/dao/psql"
	"github.com/0x726f6f6b6965/mono-service/services/user-service/internal/service"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	godotenv.Load()
	psql.InitUserDB()
	server := service.NewUserService()
	grpcServer := grpc.NewServer()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("GRPC_USER_PORT")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	pbUser.RegisterUserServiceServer(grpcServer, server)
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

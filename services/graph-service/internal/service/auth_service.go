package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	pbCommon "github.com/0x726f6f6b6965/mono-service/protos/common/v1"
	pbUser "github.com/0x726f6f6b6965/mono-service/protos/user/v1"
	"github.com/0x726f6f6b6965/mono-service/services/graph-service/internal/model"
	"github.com/0x726f6f6b6965/mono-service/services/graph-service/internal/utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// create a new service
type UserService struct{}

// Register returns JWT token for authentication
func (u *UserService) Register(ctx context.Context, input model.NewUser) (string, error) {
	// create a password with bcrypt encryption
	bs, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	// create a variable to store the encrypted password
	var password string = string(bs)

	// create a new user
	newUser := pbUser.SetUserRequest{
		Id:       uuid.New().String(),
		Username: input.Username,
		Password: password,
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", os.Getenv("GRPC_USER"),
		os.Getenv("GRPC_USER_PORT")), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("%s, %s, %v", os.Getenv("GRPC_USER"),
			os.Getenv("GRPC_USER_PORT"), err)
		return "", err
	}
	defer conn.Close()
	userClient := pbUser.NewUserServiceClient(conn)

	resp, err := userClient.SetUser(ctx, &newUser)
	if err != nil {
		return "", err
	}

	if !resp.Ok {
		return "", errors.New("Register failed")
	}

	// generate a new JWT token
	token, err := utils.GenerateNewAccessToken(input.Username, "guest")

	// if token generation failed, return an empty string
	if err != nil {
		return "", err
	}

	// return the JWT token
	return token, nil
}

// Login returns JWT token for authentication
func (u *UserService) Login(ctx context.Context, input model.LoginInput) (string, error) {

	// Set up a connection to the server.
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", os.Getenv("GRPC_USER"),
		os.Getenv("GRPC_USER_PORT")), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("%s, %s, %v", os.Getenv("GRPC_USER"),
			os.Getenv("GRPC_USER_PORT"), err)
		return "", err
	}
	defer conn.Close()
	userClient := pbUser.NewUserServiceClient(conn)
	// create a GetUserRequest
	grpcReq := &pbUser.GetUserRequest{
		Id:       uuid.New().String(),
		Username: input.Username,
	}
	resp, err := userClient.GetUser(ctx, grpcReq)
	if err != nil {
		return "", err
	}

	// compare the user password with the password from the input
	err = bcrypt.CompareHashAndPassword([]byte(resp.Password), []byte(input.Password))

	// If the password does not match, return the empty string
	if err != nil {
		return "", err
	}

	// generate a JWT token
	token, err := utils.GenerateNewAccessToken(resp.Username, resp.RoleType.String())

	// if token generation failed, return the empty string
	if err != nil {
		return "", err
	}
	// return the JWT token
	return token, nil
}

func (u *UserService) EditRole(ctx context.Context, input *model.EditRole) (*model.UserRole, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", os.Getenv("GRPC_USER"),
		os.Getenv("GRPC_USER_PORT")), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("%s, %s, %v", os.Getenv("GRPC_USER"),
			os.Getenv("GRPC_USER_PORT"), err)
		return nil, err
	}
	defer conn.Close()
	userClient := pbUser.NewUserServiceClient(conn)
	grpcReq := &pbUser.SetUserRoleRequest{
		Id:       uuid.New().String(),
		Username: input.Username,
		RoleType: utils.GetRoleType(input.Role),
	}
	log.Println(grpcReq)
	log.Println(pbCommon.RoleType_value[input.Role])
	resp, err := userClient.SetUserRole(ctx, grpcReq)

	if err != nil {
		return nil, err
	}
	if resp.Ok {
		return &model.UserRole{
			Username: input.Username,
			Role:     input.Role,
		}, nil
	}
	return nil, errors.New("edit fail")
}

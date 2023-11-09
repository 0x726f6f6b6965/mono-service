package service

import (
	"context"

	pbCommon "github.com/0x726f6f6b6965/mono-service/protos/common/v1"
	pbUser "github.com/0x726f6f6b6965/mono-service/protos/user/v1"
	"github.com/0x726f6f6b6965/mono-service/services/common/dao/models"
	"github.com/0x726f6f6b6965/mono-service/services/user-service/internal/dao/psql/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userService struct {
	pbUser.UnimplementedUserServiceServer
}

func NewUserService() pbUser.UserServiceServer {
	return &userService{}
}

// GetUser implements v1.UserServiceServer.
func (*userService) GetUser(ctx context.Context, req *pbUser.GetUserRequest) (*pbUser.GetUserResponse, error) {
	userInfo, _ := user.UserObj.GetUser(req.Username)
	if userInfo.ID <= 0 {
		err := status.Error(codes.Unavailable, "username is unavailable")
		return nil, err
	}
	var resp pbUser.GetUserResponse
	resp.Id = req.Id
	resp.Username = userInfo.Username
	resp.Password = userInfo.Password
	resp.RoleType = pbCommon.RoleType(userInfo.RoleType)

	return &resp, nil
}

// SetUser implements v1.UserServiceServer.
func (*userService) SetUser(ctx context.Context, req *pbUser.SetUserRequest) (*pbUser.SetUserResponse, error) {
	exist, _ := user.UserObj.IsUsernameExist(req.Username)

	if exist {
		err := status.Error(codes.Unavailable, "username is unavailable")
		return nil, err
	}
	data := models.User{
		Username: req.Username,
		Password: req.Password,
	}
	id, err := user.UserObj.Add(&data)
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
		return nil, err
	}
	var resp pbUser.SetUserResponse
	resp.Id = req.Id
	resp.Ok = id > 0
	return &resp, nil
}

func (*userService) SetUserRole(ctx context.Context, req *pbUser.SetUserRoleRequest) (*pbUser.SetUserRoleResponse, error) {
	err := user.UserObj.SetUserRole(req.Username, int(pbCommon.RoleType_value[req.RoleType.String()]))
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
		return nil, err
	}
	var resp pbUser.SetUserRoleResponse
	resp.Id = req.Id
	resp.Ok = err == nil
	return &resp, nil
}

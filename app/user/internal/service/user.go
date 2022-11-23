package service

import (
	pb "base/api/user/v1"
	"base/app/user/internal/biz"
	"base/pkg/types"
	"context"
)

type UserService struct {
	pb.UnimplementedUserServer
	user *biz.User
}

func NewUserService(u *biz.User) *UserService {
	return &UserService{user: u}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserInfo, error) {
	save, err := s.user.Save(ctx, &biz.SaveUser{UserName: req.UserName, PassWord: req.PassWord, Phone: req.Phone, Admin: int32(req.Admin.Number())})
	if err != nil {
		return nil, err
	}
	return &pb.UserInfo{
		Id:       int64(save.ID),
		UserName: save.Username,
		Phone:    save.Phone,
		Admin:    save.Admin,
	}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	err := s.user.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserReply{}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	err := s.user.Del(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.DeleteUserReply{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserInfo, error) {
	user, err := s.user.GetUserById(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return user.GenerateToPbUserInfo(), nil
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	users, err := s.user.GetUserPage(ctx, &types.Page{Size: int(req.Size), Index: int(req.Index)})
	if err != nil {
		return nil, err
	}
	res := &pb.ListUserReply{List: make([]*pb.UserInfo, 0, 10)}
	for _, user := range users {
		res.List = append(res.List, user.GenerateToPbUserInfo())
	}

	return res, nil
}
func (s *UserService) GetUserByPhone(ctx context.Context, req *pb.GetUserByPhoneRequest) (*pb.UserInfo, error) {
	user, err := s.user.GetUserByPhone(ctx, req.GetPhone())
	if err != nil {
		return nil, err
	}
	return user.GenerateToPbUserInfo(), nil
}
func (s *UserService) VerifyPassword(ctx context.Context, req *pb.VerifPasswordRequest) (*pb.VerifPasswordReply, error) {
	return s.user.VerifyPassword(ctx, req)
}

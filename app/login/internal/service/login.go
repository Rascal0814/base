package service

import (
	"base/app/login/internal/biz"
	"context"

	pb "base/api/login/v1"
)

type LoginService struct {
	lg *biz.LoginUsecase
	pb.UnimplementedLoginServer
}

func NewLoginService(lg *biz.LoginUsecase) *LoginService {
	return &LoginService{lg: lg}
}

func (s *LoginService) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterReply, error) {
	err := s.lg.Register(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterReply{}, nil
}
func (s *LoginService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginReply, error) {
	token, err := s.lg.Login(ctx, &biz.LoginReq{Phone: req.GetPhone(), PassWord: req.GetPassword()})
	if err != nil {
		return nil, err
	}
	return &pb.LoginReply{Token: token}, nil
}
func (s *LoginService) Logout(ctx context.Context, req *pb.LogoutReq) (*pb.LogoutReply, error) {
	return &pb.LogoutReply{}, nil
}

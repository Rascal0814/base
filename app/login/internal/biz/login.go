package biz

import (
	loginV1 "base/api/login/v1"
	userPb "base/api/user/v1"
	"base/pkg/jwt"
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
)

var (
	// ErrUserNotRegister is user not found.
	ErrUserNotRegister         = errors.NotFound("login", "用户尚未注册")
	ErrUserPassword            = errors.NotFound("login", "密码有误")
	ErrInternalInGenerateToken = errors.NotFound("token", "系统内部错误")
)

// GreeterUsecase is a Greeter usecase.
type LoginUsecase struct {
	log      *log.Helper
	redisCli redis.Cmdable
	user     userPb.UserClient
}

// NewGreeterUsecase new a Greeter usecase.
func NewLoginUsecase(logger log.Logger, uc userPb.UserClient, redisCmd redis.Cmdable) *LoginUsecase {
	return &LoginUsecase{log: log.NewHelper(logger), redisCli: redisCmd, user: uc}
}

func (l *LoginUsecase) Login(ctx context.Context, req *LoginReq) (string, error) {
	user, err := l.user.GetUserByPhone(ctx, &userPb.GetUserByPhoneRequest{Phone: req.Phone})
	if err != nil {
		return "", err
	}

	pass, err := l.user.VerifyPassword(ctx, &userPb.VerifPasswordRequest{Id: int64(user.Id), Password: req.PassWord})
	if err != nil {
		return "", err
	}

	if !pass.Pass {
		return "", ErrUserPassword
	}
	return jwt.JwtEncrypt(int(user.Id), user.UserName)
}

func (l *LoginUsecase) Logout(ctx context.Context, id string) error {
	return l.redisCli.Del(ctx, id).Err()
}

func (l *LoginUsecase) Register(ctx context.Context, req *loginV1.RegisterReq) error {
	_, err := l.user.CreateUser(ctx, &userPb.CreateUserRequest{
		UserName: req.GetUsername(),
		PassWord: req.GetPassword(),
		Admin:    0,
		Phone:    req.GetPhone(),
	})
	return err
}

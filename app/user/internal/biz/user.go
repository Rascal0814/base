package biz

import (
	pb "base/api/user/v1"
	"base/app/user/internal/data/models"
	"base/pkg/password"
	"base/pkg/types"
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound("user", "未找到用户信息")
	ErrUserExist    = errors.InternalServer("user", "手机号码已注册")
	UtilError       = errors.InternalServer("util", "服务内部错误")
)

// UserRepo is a Greater repo.
type UserRepo interface {
	Save(context.Context, *models.User) (*models.User, error)
	Update(context.Context, *models.User) error
	Del(context.Context, int64) error
	FindUserByID(context.Context, int64) (*models.User, error)
	FindUserByPhone(context.Context, string) (*models.User, error)
	PageUser(context.Context, *types.Page) ([]*models.User, error)
	ListAllUser(context.Context) ([]*models.User, error)
}

// User IS A GREETER USECASE.
type User struct {
	repo UserRepo
	log  *log.Helper
}

// NewUse new a Greeter usecase.
func NewUse(repo UserRepo, logger log.Logger) *User {
	return &User{repo: repo, log: log.NewHelper(logger)}
}

// Save ...
func (uc *User) Save(ctx context.Context, g *SaveUser) (*models.User, error) {
	user, err := uc.repo.FindUserByPhone(ctx, g.Phone)
	if err != nil {
		return nil, err
	}
	if user.ID != 0 {
		return nil, ErrUserExist
	}
	return uc.repo.Save(ctx, g.Generate())
}

// GetUserByPhone ...
func (uc *User) GetUserByPhone(ctx context.Context, p string) (*models.User, error) {
	return uc.repo.FindUserByPhone(ctx, p)
}

// GetUserById ...
func (uc *User) GetUserById(ctx context.Context, id int64) (*models.User, error) {
	return uc.repo.FindUserByID(ctx, id)

}

// GetUserPage ...
func (uc *User) GetUserPage(ctx context.Context, page *types.Page) ([]*models.User, error) {
	return uc.repo.PageUser(ctx, page)
}

// VerifyPassword ...
func (uc *User) VerifyPassword(ctx context.Context, req *pb.VerifPasswordRequest) (*pb.VerifPasswordReply, error) {
	var (
		user *models.User
		err  error
	)
	user, err = uc.repo.FindUserByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	if user.PasswordHash == "" {
		return nil, ErrUserNotFound
	}
	return &pb.VerifPasswordReply{Pass: password.CheckPasswordHash(req.GetPassword(), user.PasswordHash)}, nil
}

func (uc *User) Del(ctx context.Context, id int64) error {
	return uc.repo.Del(ctx, id)
}

func (uc *User) Update(ctx context.Context, req *pb.UpdateUserRequest) error {
	var (
		user = &models.User{}
	)
	if req.GetPhone() != "" {
		user, _ = uc.repo.FindUserByPhone(ctx, req.GetPhone())
	}
	if user.ID != 0 {
		return ErrUserExist
	}
	user = &models.User{
		ID:       req.GetId(),
		Username: req.GetUserName(),
		Admin:    int32(req.Admin.Number()),
		Phone:    req.GetPhone(),
	}

	return uc.repo.Update(ctx, user)
}

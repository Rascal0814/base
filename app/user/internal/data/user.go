package data

import (
	"base/app/user/internal/data/models"
	"base/pkg/jwt"
	"base/pkg/types"
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"strconv"

	"base/app/user/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	*Data
	log      *log.Helper
	redisCli redis.Cmdable
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger, redis redis.Cmdable) biz.UserRepo {
	return &userRepo{
		data,
		log.NewHelper(logger),
		redis,
	}
}

func (r *userRepo) Save(ctx context.Context, u *models.User) (*models.User, error) {
	err := r.Orm.Save(u).Error
	return u, err
}

// Update todo 当想把用户admin更改为普通用户时修改不了
func (r *userRepo) Update(ctx context.Context, u *models.User) error {
	r.redisCli.Del(ctx, types.UserInfoRedisPreKey+strconv.Itoa(int(u.ID)))
	return r.Orm.Debug().Updates(u).Error
}

func (r *userRepo) FindUserByID(ctx context.Context, id int64) (*models.User, error) {
	var res = new(models.User)
	// redis
	err := func() error {
		result, err := r.redisCli.Get(ctx, types.UserInfoRedisPreKey+strconv.Itoa(int(id))).Bytes()
		if err != nil {
			return err
		}
		err = json.Unmarshal(result, res)
		if err != nil {
			return err
		}
		return nil
	}()
	if err == nil {
		return res, err
	}
	// redis 没有
	err = r.Orm.First(res, id).Error
	if err != nil {
		return nil, err
	}
	marshal, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	err = r.redisCli.Set(ctx, types.UserInfoRedisPreKey+strconv.Itoa(int(res.ID)), marshal, jwt.ExpiresTime).Err()
	if err != nil {
		r.log.Errorf("set redis error:%s", err)
	}
	return res, err
}

func (r *userRepo) FindUserByPhone(c context.Context, n string) (*models.User, error) {
	res := &models.User{}
	err := r.Orm.Where("phone=?", n).First(res).Error
	return res, err
}

func (r *userRepo) PageUser(c context.Context, req *types.Page) ([]*models.User, error) {
	res := make([]*models.User, 0, 10)
	err := r.Orm.Offset((req.Index - 1) * req.Size).Limit(int(req.Size)).Find(&res).Error
	return res, err
}

func (r *userRepo) ListAllUser(context.Context) ([]*models.User, error) {
	res := make([]*models.User, 0, 10)
	err := r.Orm.Find(&res).Error
	return res, err
}

func (r *userRepo) Del(ctx context.Context, id int64) error {
	r.redisCli.Del(ctx, types.UserInfoRedisPreKey+strconv.Itoa(int(id)))
	return r.Orm.Delete(&models.User{}, id).Error
}

package biz

import (
	"base/app/user/internal/data/models"
	"base/pkg/password"
)

// SaveUser is a Greeter model.
type SaveUser struct {
	UserName string `json:"userName"`
	PassWord string `json:"passWord"`
	Phone    string `json:"phone"`
	Admin    int32  `json:"admin"`
}

// Generate 生成数据库实例
func (s SaveUser) Generate() *models.User {
	password, _ := password.HashPassword(s.PassWord)
	return &models.User{
		Username:     s.UserName,
		PasswordHash: password,
		Admin:        s.Admin,
		Phone:        s.Phone,
	}
}

package models

import (
	pb "base/api/user/v1"
	"database/sql"
	"time"
)

// User is the model entity for the User schema.
type User struct {
	ID        int64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
	// Username holds the value of the "username" field.
	Username string `json:"user_name"`
	// PasswordHash holds the value of the "password_hash" field.
	PasswordHash string `json:"password_hash"`
	// admin 0:normal 1:admin 2:super_admin
	Admin int32 `json:"admin"gorm:"size:1"`
	// phone number must input
	Phone string `json:"phone"`
}

func (u *User) GenerateToPbUserInfo() *pb.UserInfo {
	return &pb.UserInfo{
		Id:       int64(u.ID),
		UserName: u.Username,
		Admin:    u.Admin,
		Phone:    u.Phone,
	}
}

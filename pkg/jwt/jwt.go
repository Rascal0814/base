package jwt

import (
	"github.com/go-kratos/kratos/v2/errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	BearerWord string = "Bearer"

	// AuthorizationKey holds the key used to store the JWT Token in the request tokenHeader.
	AuthorizationKey string = "Authorization"

	// reason holds the error reason.
	reason string = "UNAUTHORIZED"

	ExpiresTime = 7 * 24 * time.Hour
)
var (
	ErrMissingJwtToken        = errors.Unauthorized(reason, "JWT token is missing")
	ErrMissingKeyFunc         = errors.Unauthorized(reason, "keyFunc is missing")
	ErrTokenInvalid           = errors.Unauthorized(reason, "Token is invalid")
	ErrTokenExpired           = errors.Unauthorized(reason, "JWT token has expired")
	ErrTokenParseFail         = errors.Unauthorized(reason, "Fail to parse JWT token ")
	ErrUnSupportSigningMethod = errors.Unauthorized(reason, "Wrong signing method")
	ErrWrongContext           = errors.Unauthorized(reason, "Wrong context for middleware")
	ErrNeedTokenProvider      = errors.Unauthorized(reason, "Token provider is missing")
	ErrSignToken              = errors.Unauthorized(reason, "Can not sign token.Is the key correct?")
	ErrGetKey                 = errors.Unauthorized(reason, "Can not get key while signing token")
)

type MyClaims struct {
	UserId   int    `json:"userId"`
	UserName string `json:"userName"`
	jwt.StandardClaims
}

var mySiginKey = []byte("Hml@200814")

// JwtEncrypt 加密
func JwtEncrypt(uid int, Name string) (string, error) {
	mc := MyClaims{
		UserId:   uid,
		UserName: Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ExpiresTime).Unix(), //过期时间
			NotBefore: time.Now().Unix(),                  //开始生效
			Issuer:    "Hml",                              //签发人
			Subject:   "user Token",                       //主题
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, mc)

	// 加密
	s, err := t.SignedString(mySiginKey)
	if err != nil {
		return "", err
	}
	return s, nil
}

// JwtDecrypt 解密
func JwtDecrypt(s string) (*jwt.Token, MyClaims, error) {
	claims := &MyClaims{}
	token, err := jwt.ParseWithClaims(s, claims, func(t *jwt.Token) (interface{}, error) {
		return mySiginKey, nil
	})

	return token, *claims, err
}

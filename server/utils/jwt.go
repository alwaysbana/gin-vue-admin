package utils

import (
	"gin-vue-admin/server/initialize/global"
	"gin-vue-admin/server/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(user model.User) (string, error) {
	//创建 customerclaims
	claims := CustomClaims{
		ID:       user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    global.GVA_CONFIG.Jwt.Issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(global.GVA_CONFIG.Jwt.Expiration) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedStr, err := token.SignedString([]byte(global.GVA_CONFIG.Jwt.Secret))

	return signedStr, err

}

func ParseToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&CustomClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(global.GVA_CONFIG.Jwt.Secret), nil
		},
	)

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}

package utils

import (
	"ByteDance_Douyin/config"
	"ByteDance_Douyin/model"
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

var jwtKey = []byte(config.C.JWT.SecretKey)

type Claims struct {
	UserId int64
	jwt.StandardClaims
}

//生成token并返回
func SignToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(7 * time.Hour)
	claims := &Claims{
		UserId: user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "douyin",
			Subject:   "user token",
		}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}

// 解析token并返回claims
func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return token.Claims.(*Claims), err
}

// JWTAuth 用于验证token，并返回token对应的userid
func AuthToken(token string) (int64, error) {
	if token == "" {
		return 0, errors.New("token为空")
	}
	claim, err := ParseToken(token)
	if err != nil {
		return 0, errors.New("token过期")
	}
	return claim.UserId, nil
}

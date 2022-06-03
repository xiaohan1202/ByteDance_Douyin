package utils

import (
	"ByteDance_Douyin/config"
	"github.com/golang-jwt/jwt"
	"time"
)

var jwtKey = []byte(config.C.JWT.SecretKey)

type Claims struct {
	UserId int64
	jwt.StandardClaims
}

//生成token并返回
func SignToken(userId int64) (string, error) {
	expirationTime := time.Now().Add(7 * time.Hour)
	claims := &Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}

// 解析token并返回claims
func ParseToken(tokenString string) (int64, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return -1, err
	}
	return claims.UserId, err
}

//// JWTAuth 用于验证token，并返回token对应的userid
//func AuthToken(token string) (int64, error) {
//	if token == "" {
//		return 0, errors.New("token为空")
//	}
//	claim, err := ParseToken(token)
//	if err != nil {
//		return 0, errors.New("token过期")
//	}
//	return claim.UserId, nil
//}

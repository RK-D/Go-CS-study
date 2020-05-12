package common

import (
	"github.com/dgrijalva/jwt-go"
	"pers.study/cstest/model"
	"time"
)

var jwtKey = []byte("a_secret_correct")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(3 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "per/RK-D",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

//token : 协议头(加密协议) +   claims信息 +前面两者加上jwtKey hash的值
//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjMsImV4cCI6MTU4OTQyNTYyNywiaWF0IjoxNTg5MTY2NDI3LCJpc3MiOiJwZXIvUkstRCIsInN1YiI6InVzZXIgdG9rZW4ifQ.urwn-b5rthHW43FKHO3sX_hpzI6ILAU-a_NWVOP8p1k
//

//解析token的函数  tokenString里面解析出claims 返回
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
	return token, claims, err
}

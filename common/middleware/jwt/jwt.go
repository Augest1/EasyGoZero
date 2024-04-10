package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Jwt struct {
	AccessSecret string
	AccessExpire int64
}

func NewJwt(accessSecret string, accessExpire int64) *Jwt {
	return &Jwt{
		AccessSecret: accessSecret,
		AccessExpire: accessExpire,
	}
}

type Payload struct {
	UserId int64 `json:"userId"`
}

func (j *Jwt) GenJwtToken(payload *Payload) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Unix() + j.AccessExpire
	claims["iat"] = time.Now().Unix()
	claims["userId"] = payload.UserId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(j.AccessSecret))
}

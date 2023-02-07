// @author cold bin
// @date 2023/2/7

package _jwt

import (
	"eliminate-male-appearance-anxiety/global"
	"errors"
	"github.com/cold-bin/goutil/general/mlog"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type SmaaClaims struct {
	UserId uint `json:"user_id,omitempty"`
	jwt.RegisteredClaims
}

var secret = "cold-bin is gopher~"

func CreateToken(userId uint) (at string, rt string, err error) {
	// 签发at
	tokenA := jwt.NewWithClaims(jwt.SigningMethodHS256, SmaaClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:   "cold bin",
			Subject:  "smaa app",
			Audience: []string{"smaa app"},
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(global.Set.App.ExpAt) * time.Hour)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        "",
		},
	})
	at, err = tokenA.SignedString(secret)
	if err != nil {
		return
	}
	// 签发rt
	tokenR := jwt.NewWithClaims(jwt.SigningMethodHS256, SmaaClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(global.Set.App.ExpRt) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "cold bin",
		},
	})
	rt, err = tokenR.SignedString(secret)
	if err != nil {
		return
	}

	mlog.Debugf("签发access_token: %s，refresh_token: %s", at, rt)
	return
}

var ErrInvalidToken = errors.New("token is invalid")

func ParesToken(tokenStr string) (*SmaaClaims, error) {
	c, err := jwt.ParseWithClaims(tokenStr, &SmaaClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return nil, err
	}
	if claims, ok := c.Claims.(*SmaaClaims); ok && c.Valid { // 校验token
		return claims, nil
	}

	return nil, ErrInvalidToken
}

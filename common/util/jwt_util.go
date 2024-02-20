package util

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

var JwtUtil = &jwtUtil{}

const secret = "abc"

type Payload struct {
	Uid        int64
	AppId      int
	DeviceType int
	Cts        int64
	jwt.StandardClaims
}

type jwtUtil struct {
}

func NewPayload(uid int64, appId int, deviceType int) *Payload {
	return &Payload{
		Uid:        uid,
		AppId:      appId,
		DeviceType: deviceType,
		Cts:        time.Now().UnixMilli(),
	}
}

func (*jwtUtil) Encode(se string, payload *Payload) (result string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	if len(se) <= 0 {
		result, err = token.SignedString([]byte(secret))
	} else {
		result, err = token.SignedString([]byte(se))
	}

	if err != nil {
		logger.Errorf("JwtUtil Encode failed:%v", err)
	}
	return result, err
}

func (*jwtUtil) Parse(se string, tokenString string) (*Payload, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Payload{}, func(token *jwt.Token) (i interface{}, err error) {
		if len(se) > 0 {
			return []byte(se), nil
		} else {
			return []byte(secret), nil
		}
	})

	if err != nil {
		logger.Errorf("JwtUtil Parse failed:%v", err)
		return nil, err
	}
	if claims, ok := token.Claims.(*Payload); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("Invalid token")
}

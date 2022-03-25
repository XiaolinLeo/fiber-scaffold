package utils

import (
	"log"
	"time"

	"github.com/gbrlsnchs/jwt/v3"
)

type CustomPayload struct {
	jwt.Payload
	Username string `json:"username,omitempty"`
}

var hs = jwt.NewHS256([]byte("secret"))

//签名 generate token
func Sign(username string) (string, error) {
	now := time.Now()
	pl := CustomPayload{
		Payload: jwt.Payload{
			ExpirationTime: jwt.NumericDate(now.Add(time.Second * 1200)),
			IssuedAt:       jwt.NumericDate(now),
		},
		Username: username,
	}

	token, err := jwt.Sign(pl, hs)
	if err != nil {
		log.Println("generate token failed")
	}

	return string(token), err
}

//解密验证token
func Verity(token []byte) (*CustomPayload, error) {
	payload := &CustomPayload{}
	var (
		now             = time.Now()
		expValidator    = jwt.ExpirationTimeValidator(now)
		iatValidator    = jwt.IssuedAtValidator(now)
		validatePayload = jwt.ValidatePayload(&payload.Payload, expValidator, iatValidator)
	)
	_, err := jwt.Verify(token, hs, payload, validatePayload)
	if err != nil {
		return nil, err
	}

	return payload, err
}

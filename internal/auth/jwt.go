package auth

import (
	"github.com/glide-im/api/internal/api/comm"
	"github.com/golang-jwt/jwt"
	"time"
)

var jwtSecret []byte

func genJwt(payload JwtClaims) (string, error) {
	expireAt := time.Now().Add(time.Hour * 24)
	return genJwtExp(payload, expireAt)
}

func genJwtExp(payload JwtClaims, expiredAt time.Time) (string, error) {
	payload.ExpiresAt = expiredAt.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	t, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return t, nil
}

func parseJwt(token string) (*JwtClaims, error) {
	j := JwtClaims{}
	t, err := jwt.ParseWithClaims(token, &j, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	jwtToken, ok := t.Claims.(*JwtClaims)
	if !ok {
		return nil, comm.NewApiBizError(1, "invalid token")
	}
	return jwtToken, nil
}

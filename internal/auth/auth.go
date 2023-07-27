package auth

import (
	"crypto/sha512"
	"github.com/glide-im/api/internal/config"
	"github.com/glide-im/glide/pkg/gate"
	"github.com/golang-jwt/jwt"
)

var (
	GUEST_DEVICE   = int64(-1)
	MOBILE_DEVICE  = int64(2)
	DEFAULT_DEVICE = int64(1)
)

func SetJwtSecret(secret []byte) {
	jwtSecret = secret
}

type JwtClaims struct {
	jwt.StandardClaims
	Uid    int64 `json:"uid"`
	Device int64 `json:"device"`
	Ver    int64 `json:"ver"`
	AppId  int64 `json:"app_id"`
}

func ParseToken(token string) (*JwtClaims, error) {
	result, err := parseJwt(token)
	if err != nil {
		return nil, err
	}
	return &JwtClaims{
		Uid:    result.Uid,
		Device: result.Device,
		AppId:  result.AppId,
	}, nil
}

func GenerateTokenExpire(uid int64, device int64, expireHour int64) (string, error) {
	return genJwt(JwtClaims{
		StandardClaims: jwt.StandardClaims{},
		Uid:            uid,
		Device:         device,
		Ver:            1,
		AppId:          1,
	}, expireHour)
}

func GenerateCredentials(c *gate.ClientAuthCredentials) (*gate.EncryptedCredential, error) {

	key := sha512.New().Sum([]byte(config.ApiHttp.IMServerSecret))
	cbcCrypto := gate.NewAesCBCCrypto(key)
	encryptCredentials, err := cbcCrypto.EncryptCredentials(c)
	if err != nil {
		return nil, err
	}

	enc := gate.EncryptedCredential{
		Version:    1,
		Credential: string(encryptCredentials),
	}
	return &enc, nil
}

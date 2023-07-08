package auth

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenJwt(t *testing.T) {

	jwt, err := genJwt(JwtClaims{
		Uid:    1,
		Device: 1,
		Ver:    1,
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Log(jwt)
	}
}

func TestParseJwt(t *testing.T) {
	jwt, err := parseJwt("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODg3NDI3NDYsInVpZCI6IjExMSIsImRldmljZSI6IjIyIiwidmVyIjoxNjg4NzM5MTQ2fQ.eJ2DEQdBCoB03lAmA7Astnrn4B2WIVpQkvxiumxu9Ic")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(jwt)
	}
}

func TestGenerateTokenExpire(t *testing.T) {
	token, err := GenerateTokenExpire(111, 22, 1)
	assert.Nil(t, err)
	t.Log(token)
}

func TestParseToken(t *testing.T) {
	token, err := ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODg4MjU4MzMsInVpZCI6MTExLCJkZXZpY2UiOjIyLCJ2ZXIiOjEsImFwcF9pZCI6MX0.4q7dbE-wHzl-dTUiT4FI0QFYcZ3jyWsI4Y8WkZiEBX0")
	assert.Nil(t, err)
	t.Log(token)
}

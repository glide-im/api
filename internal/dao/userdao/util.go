package userdao

import (
	"golang.org/x/crypto/bcrypt"
	"math/rand"
)

var (
	table = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func genToken(length int) string {
	res := ""
	for i := 0; i < length; i++ {
		idx := rand.Intn(62)
		res = res + table[idx:idx+1]
	}

	return res
}

func PasswordHash(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return password
	}
	return string(bytes)
}

func PasswordVerify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

package pkg

import (
	"golang.org/x/crypto/bcrypt"
	"math/rand/v2"
)

var passwordRunes = []rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890@")

func GeneratePassword(n int) string {
	res := make([]rune, n)

	for i := range string(res) {
		res[i] = passwordRunes[rand.IntN(len(passwordRunes))]
	}

	return string(res)
}

func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic("Failed to hash password: " + err.Error())
	}

	return string(hashedPassword)
}

func ComparePasswords(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}

	return true
}

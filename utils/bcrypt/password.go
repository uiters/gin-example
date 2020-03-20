package bcrypt

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string  {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Print(err)
	}
	return string(hashedPassword)
}

func ComparePasswordAndHashedPassword(password, hashedPassword string) error  {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err
}
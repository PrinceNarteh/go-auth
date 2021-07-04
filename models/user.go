package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id       uint
	Name     string
	Email    string `gorm:"unique"`
	Password string
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ComparePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

package entity

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int
	Name         string
	Username     string
	HashPassword string
}

func (u User) CompareUsernameAndPass(username string, password string) bool {
	if u.Username != username {
		return false
	}

	return bcrypt.CompareHashAndPassword([]byte(u.HashPassword), []byte(password)) == nil
}

func (f factoryImpl) NewUser(name, username, password string) (*User, error) {
	if len(username)*len(password) != 0 {
		return nil, errors.New("create user: username and password must not be empty")
	}

	hashPass, err := hashBcrypt(password)
	if err != nil {
		return nil, err
	}

	return &User{
		Name:         name,
		Username:     username,
		HashPassword: hashPass,
	}, nil
}

func hashBcrypt(in string) (string, error) {
	out, err := bcrypt.GenerateFromPassword([]byte(in), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.WithMessage(err, "hash bcrypt")
	}

	return string(out), nil
}

package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"github.com/thanhpp/scm/internal/scmsrv/domain/entity"
)

const (
	expire = time.Hour * 24
	issuer = "scmsrv"
	secret = "secret"
)

type Token struct {
	JWT     string
	Refresh string
}

type ClaimsUserData struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type Claims struct {
	jwt.StandardClaims
	User ClaimsUserData `json:"user"`
}

func (c *Claims) Set(user *entity.User) {
	c.User.ID = user.ID
	c.User.Name = user.Name
	c.User.Username = user.Username
}

type JWTSrv interface {
	GenToken(user *entity.User) (*Token, error)
	Validate(token string) (*jwt.Token, error)
	Refresh(user *entity.User, refresh string) (*Token, error)
}

type jwtSrvImpl struct {
}

func NewJWTSrvImpl() JWTSrv {
	return &jwtSrvImpl{}
}

func (f jwtSrvImpl) GenToken(user *entity.User) (*Token, error) {
	claims := &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expire).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    issuer,
		},
	}
	claims.Set(user)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(secret)
	if err != nil {
		return nil, err
	}

	return &Token{
		JWT: t,
	}, nil
}
func (f jwtSrvImpl) Validate(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token: signing method")
		}

		return secret, nil
	})
}

func (f jwtSrvImpl) Refresh(user *entity.User, refresh string) (*Token, error) {
	return nil, nil
}

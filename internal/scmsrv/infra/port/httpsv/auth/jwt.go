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
	secret = "secretsecretsecretsecret"
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
	CheckCache(userID int) bool
}

type jwtSrvImpl struct {
	tokenCache map[int]*Token // userID - jwtToken
}

func NewJWTSrvImpl() JWTSrv {
	return &jwtSrvImpl{
		tokenCache: make(map[int]*Token),
	}
}

func (f *jwtSrvImpl) GenToken(user *entity.User) (*Token, error) {
	cachedToken, ok := f.tokenCache[user.ID]
	if ok {
		return cachedToken, nil
	}

	claims := &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expire).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    issuer,
		},
	}
	claims.Set(user)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, errors.WithMessage(err, "signing token")
	}

	newToken := &Token{
		JWT: t,
	}
	f.tokenCache[user.ID] = newToken

	return newToken, nil
}
func (f jwtSrvImpl) Validate(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token: signing method")
		}

		return []byte(secret), nil
	})
}

func (f jwtSrvImpl) Refresh(user *entity.User, refresh string) (*Token, error) {
	return nil, nil
}

func (f *jwtSrvImpl) CheckCache(userID int) bool {
	_, ok := f.tokenCache[userID]

	return ok
}

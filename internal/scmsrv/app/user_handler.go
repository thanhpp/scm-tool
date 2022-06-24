package app

import (
	"context"
	"errors"

	"github.com/thanhpp/scm/internal/scmsrv/domain/entity"
	"github.com/thanhpp/scm/internal/scmsrv/domain/repo"
)

type UserHandler struct {
	f        entity.Factory
	userRepo repo.UserRepo
}

func (h UserHandler) CreateUser(
	ctx context.Context, name, username, password string,
) (*entity.User, error) {
	newUser, err := h.f.NewUser(name, username, password)
	if err != nil {
		return nil, err
	}

	if err := h.userRepo.NewUser(ctx, newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}

func (h UserHandler) ValidateUser(
	ctx context.Context, username, pass string,
) (*entity.User, error) {
	user, err := h.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	if !user.CompareUsernameAndPass(username, pass) {
		return nil, errors.New("invalid username/password")
	}

	return user, nil
}

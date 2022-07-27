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
	user, err := h.userRepo.GetByUsername(ctx, username)
	if err == nil {
		if user != nil {
			return nil, errors.New("create user: duplicate username")
		}
	}

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

func (h UserHandler) GetUsers(
	ctx context.Context, limit, offset int,
) ([]*entity.User, error) {
	return h.userRepo.GetUsers(ctx, repo.GetUsersFilter{
		Limit:  limit,
		Offset: offset,
	})
}

func (h UserHandler) UpdateUserPassword(
	ctx context.Context, id int, newPass string,
) error {
	return h.userRepo.UpdateUserByID(ctx, id,
		func(ctx context.Context, u entity.User) (entity.User, error) {
			if err := u.UpdatePass(newPass); err != nil {
				return entity.User{}, err
			}

			return u, nil
		})
}

package storage

import (
	"context"

	"github.com/thanhpp/scm/internal/scmsrv/domain/entity"
	"github.com/thanhpp/scm/internal/scmsrv/domain/repo"
	"gorm.io/gorm"
)

type UserDB struct {
	gdb *gorm.DB
}

func (d DB) UserDB() *UserDB {
	return &UserDB{
		gdb: d.gdb,
	}
}

func (d UserDB) NewUser(ctx context.Context, user *entity.User) error {
	userDB := marshalUser(user)

	if err := d.gdb.WithContext(ctx).Model(&repo.User{}).Create(&userDB).Error; err != nil {
		return err
	}

	user.ID = userDB.ID

	return nil
}

func (d UserDB) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	userDB := &repo.User{}

	if err := d.gdb.WithContext(ctx).Where("username LIKE ?", username).First(&userDB).Error; err != nil {
		return nil, err
	}

	return unmarshalUser(userDB), nil
}

func (d UserDB) GetUsers(ctx context.Context, filer repo.GetUsersFilter) ([]*entity.User, error) {
	var usersDB []*repo.User

	if err := d.gdb.WithContext(ctx).Model(&repo.User{}).
		Offset(filer.Offset).Limit(filer.Limit).Order("id ASC").
		Find(&usersDB).
		Error; err != nil {
		return nil, err
	}

	users := make([]*entity.User, len(usersDB))
	for i := range users {
		users[i] = unmarshalUser(usersDB[i])
	}

	return users, nil
}

func (d UserDB) UpdateUserByID(ctx context.Context, id int, fn repo.UpdateUserFn) error {
	return d.gdb.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		userDB := new(repo.User)

		if err := tx.WithContext(ctx).Model(&repo.User{}).Where("id = ?", id).Take(userDB).
			Error; err != nil {
			return err
		}

		user := unmarshalUser(userDB)
		newUser, err := fn(ctx, *user)
		if err != nil {
			return err
		}
		newUserDB := marshalUser(&newUser)

		if err := tx.WithContext(ctx).Model(&repo.User{}).Where("id = ?", id).Updates(newUserDB).
			Error; err != nil {
			return err
		}

		return nil
	})
}

func marshalUser(in *entity.User) *repo.User {
	return &repo.User{
		ID:           in.ID,
		Name:         in.Name,
		Username:     in.Username,
		HashPassword: in.HashPassword,
	}
}

func unmarshalUser(in *repo.User) *entity.User {
	return &entity.User{
		ID:           in.ID,
		Name:         in.Name,
		Username:     in.Username,
		HashPassword: in.HashPassword,
	}
}

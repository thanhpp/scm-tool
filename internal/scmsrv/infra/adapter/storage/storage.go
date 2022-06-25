package storage

import (
	"context"

	"github.com/thanhpp/scm/internal/scmsrv/domain/entity"
	"github.com/thanhpp/scm/internal/scmsrv/domain/repo"
	"gorm.io/gorm"
)

type StorageDB struct {
	gdb *gorm.DB
}

func (d DB) StorageDB() *StorageDB {
	return &StorageDB{
		gdb: d.gdb,
	}
}

func (s StorageDB) marshal(in *entity.Storage) *repo.Storage {
	storage := &repo.Storage{
		ID:       in.ID,
		Name:     in.Name,
		Location: in.Location,
	} // ! Desc

	return storage
}

func unmarshalStorage(in *repo.Storage) *entity.Storage {
	storage := &entity.Storage{
		ID:       in.ID,
		Name:     in.Name,
		Location: in.Location,
		// ! Desc
	}

	return storage
}

func (s StorageDB) Get(ctx context.Context, id int) (*entity.Storage, error) {
	storageDB := new(repo.Storage)
	if err := s.gdb.WithContext(ctx).Where("id = ?", id).Take(storageDB).Error; err != nil {
		return nil, err
	}

	return unmarshalStorage(storageDB), nil
}

func (s StorageDB) Create(ctx context.Context, storage *entity.Storage) error {
	storageDB := s.marshal(storage)

	if err := s.gdb.WithContext(ctx).Create(storageDB).Error; err != nil {
		return err
	}

	storage.ID = storageDB.ID

	return nil
}

func (s StorageDB) GetList(ctx context.Context, filter repo.StorageFiler) ([]*entity.Storage, error) {
	storagesDB := make([]*repo.Storage, 0, filter.Limit)

	if err := s.gdb.WithContext(ctx).Model(&repo.Storage{}).
		Offset(filter.Offset).Limit(filter.Limit).
		Find(&storagesDB).Error; err != nil {
		return nil, err
	}

	storages := make([]*entity.Storage, len(storagesDB))
	for i := range storages {
		storages[i] = unmarshalStorage(storagesDB[i])
	}

	return storages, nil
}

func (s StorageDB) Update(ctx context.Context, storageID int, fn repo.StorageUpdateFn) error {
	return s.gdb.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		storage, err := s.Get(ctx, storageID)
		if err != nil {
			return err
		}

		newStorage, err := fn(storage)
		if err != nil {
			return err
		}

		if err := tx.WithContext(ctx).
			Model(&repo.Storage{}).
			Where("id = ?", storageID).
			Updates(newStorage).Error; err != nil {
			return err
		}

		return nil
	})
}

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

	storage.Serials = make([]repo.Serial, 0, len(in.Serials))
	for i := range in.Serials {
		storage.Serials = append(storage.Serials, marshalSerial(in.Serials[i]))
	}

	return storage
}

func (StorageDB) unmarshal(in *repo.Storage) *entity.Storage {
	storage := &entity.Storage{
		ID:       in.ID,
		Name:     in.Name,
		Location: in.Location,
		// ! Desc
	}

	storage.Serials = make([]entity.Serial, 0, len(in.Serials))
	for i := range in.Serials {
		storage.Serials = append(storage.Serials, unmarshalSeri(
			in.Serials[i],
		))
	}

	return storage
}

func (s StorageDB) Get(ctx context.Context, id int) (*entity.Storage, error) {
	storageDB := new(repo.Storage)
	if err := s.gdb.WithContext(ctx).Where("id = ?", id).Take(storageDB).Error; err != nil {
		return nil, err
	}

	return s.unmarshal(storageDB), nil
}

func (s StorageDB) Create(ctx context.Context, storage *entity.Storage) error {
	storageDB := s.marshal(storage)

	if err := s.gdb.WithContext(ctx).Create(storageDB).Error; err != nil {
		return err
	}

	storage.ID = storageDB.ID

	return nil
}

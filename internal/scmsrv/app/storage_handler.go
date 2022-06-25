package app

import (
	"context"

	"github.com/thanhpp/scm/internal/scmsrv/domain/entity"
	"github.com/thanhpp/scm/internal/scmsrv/domain/repo"
)

type StorageHandler struct {
	fac         entity.Factory
	storageRepo repo.StorageRepo
}

func (h StorageHandler) Create(ctx context.Context, name, desc, location string) (*entity.Storage, error) {
	newStorage, err := h.fac.NewStorage(name, desc, location)
	if err != nil {
		return nil, err
	}

	if err := h.storageRepo.Create(ctx, newStorage); err != nil {
		return nil, err
	}

	return newStorage, nil
}

func (h StorageHandler) GetListStorages(ctx context.Context, page, size int) ([]*entity.Storage, error) {
	offset, limit := genOffsetLimit(page, size)

	return h.storageRepo.GetList(ctx, repo.StorageFiler{
		Offset: offset,
		Limit:  limit,
	})
}

func (h StorageHandler) UpdateStorage(ctx context.Context, id int, name, desc, location string) error {
	return h.storageRepo.Update(ctx, id, func(s *entity.Storage) (*entity.Storage, error) {
		if len(name) != 0 {
			s.Name = name
		}

		if len(desc) != 0 {
			s.Desc = desc
		}

		if len(location) != 0 {
			s.Location = location
		}

		return s, nil
	})
}

func genOffsetLimit(page, size int) (offset, limit int) {
	offset = (page - 1) * size
	limit = size

	return
}

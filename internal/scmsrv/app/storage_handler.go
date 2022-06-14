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

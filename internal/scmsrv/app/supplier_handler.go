package app

import (
	"context"

	"github.com/thanhpp/scm/internal/scmsrv/domain/entity"
	"github.com/thanhpp/scm/internal/scmsrv/domain/repo"
)

type SupplierHandler struct {
	fac          entity.Factory
	supplierRepo repo.SupplierRepo
}

func (h SupplierHandler) GetByID(ctx context.Context, id int) (*entity.Supplier, error) {
	return h.supplierRepo.Get(ctx, id)
}

func (h SupplierHandler) GetList(ctx context.Context, page, size int) ([]*entity.Supplier, error) {
	offset, limit := genOffsetLimit(page, size)

	return h.supplierRepo.GetList(ctx, repo.SupplierFiler{
		Offset: offset,
		Limit:  limit,
	})
}

func (h SupplierHandler) Create(ctx context.Context, name, email, phone string) (*entity.Supplier, error) {
	newSupplier, err := h.fac.NewSupplier(name, email, phone)
	if err != nil {
		return nil, err
	}

	if err := h.supplierRepo.Create(ctx, newSupplier); err != nil {
		return nil, err
	}

	return newSupplier, nil
}

func (h SupplierHandler) Update(ctx context.Context, id int, name, email, phone string) error {
	oldSupplier, err := h.supplierRepo.Get(ctx, id)
	if err != nil {
		return err
	}

	if name != "" {
		oldSupplier.Name = name
	}

	if email != "" {
		oldSupplier.Email = email
	}

	if phone != "" {
		oldSupplier.Phone = phone
	}

	return h.supplierRepo.Update(ctx, oldSupplier)
}

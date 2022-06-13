package app

import (
	"context"

	"github.com/thanhpp/scm/internal/scmsrv/domain/entity"
	"github.com/thanhpp/scm/internal/scmsrv/domain/repo"
)

type SupplierHanlder struct {
	fac          entity.Factory
	supplierRepo repo.SupplierRepo
}

func (h SupplierHanlder) Create(ctx context.Context, name, email, phone string) (*entity.Supplier, error) {
	newSupplier := h.fac.NewSupplier(name, email, phone)

	if err := h.supplierRepo.Create(ctx, newSupplier); err != nil {
		return nil, err
	}

	return newSupplier, nil
}

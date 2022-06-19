package storage

import (
	"context"

	"github.com/thanhpp/scm/internal/scmsrv/domain/entity"
	"github.com/thanhpp/scm/internal/scmsrv/domain/repo"
	"gorm.io/gorm"
)

type SupplierDB struct {
	gdb *gorm.DB
}

func (d DB) SupplierDB() *SupplierDB {
	return &SupplierDB{
		gdb: d.gdb,
	}
}

func marshalSupplier(in *entity.Supplier) *repo.Supplier {
	return &repo.Supplier{
		ID:    in.ID,
		Name:  in.Name,
		Email: in.Email,
		Phone: in.Phone,
	}
}

func unmarshalSupplier(in *repo.Supplier) *entity.Supplier {
	return &entity.Supplier{
		ID:    in.ID,
		Name:  in.Name,
		Email: in.Email,
		Phone: in.Phone,
	}
}

func (d SupplierDB) Get(ctx context.Context, id int) (*entity.Supplier, error) {
	supplierDB := new(repo.Supplier)
	if err := d.gdb.WithContext(ctx).Where("id = ?", id).Take(supplierDB).Error; err != nil {
		return nil, err
	}

	return unmarshalSupplier(supplierDB), nil
}

func (d SupplierDB) Create(ctx context.Context, supplier *entity.Supplier) error {
	supplierDB := marshalSupplier(supplier)

	if err := d.gdb.WithContext(ctx).Create(supplierDB).Error; err != nil {
		return err
	}

	supplier.ID = supplierDB.ID

	return nil
}

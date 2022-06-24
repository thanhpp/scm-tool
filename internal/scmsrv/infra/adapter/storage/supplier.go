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

func (d SupplierDB) GetList(ctx context.Context, filter repo.SupplierFiler) ([]*entity.Supplier, error) {
	supplierDBs := make([]*repo.Supplier, filter.Limit)

	if err := d.gdb.WithContext(ctx).Model(&repo.Supplier{}).
		Limit(filter.Limit).Offset(filter.Offset).
		Find(&supplierDBs).Error; err != nil {
		return nil, err
	}

	suppliers := make([]*entity.Supplier, len(supplierDBs))

	for i := range suppliers {
		suppliers[i] = unmarshalSupplier(supplierDBs[i])
	}

	return suppliers, nil
}

func (d SupplierDB) Update(ctx context.Context, suplier *entity.Supplier) error {
	suplierDB := marshalSupplier(suplier)
	return d.gdb.WithContext(ctx).Updates(suplierDB).Error
}

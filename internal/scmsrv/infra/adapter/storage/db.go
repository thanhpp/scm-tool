package storage

import (
	"github.com/thanhpp/scm/internal/scmsrv/domain/entity"
	"github.com/thanhpp/scm/internal/scmsrv/domain/repo"
	"gorm.io/gorm"
)

var (
	collection []interface{} = []interface{}{
		repo.Item{}, repo.ItemType{}, repo.Storage{}, repo.Serial{},
		repo.Supplier{}, repo.ImportTicket{},
		repo.ImportTicketBillImage{}, repo.ImportTicketProductImage{}, repo.ImportTicketDetails{},
	}
)

type DB struct {
	gdb *gorm.DB
}

func NewDB(gormDB *gorm.DB) *DB {
	return &DB{
		gdb: gormDB,
	}
}

func (d DB) AutoMigrate() error {
	return d.gdb.AutoMigrate(
		collection...,
	)
}

func (d DB) DropAll() error {
	return d.gdb.Migrator().DropTable(
		collection...,
	)
}

func (d DB) marshalImportTicket(in entity.ImportTicket) *repo.ImportTicket {
	out := &repo.ImportTicket{
		ID:             in.ID,
		FromSupplierID: in.FromSupplier.ID,
		ToStorageID:    in.ToStorage.ID,
		Status:         in.Status,
		SendTime:       in.SendTime,
		ReceiveTime:    in.ReceiveTime,
		Fee:            in.Fee,
	}

	out.BillImages = make([]repo.ImportTicketBillImage, 0, len(in.BillImagePaths))
	for i := range in.BillImagePaths {
		out.BillImages = append(
			out.BillImages,
			repo.ImportTicketBillImage{
				ImportTicketID: in.ID,
				BillImagePath:  in.BillImagePaths[i],
			})
	}

	out.ProductImages = make([]repo.ImportTicketProductImage, 0, len(in.ProductImagePaths))
	for i := range in.ProductImagePaths {
		out.ProductImages = append(
			out.ProductImages,
			repo.ImportTicketProductImage{
				ImportTicketID:   in.ID,
				ProductImagePath: in.ProductImagePaths[i],
			},
		)
	}

	out.Details = make([]repo.ImportTicketDetails, 0, len(in.Details))
	for i := range in.Details {
		out.Details = append(
			out.Details,
			repo.ImportTicketDetails{
				ImportTicketID:  in.ID,
				ItemSKU:         in.Details[i].Item.SKU,
				BuyQuantity:     in.Details[i].BuyQuantity,
				ReceiveQuantity: in.Details[i].ReceiveQuantity,
				BuyPrice:        in.Details[i].BuyPrice,
			})
	}

	return out
}

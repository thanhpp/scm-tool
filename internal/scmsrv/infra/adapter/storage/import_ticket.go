package storage

import (
	"context"
	"log"

	"github.com/thanhpp/scm/internal/scmsrv/domain/entity"
	"github.com/thanhpp/scm/internal/scmsrv/domain/repo"
	"github.com/thanhpp/scm/pkg/logger"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ImportTicketDB struct {
	gdb *gorm.DB
}

func (d DB) ImportTicketDB() *ImportTicketDB {
	return &ImportTicketDB{
		gdb: d.gdb,
	}
}

func (d ImportTicketDB) marshalImportTicket(in entity.ImportTicket) *repo.ImportTicket {
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

func (ImportTicketDB) unmarshal(in *repo.ImportTicket) *entity.ImportTicket {
	out := &entity.ImportTicket{
		ID:           in.ID,
		FromSupplier: *unmarshalSupplier(&in.FromSupplier),
		ToStorage:    *unmarshalStorage(&in.ToStorage),
		Status:       in.Status,
		SendTime:     in.SendTime,
		ReceiveTime:  in.ReceiveTime,
		Fee:          in.Fee,
	}

	out.BillImagePaths = make([]string, 0, len(in.BillImages))
	for i := range in.BillImages {
		out.BillImagePaths = append(out.BillImagePaths, in.BillImages[i].BillImagePath)
	}

	out.ProductImagePaths = make([]string, 0, len(in.ProductImages))
	for i := range in.ProductImages {
		out.ProductImagePaths = append(out.ProductImagePaths, in.ProductImages[i].ProductImagePath)
	}

	out.Details = make([]entity.ImportTicketDetails, 0, len(in.Details))
	for i := range in.Details {
		out.Details = append(out.Details, entity.ImportTicketDetails{
			Item:            *unmarshalItem(in.Details[i].Item),
			BuyQuantity:     in.Details[i].BuyQuantity,
			ReceiveQuantity: in.Details[i].ReceiveQuantity,
			BuyPrice:        in.Details[i].BuyPrice,
		})
	}

	return out
}

func (d ImportTicketDB) Create(ctx context.Context, in *entity.ImportTicket) error {
	return d.gdb.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if in.ID == 0 {
			var maxID int
			if err := tx.
				Raw(`
					select case when max(id) is null then 0 else max(id) end
					from import_ticket;
				`).Scan(&maxID).Error; err != nil {
				return err
			}

			log.Println("maxID", maxID)

			in.ID = maxID + 1
		}

		dbImportTicket := d.marshalImportTicket(*in)
		logger.Debugw("create importTicket", "importTicket db", dbImportTicket)

		if err := tx.Omit(clause.Associations).Create(dbImportTicket).Error; err != nil {
			return err
		}

		if err := tx.
			Model(&repo.ImportTicketBillImage{}).
			CreateInBatches(dbImportTicket.BillImages, len(dbImportTicket.BillImages)).Error; err != nil {
			return err
		}

		if err := tx.
			Model(&repo.ImportTicketBillImage{}).
			CreateInBatches(dbImportTicket.ProductImages, len(dbImportTicket.ProductImages)).Error; err != nil {
			return err
		}

		if err := tx.
			Model(&repo.ImportTicketBillImage{}).Omit(clause.Associations).
			CreateInBatches(dbImportTicket.Details, len(dbImportTicket.Details)).Error; err != nil {
			return err
		}

		return nil
	})
}

func (d ImportTicketDB) Get(ctx context.Context, importTicketID int) (*entity.ImportTicket, error) {
	importTicketDB := new(repo.ImportTicket)

	if err := d.gdb.WithContext(ctx).Preload(clause.Associations).Where("id = ?", importTicketID).
		Take(importTicketDB).Error; err != nil {
		if err != nil {
			return nil, err
		}
	}

	importTicket := d.unmarshal(importTicketDB)

	return importTicket, nil
}

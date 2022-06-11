package storage

import (
	"context"

	"github.com/thanhpp/scm/internal/scmsrv/domain/entity"
	"github.com/thanhpp/scm/internal/scmsrv/domain/repo"
	"gorm.io/gorm"
)

type ItemDB struct {
	gdb *gorm.DB
}

func (d DB) ItemDB() *ItemDB {
	return &ItemDB{
		gdb: d.gdb,
	}
}

func (d ItemDB) marshal(in entity.Item) *repo.Item {
	item := &repo.Item{
		SKU:       in.SKU,
		Name:      in.Name,
		Desc:      in.Desc,
		SellPrice: in.SellPrice,
		ItemType: repo.ItemType{
			ID:   in.Type.ID,
			Name: in.Type.Name,
			Desc: in.Type.Desc,
		},
	}

	// ! missing Item images

	for i := range in.Serials {
		item.Serials = append(item.Serials, repo.Serial{
			Seri:           in.Serials[i].Seri,
			ItemSKU:        in.SKU,
			StorageID:      item.Serials[i].StorageID,
			ImportTicketID: item.Serials[i].ImportTicketID,
		})
	}

	return item
}

func (d ItemDB) unmarshal(in repo.Item) *entity.Item {
	item := &entity.Item{
		SKU:       in.SKU,
		Name:      in.Name,
		Desc:      in.Desc,
		SellPrice: in.SellPrice,
	}

	// ! missing Item images

	for i := range in.Serials {
		item.Serials = append(
			item.Serials,
			&entity.Serial{
				Seri: in.Serials[i].Seri,
				Item: item,
				// ? storage
			})
	}

	return item
}

// * should have option for querying serials
func (d ItemDB) GetBySKU(ctx context.Context, sku string) (*entity.Item, error) {
	itemDB := new(repo.Item)

	if err := d.gdb.WithContext(ctx).
		Model(&repo.Item{}).Where("sku LIKE ?", sku).Take(itemDB).Error; err != nil {
		return nil, err
	}

	return d.unmarshal(*itemDB), nil
}

// ? create serial and images -> returns if error (conflict)
func (d ItemDB) Create(ctx context.Context, item entity.Item) error {
	return d.gdb.WithContext(ctx).Model(&repo.Item{}).Create(d.marshal(item)).Error
}

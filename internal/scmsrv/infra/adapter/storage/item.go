package storage

import (
	"context"

	"github.com/thanhpp/scm/internal/scmsrv/domain/entity"
	"github.com/thanhpp/scm/internal/scmsrv/domain/repo"
	"github.com/thanhpp/scm/pkg/enum"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ItemDB struct {
	gdb *gorm.DB
}

func (d DB) ItemDB() *ItemDB {
	return &ItemDB{
		gdb: d.gdb,
	}
}

func (d ItemDB) marshalItem(in entity.Item) *repo.Item {
	item := &repo.Item{
		SKU:        in.SKU,
		Name:       in.Name,
		Desc:       in.Desc,
		SellPrice:  in.SellPrice,
		ItemTypeID: in.Type.ID,
	}

	for i := range in.Images {
		item.Images = append(
			item.Images,
			repo.ItemImage{
				ItemSKU: item.SKU,
				Image:   in.Images[i],
			})
	}

	for i := range in.Serials {
		item.Serials = append(item.Serials, repo.Serial{
			Seri:           in.Serials[i].Seri,
			ItemSKU:        in.SKU,
			ImportTicketID: item.Serials[i].ImportTicketID,
		})
	}

	return item
}

func unmarshalItem(in repo.Item) *entity.Item {
	item := &entity.Item{
		SKU:       in.SKU,
		Name:      in.Name,
		Desc:      in.Desc,
		SellPrice: in.SellPrice,
		Type: entity.ItemType{
			ID:   in.ItemType.ID,
			Name: in.ItemType.Name,
			Desc: in.ItemType.Desc,
		},
	}

	for i := range in.Images {
		item.Images = append(item.Images, in.Images[i].Image)
	}

	for i := range in.Serials {
		item.Serials = append(
			item.Serials,
			&entity.Serial{
				Seri: in.Serials[i].Seri,
				Item: item,
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

	return unmarshalItem(*itemDB), nil
}

func (d ItemDB) CoundAvailabeByStorageID(ctx context.Context, storageID int) (int, error) {
	var count int64
	if err := d.gdb.WithContext(ctx).Model(&repo.Serial{}).
		Joins(`
		JOIN import_ticket ON serial.import_ticket_id = import_ticket.id
		AND import_ticket.to_storage_id = ?
		`, storageID).
		Where("(serial.status = ? OR serial.status IS NULL)", enum.SerialStatusNew).
		Group("item_sku").
		Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}

// ? create serial and images -> returns if error (conflict)
func (d ItemDB) CreateItem(ctx context.Context, item entity.Item) error {
	return d.gdb.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		itemDB := d.marshalItem(item)

		if err := tx.Model(&repo.Item{}).Omit(clause.Associations).
			Create(itemDB).Error; err != nil {
			return err
		}

		if err := tx.Model(&repo.ItemImage{}).CreateInBatches(
			itemDB.Images, len(itemDB.Images),
		).Error; err != nil {
			return err
		}

		return nil
	})
}

func (d ItemDB) GetList(ctx context.Context, filer repo.ItemFilter) ([]*entity.Item, error) {
	itemsDB := make([]*repo.Item, filer.Limit)

	if err := d.gdb.
		WithContext(ctx).Preload(clause.Associations).
		Model(&repo.Item{}).Limit(filer.Limit).Offset(filer.Offset).Find(&itemsDB).Error; err != nil {
		return nil, err
	}

	items := make([]*entity.Item, len(itemsDB))
	for i := range items {
		items[i] = unmarshalItem(*itemsDB[i])
	}

	return items, nil
}

// ----------------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- ITEM TYPE ---------------------------------------------------------------

func (d ItemDB) unmarshalItemType(in *repo.ItemType) *entity.ItemType {
	return &entity.ItemType{
		ID:   in.ID,
		Name: in.Name,
		Desc: in.Desc,
	}
}

func (ItemDB) marshalItemType(in *entity.ItemType) *repo.ItemType {
	return &repo.ItemType{
		Name: in.Name,
		Desc: in.Desc,
	}
}

func (d ItemDB) GetItemType(ctx context.Context, itemTypeID int) (*entity.ItemType, error) {
	itemDB := new(repo.ItemType)

	if err := d.gdb.WithContext(ctx).Where("id = ?", itemTypeID).Take(itemDB).Error; err != nil {
		return nil, err
	}

	return d.unmarshalItemType(itemDB), nil
}

func (d ItemDB) CreateItemType(ctx context.Context, itemType *entity.ItemType) error {
	itemTypeDB := d.marshalItemType(itemType)

	if err := d.gdb.WithContext(ctx).Create(itemTypeDB).Error; err != nil {
		return err
	}

	itemType.ID = itemTypeDB.ID

	return nil
}

func (d ItemDB) GetAllItemType(ctx context.Context) ([]*entity.ItemType, error) {
	var itemTypeDBs []*repo.ItemType

	if err := d.gdb.WithContext(ctx).Model(&repo.ItemType{}).Find(&itemTypeDBs).Error; err != nil {
		return nil, err
	}

	itemTypes := make([]*entity.ItemType, len(itemTypeDBs))

	for i := range itemTypeDBs {
		itemTypes[i] = d.unmarshalItemType(itemTypeDBs[i])
	}

	return itemTypes, nil
}

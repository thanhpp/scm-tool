package storage

import (
	"context"

	"github.com/thanhpp/scm/internal/scmsrv/domain/entity"
	"github.com/thanhpp/scm/internal/scmsrv/domain/repo"
	"github.com/thanhpp/scm/pkg/logger"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SerialDB struct {
	gdb *gorm.DB
}

func (d DB) SerialDB() repo.SerialRepo {
	return &SerialDB{
		gdb: d.gdb,
	}
}

func (d SerialDB) Count(ctx context.Context, importTicketID int, itemSKU string) (int, error) {
	var seriCount int64

	if err := d.gdb.
		WithContext(ctx).Model(&repo.Serial{}).
		Where("import_ticket_id = ? AND item_sku = ?", importTicketID, itemSKU).
		Count(&seriCount).Error; err != nil {
		return 0, err
	}

	return int(seriCount), nil
}

func (d SerialDB) Get(ctx context.Context, seri string) (*entity.Serial, error) {
	serialDB := new(repo.Serial)
	if err := d.gdb.
		WithContext(ctx).
		Model(serialDB).
		Where("seri LIKE ?", seri).
		First(serialDB).Error; err != nil {
		return nil, err
	}

	itemDB := new(repo.Item)
	if err := d.gdb.
		WithContext(ctx).
		Model(itemDB).
		Preload(clause.Associations).
		Where("sku LIKE ?", serialDB.ItemSKU).
		First(itemDB).Error; err != nil {
		return nil, err
	}

	importTicketDB := new(repo.ImportTicket)
	if err := d.gdb.
		WithContext(ctx).
		Model(importTicketDB).
		Preload(clause.Associations).
		Where("id = ?", serialDB.ImportTicketID).
		First(importTicketDB).Error; err != nil {
		return nil, err
	}

	serialDB.Item = *itemDB
	serialDB.ImportTicket = *importTicketDB

	serial := unmarshalSerial(*serialDB)

	return &serial, nil
}

func (d SerialDB) GetSeriWithEmptyTokenID(ctx context.Context) ([]*entity.Serial, error) {
	var series []string

	if err := d.gdb.WithContext(ctx).
		Model(&repo.Serial{}).Select("seri").Where("token_id = ?", 0).Find(&series).Error; err != nil {
		return nil, err
	}

	result := make([]*entity.Serial, len(series))
	for i := range series {
		serial, err := d.Get(ctx, series[i])
		if err != nil {
			return nil, err
		}

		result[i] = serial
	}

	return result, nil
}

func (d SerialDB) CreateBatch(ctx context.Context, serials []*entity.Serial) error {
	var serialsDB = make([]*repo.Serial, len(serials))

	for i := range serials {
		newSerial := marshalSerial(*serials[i])
		serialsDB[i] = &newSerial
		if i == 0 {
			logger.Debugw("marshal serial", "entity", serials[i], "repo", newSerial)
		}
	}

	if err := d.gdb.
		WithContext(ctx).
		Model(&repo.Serial{}).
		CreateInBatches(serialsDB, len(serialsDB)).Error; err != nil {
		return err
	}

	return nil
}

func (s SerialDB) UpdateSerial(ctx context.Context, seri string, fn repo.UpdateSerialFn) error {
	return s.gdb.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		serial, err := s.Get(ctx, seri)
		if err != nil {
			return err
		}

		newSerial, err := fn(ctx, serial)
		if err != nil {
			return err
		}

		newSerialDB := marshalSerial(*newSerial)

		if err := tx.
			WithContext(ctx).
			Model(&repo.Serial{}).
			Where("seri = ?", seri).Updates(newSerialDB).Error; err != nil {
			return err
		}

		return nil
	})
}

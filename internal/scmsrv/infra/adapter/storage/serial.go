package storage

import (
	"context"

	"github.com/thanhpp/scm/internal/scmsrv/domain/entity"
	"github.com/thanhpp/scm/internal/scmsrv/domain/repo"
	"gorm.io/gorm"
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
		WithContext(ctx).Model(&SerialDB{}).
		Where("import_ticket_id = ? AND item_sku = ?", importTicketID, itemSKU).
		Count(&seriCount).Error; err != nil {
		return 0, err
	}

	return int(seriCount), nil
}

func (d SerialDB) CreateBatch(ctx context.Context, serials []*entity.Serial) error {
	var serialsDB = make([]*repo.Serial, len(serials))

	for i := range serials {
		newSerial := marshalSerial(*serials[i])
		serialsDB[i] = &newSerial
	}

	if err := d.gdb.
		WithContext(ctx).
		Model(&repo.Serial{}).
		CreateInBatches(serialsDB, len(serialsDB)).Error; err != nil {
		return err
	}

	return nil
}

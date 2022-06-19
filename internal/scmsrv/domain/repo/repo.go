package repo

import (
	"context"

	"github.com/thanhpp/scm/internal/scmsrv/domain/entity"
)

type ItemUpdateFn func(ctx context.Context, item *Item) (*Item, error)

type ItemRepo interface {
	// Read
	GetBySKU(ctx context.Context, sku string) (*entity.Item, error)
	GetItemType(ctx context.Context, itemTypeID int) (*entity.ItemType, error)

	// Write
	CreateItem(ctx context.Context, item entity.Item) error
	// Update(ctx context.Context, sku string, fn ItemUpdateFn) error
	// Delete(ctx context.Context, sku string) error

	CreateItemType(ctx context.Context, itemType *entity.ItemType) error
}

type ImportTicketRepo interface {
	// Write
	Create(ctx context.Context, importTicket *entity.ImportTicket) error
}

type SupplierRepo interface {
	// Read
	Get(ctx context.Context, id int) (*entity.Supplier, error)

	// Write
	Create(ctx context.Context, supplier *entity.Supplier) error
}

type StorageRepo interface {
	// Read
	Get(ctx context.Context, id int) (*entity.Storage, error)

	// Write
	Create(ctx context.Context, storage *entity.Storage) error
}

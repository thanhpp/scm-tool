package repo

import (
	"context"

	"github.com/thanhpp/scm/internal/scmsrv/domain/entity"
)

type ItemUpdateFn func(ctx context.Context, item *Item) (*Item, error)

type ItemRepo interface {
	// Read
	GetBySKU(ctx context.Context, sku string) (*Item, error)

	// Write
	Create(ctx context.Context, item *Item) error
	Update(ctx context.Context, sku string, fn ItemUpdateFn) error
	Delete(ctx context.Context, sku string) error
}

type ImportTicketRepo interface {
	// Write
	Create(ctx context.Context, importTicket *entity.ImportTicket) error
}

type SupplierRepo interface {
	Get(ctx context.Context, id int) (*entity.Supplier, error)
}

type StorageRepo interface {
	Get(ctx context.Context, id int) (*entity.Storage, error)
}

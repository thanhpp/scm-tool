package repo

import "context"

type ItemUpdateFn func(ctx context.Context, item *Item) (*Item, error)

type ItemRepo interface {
	// Read
	GetBySKU(ctx context.Context, sku string) (*Item, error)

	// Write
	Create(ctx context.Context, item *Item) error
	Update(ctx context.Context, sku string, fn ItemUpdateFn) error
	Delete(ctx context.Context, sku string) error
}

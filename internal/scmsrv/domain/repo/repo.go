package repo

import (
	"context"

	"github.com/thanhpp/scm/internal/scmsrv/domain/entity"
)

type ItemUpdateFn func(ctx context.Context, item entity.Item) (entity.Item, error)

type ItemTypeUpdateFn func(ctx context.Context, itemType *ItemType) (*ItemType, error)

type ItemFilter struct {
	Limit  int
	Offset int
}

type ItemRepo interface {
	// Read
	GetBySKU(ctx context.Context, sku string) (*entity.Item, error)
	GetList(ctx context.Context, filer ItemFilter) ([]*entity.Item, error)
	GetItemType(ctx context.Context, itemTypeID int) (*entity.ItemType, error)
	GetAllItemType(ctx context.Context) ([]*entity.ItemType, error)
	CoundAvailabeByStorageID(ctx context.Context, storageID int) (int, error)

	// Write
	CreateItem(ctx context.Context, item entity.Item) error
	UpdateItem(ctx context.Context, sku string, fn ItemUpdateFn) error
	// Delete(ctx context.Context, sku string) error

	CreateItemType(ctx context.Context, itemType *entity.ItemType) error
	UpdateItemType(ctx context.Context, itemTypeID int, fn ItemTypeUpdateFn) error
}

type ImportTicketRepoFiler struct {
}

type ImportTicketRepo interface {
	// Write
	Create(ctx context.Context, importTicket *entity.ImportTicket) error

	// Read
	Get(ctx context.Context, importTicketID int) (*entity.ImportTicket, error)
	GetGeneralInfoList(ctx context.Context, offset, limit int) ([]*entity.ImportTicket, error)
}

type SupplierFiler struct {
	Limit  int
	Offset int
}

type SupplierRepo interface {
	// Read
	Get(ctx context.Context, id int) (*entity.Supplier, error)
	GetList(ctx context.Context, filter SupplierFiler) ([]*entity.Supplier, error)

	// Write
	Create(ctx context.Context, supplier *entity.Supplier) error

	// Update
	Update(context.Context, *entity.Supplier) error
}

type StorageFiler struct {
	Limit  int
	Offset int
}

type StorageUpdateFn func(*entity.Storage) (*entity.Storage, error)

type StorageRepo interface {
	// Read
	Get(ctx context.Context, id int) (*entity.Storage, error)
	GetList(ctx context.Context, filter StorageFiler) ([]*entity.Storage, error)

	// Write
	Create(ctx context.Context, storage *entity.Storage) error
	Update(ctx context.Context, storageID int, fn StorageUpdateFn) error
}

type UpdateSerialFn func(context.Context, *entity.Serial) (*entity.Serial, error)

type SerialRepo interface {
	// Read
	Count(ctx context.Context, importTicketID int, itemSKU string) (int, error)
	Get(ctx context.Context, seri string) (*entity.Serial, error)
	GetSeriWithEmptyTokenID(ctx context.Context) ([]*entity.Serial, error)
	GetSerialsByImportTicketID(ctx context.Context, importTicketID int) ([]*entity.Serial, error)

	// Write
	CreateBatch(ctx context.Context, serials []*entity.Serial) error
	UpdateSerial(ctx context.Context, seri string, fn UpdateSerialFn) error
}

type UpdateUserFn func(context.Context, entity.User) (entity.User, error)

type GetUsersFilter struct {
	Limit  int
	Offset int
}

type UserRepo interface {
	// Read
	GetByUsername(ctx context.Context, username string) (*entity.User, error)
	GetUsers(ctx context.Context, filer GetUsersFilter) ([]*entity.User, error)

	// Create
	NewUser(ctx context.Context, user *entity.User) error
	UpdateUserByID(ctx context.Context, id int, fn UpdateUserFn) error
}

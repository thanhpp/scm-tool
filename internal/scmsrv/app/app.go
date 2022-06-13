package app

import (
	"github.com/thanhpp/scm/internal/scmsrv/domain/entity"
	"github.com/thanhpp/scm/internal/scmsrv/domain/repo"
	"github.com/thanhpp/scm/pkg/fileutil"
)

type App struct {
	ImportTicketHandler ImportTicketHandler
	SupplierHandler     SupplierHanlder
	StorageHandler      StorageHandler
}

func New(
	fac entity.Factory,
	itemRepo repo.ItemRepo, supplierRepo repo.SupplierRepo, storageRepo repo.StorageRepo,
	importTicketRepo repo.ImportTicketRepo,
	fileUtil fileutil.FileUtil,
) App {
	return App{
		ImportTicketHandler: ImportTicketHandler{
			itemRepo:         itemRepo,
			supplierRepo:     supplierRepo,
			storageRepo:      storageRepo,
			importTicketRepo: importTicketRepo,
			fac:              fac,
			fileUtil:         fileUtil,
		},
		SupplierHandler: SupplierHanlder{
			fac:          fac,
			supplierRepo: supplierRepo,
		},
		StorageHandler: StorageHandler{
			fac:         fac,
			storageRepo: storageRepo,
		},
	}
}

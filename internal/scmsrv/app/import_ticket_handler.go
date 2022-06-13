package app

import (
	"context"
	"time"

	"github.com/thanhpp/scm/internal/scmsrv/domain/entity"
	"github.com/thanhpp/scm/internal/scmsrv/domain/repo"
	"github.com/thanhpp/scm/pkg/fileutil"
)

type ImportTicketHandler struct {
	itemRepo         repo.ItemRepo
	supplierRepo     repo.SupplierRepo
	storageRepo      repo.StorageRepo
	importTicketRepo repo.ImportTicketRepo
	fac              entity.Factory
	fileUtil         fileutil.FileUtil
}

func (h ImportTicketHandler) Create(
	ctx context.Context, supplierID, storageID int,
	sendTime, receiveTime time.Time, fee float64,
	details []entity.ImportTicketDetails,
	// * images
) (*entity.ImportTicket, error) {
	supplier, err := h.supplierRepo.Get(ctx, supplierID)
	if err != nil {
		return nil, err
	}

	storage, err := h.storageRepo.Get(ctx, storageID)
	if err != nil {
		return nil, err
	}

	newImportTicket, err := h.fac.NewImportTicket(
		*supplier, *storage, sendTime, fee,
		details, nil, nil, // ! missing images
	)
	if err != nil {
		return nil, err
	}

	if err := h.importTicketRepo.Create(ctx, newImportTicket); err != nil {
		return nil, err
	}

	return newImportTicket, nil
}

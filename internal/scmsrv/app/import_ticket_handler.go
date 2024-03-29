package app

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/pkg/errors"
	"github.com/thanhpp/scm/internal/scmsrv/domain/entity"
	"github.com/thanhpp/scm/internal/scmsrv/domain/repo"
	"github.com/thanhpp/scm/internal/scmsrv/infra/adapter/nftsvclient"
	"github.com/thanhpp/scm/pkg/constx"
	"github.com/thanhpp/scm/pkg/fileutil"
)

type ImportTicketHandler struct {
	itemRepo         repo.ItemRepo
	supplierRepo     repo.SupplierRepo
	storageRepo      repo.StorageRepo
	importTicketRepo repo.ImportTicketRepo
	serialRepo       repo.SerialRepo
	fac              entity.Factory
	fileUtil         fileutil.FileUtil
	nftServiceClient *nftsvclient.NFTServiceClient
}

func (h ImportTicketHandler) Create(
	ctx context.Context, supplierID, storageID int,
	sendTime, receiveTime time.Time, fee float64,
	details []entity.ImportTicketDetails,
	billImages, productImages []*multipart.FileHeader,
) (*entity.ImportTicket, error) {
	if err := h.imagesTypeCheck(billImages); err != nil {
		return nil, err
	}

	if err := h.imagesTypeCheck(productImages); err != nil {
		return nil, err
	}

	billImagesPath, err := h.
		fileUtil.
		SaveFilesFromMultipart(constx.SaveFilePaths, "import_ticket-bill_images", billImages)
	if err != nil {
		return nil, err
	}

	productsImagesPath, err := h.
		fileUtil.
		SaveFilesFromMultipart(constx.SaveFilePaths, "import_ticket-product_images", productImages)
	if err != nil {
		return nil, err
	}

	supplier, err := h.supplierRepo.Get(ctx, supplierID)
	if err != nil {
		return nil, fmt.Errorf("get supplier error: %w", err)
	}

	storage, err := h.storageRepo.Get(ctx, storageID)
	if err != nil {
		return nil, fmt.Errorf("get storage error: %w", err)
	}

	newImportTicket, err := h.fac.NewImportTicket(
		*supplier, *storage, sendTime, receiveTime, fee,
		details, billImagesPath, productsImagesPath,
	)
	if err != nil {
		return nil, err
	}

	log.Printf("newImportTicket entity: %+v", newImportTicket)

	if err := h.importTicketRepo.Create(ctx, newImportTicket); err != nil {
		return nil, fmt.Errorf("create import ticket error: %w", err)
	}

	return newImportTicket, nil
}

// NewImportDetails too many queries. But who knows :D
func (h ImportTicketHandler) CreateImportDetails(
	ctx context.Context, sku string, buyQty, receiveQty int, butPrice float64,
) (*entity.ImportTicketDetails, error) {
	item, err := h.itemRepo.GetBySKU(ctx, sku)
	if err != nil {
		return nil, err
	}

	detail, err := h.fac.NewImportTicketDetails(*item, buyQty, receiveQty, butPrice)
	if err != nil {
		return nil, err
	}

	return detail, nil
}

func (h ImportTicketHandler) GetSerialsByImportTicketID(
	ctx context.Context, importTicketID int,
) (map[string][]*entity.Serial, error) {
	serials, err := h.serialRepo.GetSerialsByImportTicketID(ctx, importTicketID)
	if err != nil {
		return nil, fmt.Errorf("get serial by import ticket id error: %w", err)
	}

	m := make(map[string][]*entity.Serial)
	for i := range serials {
		m[serials[i].Item.SKU] = append(m[serials[i].Item.SKU], serials[i])
	}

	return m, nil
}

func (h ImportTicketHandler) GetSerialInfo(
	ctx context.Context, seri string,
) (*entity.Serial, *nftsvclient.NFTInfo, error) {
	serial, err := h.serialRepo.Get(ctx, seri)
	if err != nil {
		return nil, nil, err
	}

	nftInfo, err := h.nftServiceClient.GetNFTInfoBySeri(ctx, seri)
	if err != nil {
		return serial, new(nftsvclient.NFTInfo), nil
	}

	return serial, nftInfo, nil
}

func (h ImportTicketHandler) GenSerials(
	ctx context.Context, importTicketID int,
) ([]*entity.Serial, error) {
	// check import ticket
	importTicket, err := h.importTicketRepo.Get(ctx, importTicketID)
	if err != nil {
		return nil, err
	}

	// for each item details -
	var serials []*entity.Serial
	for i := range importTicket.Details {
		//  check if serials exist
		n, err := h.serialRepo.Count(ctx, importTicketID, importTicket.Details[i].Item.SKU)
		if err != nil {
			return nil, err
		}
		if n != 0 {
			return nil, errors.New("serial exists")
		}

		// create serials
		detailSerials, err := h.fac.
			NewSerials(importTicket, &importTicket.Details[i].Item, importTicket.Details[i].BuyQuantity)
		if err != nil {
			return nil, err
		}

		serials = append(serials, detailSerials...)
	}

	// save serials
	if err := h.serialRepo.CreateBatch(ctx, serials); err != nil {
		return nil, err
	}

	return serials, nil
}

var (
	imageTypes = []string{".jpg", ".jpeg", ".png"}
)

func (h ImportTicketHandler) imagesTypeCheck(files []*multipart.FileHeader) error {
	for i := range files {
		if !isInStringSlices(filepath.Ext(files[i].Filename), imageTypes) {
			return errors.New("file not image")
		}
	}

	return nil
}

func isInStringSlices(val string, sl []string) bool {
	for i := range sl {
		if sl[i] == val {
			return true
		}
	}

	return false
}

func (h ImportTicketHandler) GetImportTicketByID(
	ctx context.Context, id int,
) (*entity.ImportTicket, error) {
	importTicket, err := h.importTicketRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return importTicket, nil
}

func (h ImportTicketHandler) GetListImportTicket(
	ctx context.Context, offset, limit int,
) ([]*entity.ImportTicket, error) {
	importTickets, err := h.importTicketRepo.GetGeneralInfoList(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	return importTickets, nil
}

package app

import (
	"context"
	"mime/multipart"

	"github.com/thanhpp/scm/internal/scmsrv/domain/entity"
	"github.com/thanhpp/scm/internal/scmsrv/domain/repo"
	"github.com/thanhpp/scm/pkg/fileutil"
)

type ItemHandler struct {
	fac      entity.Factory
	itemRepo repo.ItemRepo
	fileUtil fileutil.FileUtil
}

func (h ItemHandler) Create(
	ctx context.Context,
	sku, name, desc string, itemTypeID int,
	images []*multipart.FileHeader,
) (*entity.Item, error) {
	itemType, err := h.itemRepo.GetItemType(ctx, itemTypeID)
	if err != nil {
		return nil, err
	}

	// ! remove images if error
	imagePaths, err := h.fileUtil.SaveFilesFromMultipart("", "item-images-"+sku, images)
	if err != nil {
		return nil, err
	}

	newItem, err := h.fac.NewItem(sku, name, desc, *itemType, imagePaths)
	if err != nil {
		return nil, err
	}

	return newItem, nil
}

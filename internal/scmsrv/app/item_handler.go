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

func (h ItemHandler) CreateItem(
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

func (h ItemHandler) GetList(ctx context.Context, page, size int) ([]*entity.Item, error) {
	offset, limit := genOffsetLimit(page, size)

	return h.itemRepo.GetList(ctx, repo.ItemFilter{
		Offset: offset,
		Limit:  limit,
	})
}

func (h ItemHandler) CreateItemType(
	ctx context.Context,
	name, desc string,
) (*entity.ItemType, error) {
	newItemType, err := h.fac.NewItemType(name, desc)
	if err != nil {
		return nil, err
	}

	if err := h.itemRepo.CreateItemType(ctx, newItemType); err != nil {
		return nil, err
	}

	return newItemType, nil
}

func (h ItemHandler) GetAllItemType(ctx context.Context) ([]*entity.ItemType, error) {
	return h.itemRepo.GetAllItemType(ctx)
}

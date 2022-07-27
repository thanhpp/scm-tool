package app

import (
	"context"
	"mime/multipart"

	"github.com/pkg/errors"
	"github.com/thanhpp/scm/internal/scmsrv/domain/entity"
	"github.com/thanhpp/scm/internal/scmsrv/domain/repo"
	"github.com/thanhpp/scm/pkg/constx"
	"github.com/thanhpp/scm/pkg/fileutil"
	"github.com/thanhpp/scm/pkg/logger"
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
		return nil, errors.WithMessage(err, "get item type")
	}

	// ! remove images if error
	imagePaths, err := h.fileUtil.SaveFilesFromMultipart(constx.SaveFilePaths, "item-images-"+sku, images)
	if err != nil {
		return nil, errors.WithMessage(err, "save images")
	}

	newItem, err := h.fac.NewItem(sku, name, desc, *itemType, imagePaths)
	if err != nil {
		return nil, err
	}

	if err := h.itemRepo.CreateItem(ctx, *newItem); err != nil {
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

func (h ItemHandler) UpdateItem(
	ctx context.Context,
	sku, name, desc string, itemTypeID int, sellPrice float64,
	newImages []*multipart.FileHeader, deleteImages []string,
) error {
	itemType, err := h.itemRepo.GetItemType(ctx, itemTypeID)
	if err != nil {
		return err
	}

	return h.itemRepo.UpdateItem(ctx, sku,
		func(ctx context.Context, item entity.Item) (entity.Item, error) {
			logger.Debugw("update item", "current images", item.Images)
			for i := range deleteImages {
				if !item.DeleteImages(constx.SaveFilePaths + "/" + deleteImages[i]) {
					return entity.Item{}, errors.New("delete not exist images: " + deleteImages[i])
				}
			}
			logger.Debugw("update item", "after delete images", item.Images)

			newImagePaths, err := h.fileUtil.
				SaveFilesFromMultipart(constx.SaveFilePaths, "item-images-"+sku, newImages)
			if err != nil {
				return entity.Item{}, err
			}

			item.Images = append(item.Images, newImagePaths...)

			if err := item.SetName(name); err != nil {
				return entity.Item{}, err
			}

			item.Desc = desc
			item.Type = *itemType
			item.SellPrice = sellPrice

			return item, nil
		},
	)
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

func (h ItemHandler) UpdateItemType(ctx context.Context, id int, name, desc string) error {
	if len(name) == 0 {
		return errors.New("update item type: empty name")
	}

	return h.itemRepo.UpdateItemType(ctx, id,
		func(ctx context.Context, itemType *repo.ItemType) (*repo.ItemType, error) {
			if itemType == nil {
				return nil, errors.New("update nil item type")
			}

			itemType.Name = name
			itemType.Desc = desc

			return itemType, nil
		},
	)
}

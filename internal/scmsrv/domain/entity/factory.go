package entity

import (
	"time"

	"github.com/pkg/errors"
)

type Factory interface {
	NewImportTicket(
		fromSupplier Supplier, toStorage Storage, sendTime time.Time, fee float64, details []ImportTicketDetails,
		billImagePaths []string, productImagePaths []string,
	) (*ImportTicket, error)
	NewImportTicketDetails(item Item, buyQuantity, receiveQuantity int, buyPrice float64) (*ImportTicketDetails, error)
	NewSupplier(name, phone, email string) (*Supplier, error)
	NewStorage(name, desc, location string) (*Storage, error)
	NewItem(sku, name, desc string, itemType ItemType, imagePaths []string) (*Item, error)
}

type factoryImpl struct{}

func NewFactory() Factory {
	return factoryImpl{}
}

func (f factoryImpl) NewImportTicket(
	fromSupplier Supplier, toStorage Storage, sendTime time.Time, fee float64, details []ImportTicketDetails,
	billImagePaths []string, productImagePaths []string,
) (*ImportTicket, error) {
	if sendTime.IsZero() {
		return nil, errors.New("create import ticket: empty send time")
	}

	if len(details) == 0 {
		return nil, errors.New("create import ticket: empty details")
	}

	skuMap := make(map[string]struct{}, len(details))
	for i := range details {
		if _, ok := skuMap[details[i].Item.SKU]; ok {
			return nil, errors.New("create import ticket: duplicate sku" + details[i].Item.SKU)
		}
		skuMap[details[i].Item.SKU] = struct{}{}
	}

	importTicket := &ImportTicket{
		FromSupplier: fromSupplier,
		ToStorage:    toStorage,
		Status:       ImportTicketStatusNew,
		SendTime:     sendTime,
		ReceiveTime:  time.Time{},

		BillImagePaths:    billImagePaths,
		ProductImagePaths: productImagePaths,

		Details: details,
	}

	return importTicket, nil
}

func (factoryImpl) NewImportTicketDetails(
	item Item, buyQuantity, receiveQuantity int, buyPrice float64,
) (*ImportTicketDetails, error) {
	if item.IsEmpty() {
		return nil, errors.New("create import ticket detail: empty item")
	}

	if buyQuantity == 0 {
		return nil, errors.New("create import ticket detail: zero buy quantity")
	}

	return &ImportTicketDetails{
		Item:            item,
		BuyQuantity:     buyQuantity,
		ReceiveQuantity: receiveQuantity,
		BuyPrice:        buyPrice,
	}, nil
}

func (f factoryImpl) NewSupplier(name, phone, email string) (*Supplier, error) {
	if len(name) == 0 {
		return nil, errors.New("create supplier: empty name")
	}

	if len(phone)+len(email) == 0 {
		return nil, errors.New("create supplier: empty contact")
	}

	return &Supplier{
		Name:  name,
		Phone: phone,
		Email: email,
	}, nil
}

func (f factoryImpl) NewStorage(name, desc, location string) (*Storage, error) {
	if len(name) == 0 {
		return nil, errors.New("create storage: empty name")
	}

	if len(location) == 0 {
		return nil, errors.New("create storage: empty location")
	}

	return &Storage{
		Name:     name,
		Desc:     desc,
		Location: location,
	}, nil
}

func (factoryImpl) NewItem(sku, name, desc string, itemType ItemType, imagePaths []string) (*Item, error) {
	if len(sku) == 0 {
		return nil, errors.New("create item: empty sku")
	}

	newItem := &Item{
		SKU:  sku,
		Name: name,
		Desc: desc,
		Type: itemType,
	}
	newItem.Images = imagePaths

	return newItem, nil
}

package entity

import (
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/xid"
	"github.com/thanhpp/scm/pkg/enum"
)

type Factory interface {
	NewImportTicket(
		fromSupplier Supplier, toStorage Storage, sendTime, receiveTime time.Time, fee float64, details []ImportTicketDetails,
		billImagePaths []string, productImagePaths []string,
	) (*ImportTicket, error)
	NewImportTicketDetails(item Item, buyQuantity, receiveQuantity int, buyPrice float64) (*ImportTicketDetails, error)
	NewSupplier(name, email, phone string) (*Supplier, error)
	NewStorage(name, desc, location string) (*Storage, error)
	NewItem(sku, name, desc string, itemType ItemType, imagePaths []string) (*Item, error)
	NewItemType(name, desc string) (*ItemType, error)
	NewSerials(importTicket *ImportTicket, item *Item, num int) ([]*Serial, error)
	NewUser(name, username, password string) (*User, error)
}

type factoryImpl struct{}

func NewFactory() Factory {
	return factoryImpl{}
}

func (f factoryImpl) NewImportTicket(
	fromSupplier Supplier, toStorage Storage, sendTime, receiveTime time.Time, fee float64, details []ImportTicketDetails,
	billImagePaths []string, productImagePaths []string,
) (*ImportTicket, error) {
	// if sendTime.IsZero() {
	// 	return nil, errors.New("create import ticket: empty send time")
	// }

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

	if receiveTime.IsZero() {
		receiveTime = sendTime
	}

	importTicket := &ImportTicket{
		FromSupplier: fromSupplier,
		ToStorage:    toStorage,
		Status:       enum.ImportTicketStatusNew,
		SendTime:     sendTime,
		ReceiveTime:  receiveTime,
		Fee:          fee,

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

func (f factoryImpl) NewSupplier(name, email, phone string) (*Supplier, error) {
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

	sku = strings.ReplaceAll(sku, " ", "")

	newItem := &Item{
		SKU:  sku,
		Name: name,
		Desc: desc,
		Type: itemType,
	}
	newItem.Images = imagePaths

	return newItem, nil
}

func (factoryImpl) NewItemType(name, desc string) (*ItemType, error) {
	if len(name) == 0 {
		return nil, errors.New("create item type: empty name")
	}

	return &ItemType{
		Name: name,
		Desc: desc,
	}, nil
}

func (factoryImpl) NewSerials(importTicket *ImportTicket, item *Item, num int) ([]*Serial, error) {
	if importTicket == nil {
		return nil, errors.New("create serials: empty import ticket")
	}

	if item == nil {
		return nil, errors.New("create serials: empty item")
	}

	if num == 0 {
		return nil, errors.New("create serials: zero num")
	}

	serials := make([]*Serial, num)
	for i := range serials {
		serials[i] = &Serial{
			Seri:         stringToInt(xid.New().String()),
			Status:       enum.SerialStatusNew,
			ImportTicket: importTicket,
			Item:         item,
		}
	}

	return serials, nil
}

func stringToInt(in string) string {
	b := new(strings.Builder)
	b.Grow(len(in))

	for i := range in {
		b.WriteString(fmt.Sprintf("%d", in[i]))
	}

	return b.String()
}

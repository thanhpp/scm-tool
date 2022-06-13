package entity

import "time"

type Factory interface {
	NewImportTicket(
		fromSupplier Supplier, toStorage Storage, sendTime time.Time, fee float64, details []ImportTicketDetails,
		billImagePaths []string, productImagePaths []string,
	) (*ImportTicket, error)
	NewSupplier(name, phone, email string) *Supplier
}

type factoryImpl struct{}

func NewFactory() Factory {
	return factoryImpl{}
}

func (f factoryImpl) NewImportTicket(
	fromSupplier Supplier, toStorage Storage, sendTime time.Time, fee float64, details []ImportTicketDetails,
	billImagePaths []string, productImagePaths []string,
) (*ImportTicket, error) {
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

func (f factoryImpl) NewSupplier(name, phone, email string) *Supplier {
	return &Supplier{
		Name:  name,
		Phone: phone,
		Email: email,
	}
}

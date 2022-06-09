package entity

import (
	"time"

	"github.com/thanhpp/scm/pkg/enum"
)

type ImportTicketStatus int

const (
	ImportTicketStatusNew = 1 + iota
)

type ImportTicket struct {
	ID           int
	FromSupplier Supplier
	ToStorage    Storage
	Status       enum.ImportTicketStatus
	SendTime     time.Time
	ReceiveTime  time.Time
	Fee          float64

	BillImagePaths    []string
	ProductImagePaths []string

	Details []ImportTicketDetails
}

type ImportTicketDetails struct {
	Item            Item
	BuyQuantity     int
	ReceiveQuantity int
	BuyPrice        float64
}

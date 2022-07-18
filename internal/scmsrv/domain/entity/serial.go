package entity

import "github.com/thanhpp/scm/pkg/enum"

type Serial struct {
	Seri         string
	Status       enum.SerialStatus
	TokenID      int
	Item         *Item
	ImportTicket *ImportTicket
}

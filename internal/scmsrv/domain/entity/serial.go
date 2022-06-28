package entity

import "github.com/thanhpp/scm/pkg/enum"

type Serial struct {
	Seri         string
	Status       enum.SerialStatus
	Item         *Item
	ImportTicket *ImportTicket
}

package entity

import "github.com/google/uuid"

type Storage struct {
	ID       uuid.UUID
	Name     string
	Desc     string
	Location string

	Serials []Serial
}

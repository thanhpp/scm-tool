package entity

type Storage struct {
	ID       int
	Name     string
	Desc     string
	Location string

	Serials []Serial
}

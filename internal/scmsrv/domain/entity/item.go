package entity

type Item struct {
	SKU       string
	Name      string
	Desc      string
	Images    []string
	SellPrice float64
	Type      ItemType

	// default supplier
	// default storage

	Serials []*Serial
}

func (i Item) IsEmpty() bool {
	return len(i.SKU) == 0
}

type ItemType struct {
	ID   int
	Name string
	Desc string
}

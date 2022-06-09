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

type ItemType struct {
	ID   int
	Name string
	Desc string
}

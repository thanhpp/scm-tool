package entity

import "errors"

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

func (it *Item) DeleteImages(image string) bool {
	for i := range it.Images {
		if it.Images[i] == image {
			it.Images = append(it.Images[:i], it.Images[i+1:]...)
			return true
		}
	}

	return false
}

func (it *Item) SetName(name string) error {
	if len(name) == 0 {
		return errors.New("set empty item name")
	}

	it.Name = name

	return nil
}

func (i Item) IsEmpty() bool {
	return len(i.SKU) == 0
}

type ItemType struct {
	ID   int
	Name string
	Desc string
}

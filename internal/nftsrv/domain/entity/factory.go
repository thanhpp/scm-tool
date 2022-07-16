package entity

import "errors"

type Factory struct{}

func (f Factory) NewSeriNFT(seri string, metadata map[string]interface{}) (*SerialNFT, error) {
	if metadata == nil {
		return nil, errors.New("new seri nft: empty metadata")
	}

	if _, ok := metadata["name"]; ok {
		return nil, errors.New("new seri nft: duplicate 'name' entry")
	}

	metadata["name"] = seri

	return nil, nil
}

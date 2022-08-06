package sharedto

import (
	"errors"
	"strconv"

	"github.com/thanhpp/scm/internal/scmsrv/domain/entity"
)

const (
	MintNFTRequestQueue = "mintnft"
	SeriNFTInfoQueue    = "serinftinfo"
)

type MetadataCreateSeriNFT struct {
	ItemName       string `json:"item_name"`
	ItemDesc       string `json:"item_desc"`
	SupplierName   string `json:"supplier_name"`
	ImportTicketID string `json:"import_ticket_id"`
}

type ReqCreateSeriNFT struct {
	Seri     string                `json:"seri"`
	Metadata MetadataCreateSeriNFT `json:"metadata"`
}

func (req *ReqCreateSeriNFT) SetData(serial *entity.Serial) error {
	if serial == nil || serial.Item == nil || serial.ImportTicket == nil {
		return errors.New("set data create serinft is nil")
	}

	req.Seri = serial.Seri

	req.Metadata = MetadataCreateSeriNFT{
		ItemName:       serial.Item.Name,
		ItemDesc:       serial.Item.Desc,
		SupplierName:   serial.ImportTicket.FromSupplier.Name,
		ImportTicketID: strconv.Itoa(serial.ImportTicket.ID),
	}

	return nil
}

type SeriNFTInfo struct {
	Seri    string `json:"seri"`
	TxHash  string `json:"tx_hash"`
	IPFSCid string `json:"ipfs_cid"`
	TokenID uint64 `json:"token_id"`
}

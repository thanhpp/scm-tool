package nftsvclient

import (
	"errors"
	"strconv"

	"github.com/thanhpp/scm/internal/scmsrv/domain/entity"
	"github.com/thanhpp/scm/pkg/ginutil"
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
		return newError("set req create seri nft", "nil check", errors.New("serial or import ticket is nil"))
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

type RespCreateSeriNFT struct {
	ginutil.RespTemplateError
	Data struct {
		Seri    string `json:"seri"`
		TxHash  string `json:"tx_hash"`
		IPFSCid string `json:"ipfs_cid"`
	} `json:"data"`
}

type NFTInfo struct {
	Seri    string `json:"seri"`
	TxHash  string `json:"tx_hash"`
	IPFSCid string `json:"ipfs_cid"`
	TokenID uint64 `json:"token_id"`
}

type RespGetNFTInfo struct {
	ginutil.RespTemplateError
	Data NFTInfo `json:"data"`
}

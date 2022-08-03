package dto

import (
	"github.com/thanhpp/scm/internal/nftsrv/domain/entity"
	"github.com/thanhpp/scm/pkg/ginutil"
)

type MintSeriNFTReq struct {
	Seri     string            `json:"seri"`
	Metadata map[string]string `json:"metadata"`
}

type DataMintSeriNFTResp struct {
	Seri    string `json:"seri"`
	TxHash  string `json:"tx_hash"`
	IPFSCid string `json:"ipfs_cid"`
}

func (d *DataMintSeriNFTResp) set(in *entity.SerialNFT) {
	d.Seri = in.Seri
	d.TxHash = in.TxHash
	d.IPFSCid = in.IPFSHash
}

type MintSeriNFTResp struct {
	ginutil.RespTemplateError
	Data DataMintSeriNFTResp `json:"data"`
}

func (r *MintSeriNFTResp) SetData(in *entity.SerialNFT) {
	r.Data.set(in)
}

type BurnSeriNFTsReq struct {
	Serials []string `json:"serials"`
}

type BurnSerisNFTResp struct {
	ginutil.RespTemplateError
}

type TransferSeriNFTReq struct {
	Serials []string `json:"serials"`
	To      string   `json:"to"`
}

type TransferSeriNFTResp struct {
	ginutil.RespTemplateError
}

type DataGetSeriNFTResp struct {
	DataMintSeriNFTResp
	TokenID uint64 `json:"token_id"`
}

func (d *DataGetSeriNFTResp) set(in *entity.SerialNFT) {
	d.DataMintSeriNFTResp.set(in)
	d.TokenID = uint64(in.TokenID)
}

type GetSeriNFTResp struct {
	ginutil.RespTemplateError
	Data struct {
		DataGetSeriNFTResp
		Owner          string `json:"owner"`
		TransferTxHash string `json:"transfer_tx_hash"`
	} `json:"data"`
}

func (r *GetSeriNFTResp) SetData(in *entity.SerialNFT) {
	r.Data.set(in)
	r.Data.Owner = in.Owner
	r.Data.TransferTxHash = in.TransferTxHash
}

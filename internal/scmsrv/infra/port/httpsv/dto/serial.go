package dto

import (
	"github.com/thanhpp/scm/internal/scmsrv/domain/entity"
	"github.com/thanhpp/scm/internal/scmsrv/infra/adapter/nftsvclient"
	"github.com/thanhpp/scm/pkg/ginutil"
)

type SerialInfoRespItemData struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type SerialInfoRespStorageData struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

type SerialInfoRespSupplierData struct {
	Name string `json:"name"`
}

type SerialInfoRespData struct {
	Seri           string                     `json:"seri"`
	ImportTicketID int                        `json:"import_ticket_id"`
	StorageData    SerialInfoRespStorageData  `json:"storage"`
	SupplierData   SerialInfoRespSupplierData `json:"supplier"`
	ItemData       SerialInfoRespItemData     `json:"item"`
	NFTInfo        nftsvclient.NFTInfo        `json:"nft_info"`
}

func (d *SerialInfoRespData) set(in *entity.Serial, nftInfo *nftsvclient.NFTInfo) {
	d.Seri = in.Seri
	d.ImportTicketID = in.ImportTicket.ID

	d.StorageData = SerialInfoRespStorageData{
		Name:     in.ImportTicket.ToStorage.Name,
		Location: in.ImportTicket.ToStorage.Location,
	}

	d.ItemData = SerialInfoRespItemData{
		Name: in.Item.Name,
		Desc: in.Item.Desc,
	}

	d.SupplierData = SerialInfoRespSupplierData{
		Name: in.ImportTicket.FromSupplier.Name,
	}

	d.NFTInfo = *nftInfo
}

type SerialInfoResp struct {
	ginutil.RespTemplateError
	Data SerialInfoRespData `json:"data"`
}

func (resp *SerialInfoResp) SetData(in *entity.Serial, nftInfo *nftsvclient.NFTInfo) {
	resp.Data.set(in, nftInfo)
}

type SerialInfoRespDataWithoutNFTInfo struct {
	Seri           string                     `json:"seri"`
	ImportTicketID int                        `json:"import_ticket_id"`
	StorageData    SerialInfoRespStorageData  `json:"storage"`
	SupplierData   SerialInfoRespSupplierData `json:"supplier"`
	ItemData       SerialInfoRespItemData     `json:"item"`
}

type GetSerialsByImportTicketIDResp struct {
	ginutil.RespTemplateError
	Data []GetSerialsByImportTicketIDData `json:"data"`
}

type GetSerialsByImportTicketIDData struct {
	SKU    string   `json:"sku"`
	Series []string `json:"series"`
}

func (resp *GetSerialsByImportTicketIDResp) SetData(m map[string][]*entity.Serial) {
	for k, v := range m {
		resp.Data = append(resp.Data, GetSerialsByImportTicketIDData{
			SKU:    k,
			Series: make([]string, len(v)),
		})
		for i := range v {
			resp.Data[len(resp.Data)-1].Series[i] = v[i].Seri
		}
	}
}

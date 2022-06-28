package dto

import (
	"github.com/thanhpp/scm/internal/scmsrv/domain/entity"
	"github.com/thanhpp/scm/pkg/ginutil"
)

type CreateStorageReq struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Desc     string `json:"desc"`
}

type StorageInfoRespData struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Desc     string `json:"desc"`
}

func (d *StorageInfoRespData) set(stg *entity.Storage) {
	d.ID = stg.ID
	d.Name = stg.Name
	d.Location = stg.Location
	d.Desc = stg.Desc
}

type StorageInfoResp struct {
	ginutil.RespTemplateError
	Data StorageInfoRespData `json:"data"`
}

func (resp *StorageInfoResp) SetData(stg *entity.Storage) {
	resp.Data.set(stg)
}

type GetListStoragesResp struct {
	ginutil.RespTemplateError
	Data []StorageInfoRespData `json:"data"`
}

func (resp *GetListStoragesResp) SetData(storages []*entity.Storage) {
	resp.Data = make([]StorageInfoRespData, len(storages))

	for i := range storages {
		resp.Data[i].set(storages[i])
	}
}

type StorageDetailRespData struct {
	Storage        StorageInfoRespData `json:"storage"`
	AvailableItems int                 `json:"available_items"`
}

func (data *StorageDetailRespData) set(stg *entity.Storage, availableItems int) {
	data.Storage.set(stg)
	data.AvailableItems = availableItems
}

type StorageDetailResp struct {
	ginutil.RespTemplateError
	Data StorageDetailRespData `json:"data"`
}

func (resp *StorageDetailResp) SetData(stg *entity.Storage, availableItems int) {
	resp.Data.set(stg, availableItems)
}

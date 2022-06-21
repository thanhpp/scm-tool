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
	Name     string `json:"name"`
	Location string `json:"location"`
	Desc     string `json:"desc"`
}

type StorageInfoResp struct {
	ginutil.RespTemplateError
	Data StorageInfoRespData `json:"data"`
}

func (resp *StorageInfoResp) SetData(stg *entity.Storage) {
	resp.Data = StorageInfoRespData{
		Name:     stg.Name,
		Location: stg.Location,
		Desc:     stg.Desc,
	}
}

type GetListStoragesResp struct {
	ginutil.RespTemplateError
	Data []StorageInfoRespData `json:"data"`
}

func (resp *GetListStoragesResp) SetData(storages []*entity.Storage) {
	resp.Data = make([]StorageInfoRespData, len(storages))

	for i := range storages {
		resp.Data[i] = StorageInfoRespData{
			Name:     storages[i].Name,
			Location: storages[i].Location,
			Desc:     storages[i].Desc,
		}
	}
}

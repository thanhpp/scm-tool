package dto

import (
	"github.com/thanhpp/scm/internal/scmsrv/domain/entity"
	"github.com/thanhpp/scm/pkg/ginutil"
)

type CreateSupplierReq struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type SupplierInfoRespData struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

func (d *SupplierInfoRespData) set(supplier *entity.Supplier) {
	d.ID = supplier.ID
	d.Name = supplier.Name
	d.Phone = supplier.Phone
	d.Email = supplier.Email
}

type SupplierInfoResp struct {
	ginutil.RespTemplateError
	Data SupplierInfoRespData `json:"data"`
}

func (resp *SupplierInfoResp) SetData(supplier *entity.Supplier) {
	resp.Data.set(supplier)
}

type GetListSupplierResp struct {
	ginutil.RespTemplateError
	Data []SupplierInfoRespData `json:"data"`
}

func (resp *GetListSupplierResp) SetData(suppliers []*entity.Supplier) {
	resp.Data = make([]SupplierInfoRespData, len(suppliers))

	for i := range suppliers {
		resp.Data[i].set(suppliers[i])
	}
}

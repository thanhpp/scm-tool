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

type SupplierInfoResp struct {
	ginutil.RespTemplateError
	Data SupplierInfoRespData `json:"data"`
}

func (resp *SupplierInfoResp) SetData(supplier *entity.Supplier) {
	resp.Data = SupplierInfoRespData{
		ID:    supplier.ID,
		Name:  supplier.Name,
		Phone: supplier.Phone,
		Email: supplier.Email,
	}
}

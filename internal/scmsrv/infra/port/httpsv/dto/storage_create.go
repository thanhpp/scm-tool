package dto

type CreateStorageReq struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Desc     string `json:"desc"`
}

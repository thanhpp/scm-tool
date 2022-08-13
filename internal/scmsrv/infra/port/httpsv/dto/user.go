package dto

import (
	"github.com/thanhpp/scm/internal/scmsrv/domain/entity"
	"github.com/thanhpp/scm/pkg/ginutil"
)

type CreateUserReq struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserInfoRespData struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

func (d *UserInfoRespData) set(user *entity.User) {
	d.ID = user.ID
	d.Name = user.Name
	d.Username = user.Username
}

type UserInfoResp struct {
	ginutil.RespTemplateError
	Data UserInfoRespData `json:"data"`
}

func (resp *UserInfoResp) SetData(user *entity.User) {
	resp.Data.set(user)
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRespData struct {
	Token string `json:"token"`
	User  struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"user"`
}

type LoginResp struct {
	ginutil.RespTemplateError
	Data LoginRespData `json:"data"`
}

func (resp *LoginResp) SetData(token string, user *entity.User) {
	resp.Data.Token = token
	resp.Data.User.ID = user.ID
	resp.Data.User.Name = user.Name
}

type RespGetUsers struct {
	ginutil.RespTemplateError
	Data []UserInfoRespData `json:"data"`
}

func (resp *RespGetUsers) SetData(users []*entity.User) {
	resp.Data = make([]UserInfoRespData, len(users))
	for i := range resp.Data {
		resp.Data[i].set(users[i])
	}
}

type ReqUpdateUserPass struct {
	NewPass string `json:"new_pass"`
}

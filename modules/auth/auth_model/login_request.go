package auth_model

import (
	"relia_system/modules/user/user_model"
)

type LoginRequest struct {
	Cmt      string `json:"cmt"`
	Cccd     string `json:"cccd"`
	LoginId  string `json:"loginId" form:"loginId"`
	UseCmt   bool   `json:"use_cmt"`
	Password string `json:"password" form:"password"`
}

func (LoginRequest) TableName() string {
	return user_model.User{}.TableName()
}

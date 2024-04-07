package response

import "web2app/model"

type LoginResponse struct {
	User      model.SysUserRes `json:"user"`
	Token     string           `json:"token"`
	ExpiresAt int64            `json:"expiresAt"`
}

package model

import "time"

type JwtBlacklist struct {
	CreateTime time.Time `json:"createTime" bson:"createTime"`
	Jwt        string    `json:"jwt" bson:"jwt"`
}

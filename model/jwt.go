package model

import (
	"github.com/golang-jwt/jwt/v4"
)

// Custom claims structure
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.StandardClaims
}

type BaseClaims struct {
	Email    string
	ID       string
	UserName string
	NickName string
	RoleId   string
}

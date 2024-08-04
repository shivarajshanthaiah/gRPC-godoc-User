package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	UserID int    `json:"userid"`
	Email  string `json:"email"`
	Role   string `json:"Role"`
	Claims jwt.StandardClaims
}

func (*Claims) Valid() error {
	panic("unimplemented")
}
package client

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

//Client is the struct that holds client information
type Client struct {
	gorm.Model
	FirstName string `gorm:"first_name" json:"first_name,omitempty"`
	LastName  string `gorm:"last_name" json:"last_name,omitempty"`
	Email     string `gorm:"email" json:"email,omitempty"`
	Secret    string `gorm:"secret" json:"secret,omitempty"`
}

//NewClient creates a new model
func NewClient() *Client {
	c := new(Client)
	return c
}

type UserToken struct {
	UID string `json:"id"`
	jwt.StandardClaims
}

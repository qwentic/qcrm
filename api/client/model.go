package client

import "github.com/jinzhu/gorm"

//Client is the struct that holds client information
type Client struct {
	gorm.Model
	FirstName string `gorm:"first_name"		json:"first_name"`
	LastName  string `gorm:"last_name"		json:"last_name"`
	Email     string `gorm:"email"		json:"email"`
	Secret    string `gorm: "secret"   json:"secret"`
}

func NewClient() *Client {
	c := new()
}

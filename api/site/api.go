package site

import "github.com/jinzhu/gorm"

var db *gorm.DB

//API holds all api handler
type API struct {
	DB *gorm.DB
}

//NewAPI creates a new api and insert the database provided in it
func NewAPI(d *gorm.DB) *API {
	a := new(API)
	a.DB = d
	db = d
	return a
}

package site

import "github.com/jinzhu/gorm"

//Site holds site information
type Site struct {
	gorm.Model
	Name      string `sql:"name" json:"name,omitempty"`
	URL       string `sql:"url" json:"url,omitempty"`
	ContactID uint   `sql:"contact_id" json:"contact_id,omitempty"`
}

//NewSite creates a new site
func NewSite() *Site {
	s := new(Site)
	return s
}

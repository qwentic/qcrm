package contact

import "github.com/jinzhu/gorm"

//Contact holds api(s)
type Contact struct {
	gorm.Model
	FirstName string `gorm:"first_name" json:"first_name,omitempty"`
	LastName  string `gorm:"last_name" json:"last_name,omitempty"`
	Image     string `gorm:"image" json:"image,omitempty"`
	Email     string `gorm:"email" json:"email,omitempty"`
	Phone     string `gorm:"phone" json:"phone,omitempty"`
	CompanyID uint   `gorm:"company_id" json:"company_id,omitempty"`
}

//NewContact creates a new model
func NewContact() *Contact {
	c := new(Contact)
	return c
}

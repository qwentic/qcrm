package company

import "github.com/jinzhu/gorm"

//Company holds information for company
type Company struct {
	gorm.Model
	Name         string  `gorm:"name" json:"name,omitempty"`
	Responsible  string  `gorm:"responsible" json:"responsible,omitempty"`
	Type         string  `gorm:"type" json:"type,omitempty"`
	IndustryID   uint    `gorm:"industry_id" json:"industry_id,omitempty"`
	Employees    string  `gorm:"employees" json:"employees,omitempty"`
	AnnualIncome float64 `gorm:"annual_income" json:"annual_income,omitempty"`
	Currency     string  `gorm:"currency" json:"currency,omitempty"`
	Comment      string  `gorm:"comment" json:"comment,omitempty"`
}

//NewCompany creates a new model
func NewCompany() *Company {
	c := new(Company)
	return c
}

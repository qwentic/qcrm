package industry

import "github.com/jinzhu/gorm"

//Industry holds indutries details to insert into database
type Industry struct {
	gorm.Model
	Name string `gorm:"name" json:"name,omitempty"`
}

//NewIndustry creates a new model
func NewIndustry() *Industry {
	indus := new(Industry)
	return indus
}

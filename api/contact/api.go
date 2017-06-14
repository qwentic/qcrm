package contact

import (
	"errors"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/UnnoTed/govalidator"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/qwentic/qcrm/api/response"
)

var db *gorm.DB

// API holds all related api handlers
type API struct {
	DB *gorm.DB
}

// NewAPI creates a new api and insert the
// database provided in it
func NewAPI(d *gorm.DB) *API {
	a := new(API)
	a.DB = d
	db = d
	return a

}

//Contactinfo struct to hold post request for contact
type Contactinfo struct {
	FirstName string `json:"first_name,omitempty" valid:"alphanum,length(2|100),required"`
	LastName  string `json:"last_name,omitempty" valid:"alphanum,length(2|100),required"`
	Email     string `json:"email,omitempty" valid:"email,length(6|255),required"`
	Phone     string `json:"phone,omitempty" valid:"number,length(10|12),required"`
}

//PostContact handles post request to create new contact
func (a *API) PostContact(c echo.Context) error {
	cc := &Contactinfo{}
	if err := c.Bind(&cc); err != nil {
		return response.Error(c, err)
	}
	log.SetLevel(log.DebugLevel)
	log.WithFields(log.Fields{
		"FirstName": cc.FirstName,
		"LastName":  cc.LastName,
		"Email":     cc.Email,
	}).Debug("[PostContact] data:")

	// validates cc
	valid, err := govalidator.ValidateStruct(cc)
	if err != nil {
		return response.Error(c, err)
	}

	// check if usr is valid
	if !valid {
		return response.Error(c, errors.New("The data you sent is not valid"))
	}

	//copy cc data into a new contact
	contact := NewContact()

	contact.FirstName = cc.FirstName
	contact.LastName = cc.LastName
	contact.Email = cc.Email
	contact.Phone = cc.Phone
	//check if the email exist in the database by counting the rows count
	ct := NewContact()
	var count int
	a.DB.Where("email = ?", contact.Email).Find(&ct).Count(&count)
	if count > 0 {
		return errors.New("Email already exists")
	}

	// insert the user into the database
	db.Create(&contact)
	if db.NewRecord(contact) {
		return response.Error(c, errors.New("Fail to insert record"))
	}

	return response.SuccessCreated(c, map[string]interface{}{
		"id": strconv.FormatInt(int64(27), 10),
	})
}

//PutContact handles put request to update the existing contact
func (a *API) PutContact(c echo.Context) error {

	return response.SuccessCreated(c, map[string]interface{}{
		"id": strconv.FormatInt(int64(27), 10),
	})
}

//GetContact handles get request to return existing contact
func (a *API) GetContact(c echo.Context) error {

	return response.SuccessCreated(c, map[string]interface{}{
		"id": strconv.FormatInt(int64(27), 10),
	})
}

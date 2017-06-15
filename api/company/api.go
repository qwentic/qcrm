package company

import (
	"errors"
	"strconv"

	"github.com/UnnoTed/govalidator"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/qwentic/qcrm/api/response"
)

var db *gorm.DB

//API holds all related api(s) handler
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

//PostCompany handles post request to create new company
func (a *API) PostCompany(c echo.Context) error {
	cmp := &Company{}
	if err := c.Bind(cmp); err != nil {
		return response.Error(c, err)
	}
	log.Info("test in company")
	valid, err := govalidator.ValidateStruct(cmp)

	if err != nil {
		return response.Error(c, err)
	}

	if !valid {
		return response.Error(c, errors.New("The data you sent is not valid"))

	}
	//check if company already exists
	if !a.DB.Where("name = ?", cmp.Name).First(&cmp).RecordNotFound() {
		return response.Error(c, errors.New("Company exists"))
	}
	cdb := a.DB.Create(&cmp)
	if cdb.Error != nil {
		response.Error(c, cdb.Error)
	}
	return response.SuccessCreated(c, map[string]interface{}{
		"id": strconv.FormatInt(int64(cmp.ID), 10),
	})
}

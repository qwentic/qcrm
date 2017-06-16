package contact

import (
	"errors"
	"strconv"

	Log "github.com/Sirupsen/logrus"
	"github.com/UnnoTed/govalidator"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/qwentic/qcrm/api/industry"
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

// Init initiates the process
func Init(d *gorm.DB) error {
	Log.SetLevel(Log.DebugLevel)
	Log.Debug("[Algorithm.Init]: Initiating for saving industries to database...")

	var err error

	if d != nil {
		Log.Debug("[Algorithm.Init]: MySQL: OK")
		db = d
	} else {
		return errors.New("No db found")
	}

	// industries
	err = PreDefinedIndustries()
	if err != nil {
		return err
	}

	Log.Debug("[Algorithm.Init]: Init for industries insertion Done .")
	return nil
}

func PreDefinedIndustries() error {
	//create transaction
	tx := db.Begin()

	for _, indus := range industry.All {
		_indus := industry.NewIndustry()
		indb := db.Where("name = ?", indus.Name).First(&_indus)
		if indb.Error != nil && !indb.RecordNotFound() {
			return indb.Error
		}
		//ignore when exists
		if !indb.RecordNotFound() {
			continue
		}
		if err := tx.Create(&indus).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// insert industries into the db
	tx.Commit()

	return nil
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
	/*userID, err := util.GetUserID(c)
	if err != nil {
		return response.Error(c, err)
	}
	//get name of responsible person to create contact
	cl := client.NewClient()

	resdb := a.DB.Where("id = ?", userID).First(&cl)
	if resdb.Error != nil && !resdb.RecordNotFound() {
		return response.Error(c, resdb.Error)
	}
	*/
	cc := &CSInfo{}
	if err := c.Bind(&cc); err != nil {
		return response.Error(c, err)
	}
	Log.SetLevel(Log.DebugLevel)
	Log.WithFields(Log.Fields{
		"FirstName": cc.FirstName,
		"LastName":  cc.LastName,
		"Email":     cc.Email,
	}).Debug("[PostContact] data:")

	// validates cc
	valid, err := govalidator.ValidateStruct(cc)
	if err != nil {
		return response.Error(c, err)
	}

	// check if cc is valid
	if !valid {
		return response.Error(c, errors.New("The data you sent is not valid"))
	}

	//copy cc data into a new contact
	/*	contact := NewContact()

		contact.FirstName = cc.FirstName
		contact.LastName = cc.LastName
		contact.Email = cc.Email
		contact.Phone = cc.Phone
		contact.CompanyID = cc.CompanyID
	*/ //check if the email exist in the database by counting the rows count
	ct := NewContact()
	var count int
	a.DB.Where("email = ?", cc.Email).Find(&ct).Count(&count)
	if count > 0 {
		return errors.New("Email already exists")
	}

	// insert the contact into the database
	db.Create(&cc.Contact)
	if db.NewRecord(cc.Contact) {
		return response.Error(c, errors.New("Fail to insert record"))
	}
	tx := a.DB.Begin()
	//insert sites url
	if cc.Site != nil {
		for _, s := range cc.Site {
			s.ContactID = cc.ID
			if err := tx.Create(&s).Error; err != nil {
				return response.Error(c, err)
			}
		}
		err := tx.Commit().Error
		if err != nil {
			return response.Error(c, err)
		}
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

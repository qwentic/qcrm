package client

import (
	"strconv"

	"errors"

	log "github.com/Sirupsen/logrus"
	"github.com/UnnoTed/govalidator"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/qwentic/qcrm/api/response"
	"github.com/qwentic/qcrm/api/util"
	"golang.org/x/crypto/bcrypt"
)

const (
	issuer = "auth.service"
)

var (
	db *gorm.DB
)

// API holds all related api handlers
type API struct {
	DB *gorm.DB
}

// NewAPI creates a new api and insert the
// database provided in it
func NewAPI(d *gorm.DB) *API {
	log.Debug("api")
	a := new(API)
	a.DB = d
	db = d
	return a
}

// Register handles post requests to register a new user account
type Register struct {
	FirstName string `json:"first_name" valid:"alphanum,length(2|100),required"`
	LastName  string `json:"last_name"  valid:"alphanum,length(2|100),required"`
	Email     string `json:"email"      valid:"email,length(6|255),required"`
	Secret    string `json:"password"   valid:"length(6|255),required"`

	Admin bool `json:"admin,omitempty"`
}

// PostRegister handles post requests to register a new user account
func (a *API) PostRegister(c echo.Context) error {
	usr := &Register{}
	// insert the body data into usr
	if err := c.Bind(&usr); err != nil {
		return response.Error(c, err)
	}

	log.SetLevel(log.DebugLevel)
	log.WithFields(log.Fields{
		"FirstName": usr.FirstName,
		"Secret":    "secret",
		"Email":     usr.Email,
	}).Debug("[PostRegister]: data")

	// validates usr
	valid, err := govalidator.ValidateStruct(usr)
	if err != nil {
		return response.Error(c, err)
	}

	// check if usr is valid
	if !valid {
		return response.Error(c, errors.New("The data you sent is not valid"))
	}

	// check for weak password
	/*if dry.StringListContainsCaseInsensitive(PasswordBlackList, usr.Secret) {
		return response.Error(c, errors.ErrorWeakPassword)
	}*/

	// hashes the password, uses the default cost (10) although that can be changed
	pwd, err := bcrypt.GenerateFromPassword([]byte(usr.Secret), bcrypt.DefaultCost)
	if err != nil {
		return response.Error(c, err)
	}

	// insert the hashed password into usr
	usr.Secret = string(pwd)

	// copy usr data into a new client
	cl := NewClient()

	cl.FirstName = usr.FirstName
	cl.LastName = usr.LastName
	cl.Secret = usr.Secret
	cl.Email = usr.Email
	//cl.Admin = usr.Admin

	// checks if the email exists in the database
	// by counting the rows found
	ce := NewClient()
	var count int
	a.DB.Where("email = ?", cl.Email).Find(&ce).Count(&count)
	if count > 0 {
		return errors.New("Email already exists")
	}

	// insert the user into the database
	db.Create(&cl)
	if db.NewRecord(cl) {
		return response.Error(c, errors.New("Fail to insert record"))
	}

	/*	// verification
		vrf := &verification.Verification{
			TimeCreated: time.Now().Unix(),
			ClientID:    cl.ID,
			Code:        dry.RandomHexString(10),
		}

		// check if verification exists
		dvrf := &verification.Verification{}
		count, err := a.Model(dvrf).Where("code = ?", vrf.Code).Count()
		if err != nil && err != pg.ErrNoRows {
			return response.Error(c, err)
		}

		// when exists, create a new code
		if count > 0 {
			vrf.Code = dry.RandomHexString(20)
		}

		// create verification
		if err = a.DB.Create(&vrf); err != nil {
			return response.Error(c, err)
		}
	*/
	// create the url to activate
	// url := "http://www.localhost/engagehits_fe/activate/" + vrf.Code
	// get client's info
	/*	sibClient := &sendinblue.Client{
			FirstName: cl.FirstName,
			Email:     cl.Email,
		}
		body := "In order to complete your registration at EngagedHits.com you need to confirm your email  <br> by clicking"
		body += "<a href=" + url + "> here </a>or simply click in the link below:<br><br>" + url
		body += "<br><br><br>Thank you"
		body += "<br><br>Engaged Hits"

		// send the email
		err = sendinblue.Send(sibClient, "Activate your account at Engaged Hits", body, sendinblue.TemplateActivation)
		if err != nil {
			Log.Debug("SendInBlue Error")
			return response.Error(c, err)
		}
	*/ //}

	// respond created
	return response.SuccessCreated(c, map[string]interface{}{
		"id": strconv.FormatInt(int64(cl.ID), 10),
	})
}

//PostLogin handles post requests to login a existing user
func (a *API) PostLogin(c echo.Context) error {
	var usr = &struct {
		Email  string `json:"email"      valid:"email,length(6|255),required"`
		Secret string `json:"password"   valid:"length(6|255),required"`
	}{}

	if err := c.Bind(usr); err != nil {
		return response.Error(c, err)
	}

	log.SetLevel(log.DebugLevel)
	log.WithField("email", usr.Email).Debug("[PostLogin]: data")

	// validates usr
	valid, err := govalidator.ValidateStruct(usr)
	if err != nil {
		return response.Error(c, err)
	}

	// check if usr is valid
	if !valid {
		return response.Error(c, errors.New("the data you sent is not valid"))
	}

	// get the user from the database usr's email
	cl := NewClient()
	res := a.DB.Where("email = ?", usr.Email).First(&cl)
	if res.Error != nil {
		return response.Error(c, errors.New("Your email or password is wrong"))
	}

	// compare the password
	err = bcrypt.CompareHashAndPassword([]byte(cl.Secret), []byte(usr.Secret))
	if err != nil {
		// wrong password
		return response.Error(c, errors.New("Your email or password is wrong"))
	}

	// encrypt the uuid and create a new token
	tokenString, err := CreateToken(cl.ID, true)
	if err != nil {
		return response.Error(c, err)
	}

	// remove port from the remote address
	/*ip := c.Request().RemoteAddress()
	s := strings.Split(ip, ":")
	ip = s[0]

	// get user's ip
	req := goreq.Request{
		Method: "POST",
		Uri:    "http://ip-api.com/json/" + ip,
	}

	// get browser and kind of device from user agent
	_, _, _, _, _, deviceType, _ := uasurfer.Parse(c.Request().UserAgent())
	var accessFrom string

	// chose device type
	switch deviceType {
	case uasurfer.DevicePhone, uasurfer.DeviceTablet:
		accessFrom = "android"
	default:
		accessFrom = "desktop"
	}

	ss := &session.Session{
		AccessFrom:  accessFrom,
		Token:       tokenString,
		ClientID:    cl.ID,
		Valid:       true,
		TimeCreated: time.Now().Unix(),
		TimeExpires: time.Now().Add(48 * time.Hour).Unix(),
		MachineID:   c.Request().UserAgent(),
		MachineIP:   ip,
	}

	err = ss.Create()
	if err != nil {
		return response.Error(c, err)
	}
	*/
	// respond ok
	return response.Success(c, map[string]interface{}{
		"token": tokenString,
		"id":    strconv.FormatInt(util.Obfuscate(cl.ID), 10),
	})
}

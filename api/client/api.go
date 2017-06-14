package client

import (
	"fmt"
	"strconv"
	"time"

	"bitbucket.org/mirkorakic/engagedhits/api/referral"
	"bitbucket.org/mirkorakic/engagedhits/api/response"
	"bitbucket.org/mirkorakic/engagedhits/api/verification"
	"bitbucket.org/mirkorakic/engagedhits/config"
	"bitbucket.org/mirkorakic/engagedhits/errors"
	"bitbucket.org/mirkorakic/engagedhits/sendinblue"
	"bitbucket.org/mirkorakic/engagedhits/util"
	"github.com/Sirupsen/logrus"
	"github.com/UnnoTed/govalidator"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	dry "github.com/ungerik/go-dry"
	"golang.org/x/crypto/bcrypt"
	pg "gopkg.in/pg.v4"
)

var (
	db  *gorm.DB
	Log = logrus.New()
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

	Admin bool `json:"admin"`
}

// PostRegister handles post requests to register a new user account
func (a *API) PostRegister(c echo.Context) error {
	usr := &Register{}

	// insert the body data into usr
	if err := c.Bind(&usr); err != nil {
		return response.Error(c, err)
	}

	Log.WithFields(logrus.Fields{
		"first_name": usr.FirstName,
		"password":   "secret",
		"email":      usr.Email,
	}).Debug("[PostRegister]: data")

	// validates usr
	valid, err := govalidator.ValidateStruct(usr)
	if err != nil {
		return response.Error(c, err)
	}

	// check if usr is valid
	if !valid {
		return response.Error(c, errors.ErrorInvalid)
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

	// get the client level
	if cl.ClientLevelID, err = DefaultClientLevel(); err != nil {
		return response.Error(c, err)
	}

	cl.FirstName = usr.FirstName
	cl.LastName = usr.LastName
	cl.Secret = usr.Secret
	cl.Email = usr.Email
	cl.Admin = usr.Admin
	cl.Trial = true

	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("Cookies : ", c.Cookies())
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	// referral from cookie
	for _, cookie := range c.Cookies() {
		// client id
		if cookie.Name() == "cref" {
			Log.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
			Log.WithField("id", cookie.Value()).Debug("[PostRegister]: Found referral")
			Log.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
			// parse string id to int64
			id, err := strconv.ParseInt(cookie.Value(), 10, 64)
			if err != nil {
				return response.Error(c, err)
			}

			t := time.Now().Unix()

			// create a referral
			ref := &referral.Referral{
				TimeCreated:  t,
				TimeAccepted: t,
				Hashed:       dry.RandomHexString(32),
				Success:      true,
				ClientID:     util.Deobfuscate(id),
			}

			if err = ref.Create(); err != nil {
				return response.Error(c, err)
			}

			// insert referral's id into client
			cl.ReferralID = ref.ID
			break

			// referral id
		} else if cookie.Name() == "ref" {
			Log.WithField("id", cookie.Value()).Debug("[PostRegister]: Found referral")

			ref := &referral.Referral{}
			err = ref.Query().Where("hashed = ?", cookie.Value()).First()
			if err != nil && err != pg.ErrNoRows {
				return response.Error(c, err)
			}

			// cancel the referral when the ID is wrong
			// because it would only cause problems
			// if give a error that is hard to fix to a user
			// like cleaning cookies
			if err == pg.ErrNoRows {
				break
			}

			// the referral was found, but it is already used
			// when success is true, so it'll create a new referral
			// and give it to the user
			if ref.Success {
				t := time.Now().Unix()

				// create a referral
				ref = &referral.Referral{
					TimeCreated:  t,
					TimeAccepted: t,
					Hashed:       dry.RandomHexString(32),
					Success:      true,

					// same client's id as the used referral
					ClientID: ref.ClientID,
				}

				// create referral
				if err = ref.Create(); err != nil {
					return response.Error(c, err)
				}

				// referral wasn't used yet
				// change it's state to used
			} else {
				ref.TimeAccepted = time.Now().Unix()
				ref.Success = true
			}

			// update the referral
			if err = ref.Update(); err != nil {
				return response.Error(c, err)
			}

			// insert referral's id into client
			cl.ReferralID = ref.ID
			break
		}
	}

	if !config.Live {
		// cl.Verified = true

		// fake verification
		vrf := &verification.Verification{
			TimeVerified: time.Now().Unix(),
			TimeCreated:  time.Now().Unix(),
			ClientID:     cl.ID,
			Code:         dry.RandomHexString(32),
		}

		err = a.DB.Create(&vrf)
		if err != nil {
			return response.Error(c, err)
		}
	}

	// insert the user into the database
	if err = cl.Create(); err != nil {
		return response.Error(c, err)
	}

	//if config.Live {
	// verification
	vrf := &verification.Verification{
		TimeCreated: time.Now().Unix(),
		ClientID:    cl.ID,
		Code:        dry.RandomHexString(10),
	}

	// check if verification exists
	dvrf := &verification.Verification{}
	count, err := a.DB.Model(dvrf).Where("code = ?", vrf.Code).Count()
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

	// create the url to activate
	activateUrl := viper.GetString("activate_url")
	url := activateUrl + vrf.Code
	// url := "http://www.localhost/engagehits_fe/activate/" + vrf.Code
	// get client's info
	sibClient := &sendinblue.Client{
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
	//}

	// respond created
	return response.SuccessCreated(c, map[string]interface{}{
		"id": strconv.FormatInt(util.Obfuscate(cl.ID), 10),
	})
}

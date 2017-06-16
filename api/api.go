package api

import (
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/qwentic/qcrm/api/client"
	"github.com/qwentic/qcrm/api/company"
	"github.com/qwentic/qcrm/api/contact"
	"github.com/qwentic/qcrm/config"
)

var (
	db         *gorm.DB
	clientAPI  *client.API
	contactAPI *contact.API
	dbmu       = &sync.Mutex{}
)

// Database connects to the database specified through env vars
func Database() error {
	// check if it already is connected
	if db != nil {
		return nil
	}

	dbmu.Lock()
	defer dbmu.Unlock()
	ConStr := config.DBUser + ":" + config.DBPass + "@/" + config.DBName + "?" + config.DBParam
	//gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	db, _ = gorm.Open("mysql", ConStr)

	return nil
}

// Setup initiates the database connection and
// send it to each sub-api
// it also creates all routes' handlers
func Setup(ee **echo.Echo) error {

	e := *ee

	err := Database()
	if err != nil {
		return err
	}

	dbmu.Lock()
	defer dbmu.Unlock()

	clientAPI = client.NewAPI(db)
	contactAPI = contact.NewAPI(db)
	companyAPI := company.NewAPI(db)
	e.Use(middleware.Logger())

	_api := e.Group("/api")
	{
		c1 := _api.Group("/qw")
		{
			//user
			c1.POST("/register", clientAPI.PostRegister)
			c1.POST("/login", clientAPI.PostLogin)
			_contact := _api.Group("/qw")
			{
				// auth
				_contact.Use(client.JWT([]byte(config.JWTSecret)))
				//	_contact.Use(clientAPI.Middleware(false))

				_contact.POST("/contact", contactAPI.PostContact)
				_contact.PUT("/contact/:id", contactAPI.PutContact)
				_contact.GET("/contact/:id", contactAPI.GetContact)
				_contact.POST("/company", companyAPI.PostCompany)
			}
		}
	}

	//	defer db.Close()

	return nil
}

func GetDB() *gorm.DB {
	if db == nil {
		_ = Database()
	}

	return db
}

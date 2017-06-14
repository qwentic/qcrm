package api

import (
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/qwentic/qcrm/api/client"
	"github.com/qwentic/qcrm/config"
)

var (
	db        *gorm.DB
	clientAPI *client.API

	dbmu = &sync.Mutex{}
)

// Database connects to the database specified through env vars
func Database(test bool) error {
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
func Setup(ee **echo.Echo, test bool) error {

	e := *ee

	err := Database(test)
	if err != nil {
		return err
	}

	dbmu.Lock()
	defer dbmu.Unlock()

	clientAPI = client.NewAPI(db)

	_api := e.Group("/api")
	{
		c1 := _api.Group("/qw")
		{
			//user
			c1.POST("/register", clientAPI.PostRegister)
		}
	}

	//	defer db.Close()

	return nil
}

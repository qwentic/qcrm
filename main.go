package main

import (
	"runtime"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/gommon/log"
	"github.com/qwentic/qcrm/api"
	"github.com/qwentic/qcrm/api/contact"
	"github.com/qwentic/qcrm/config"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	e := echo.New()
	e.SetDebug(true)

	err := api.Setup(&e)
	handleErr(err)

	err = contact.Init(api.GetDB())
	handleErr(err)
	log.Info("[Main]: OK")
	log.Info("Port:", config.Port)
	e.Run(fasthttp.New(":" + config.Port))
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

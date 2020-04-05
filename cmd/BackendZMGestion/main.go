package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

func main() {

	conf, err := initConfig()
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.Logger.Fatal(e.Start(conf.Server.Listen))
}

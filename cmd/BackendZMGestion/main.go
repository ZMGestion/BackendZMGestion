package main

import (
	"BackendZMGestion/src/db"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

func main() {

	conf, err := initConfig()
	if err != nil {
		panic(err)
	}

	dbConfig := conf.Database

	db, err := db.InitDb(dbConfig.User, dbConfig.Pass, dbConfig.Host, dbConfig.Port, dbConfig.Schema)
	if err != nil {
		panic(err)
	}
	defer db.Conn.Close()

	e := echo.New()
	e.Logger.Fatal(e.Start(conf.Server.Listen))
}

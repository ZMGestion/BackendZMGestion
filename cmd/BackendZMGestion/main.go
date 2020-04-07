package main

import (
	"BackendZMGestion/src/db"
	"fmt"

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

	/*Borrar desde aqui*/
	out, err := db.CallSP("zsp_roles_listar", nil)

	if err != nil {
		panic(err)
	}

	fmt.Println(out)
	/*Hasta aqui*/

	e := echo.New()
	e.Logger.Fatal(e.Start(conf.Server.Listen))
}

package main

import (
	"BackendZMGestion/internal/db"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	conf, err := initConfig()
	if err != nil {
		panic(err)
	}

	os.Setenv("app-key", conf.App.Key)

	dbConfig := conf.Database

	db, err := db.InitDb(dbConfig.User, dbConfig.Pass, dbConfig.Host, dbConfig.Port, dbConfig.Schema)
	if err != nil {
		panic(err)
	}
	defer db.Conn.Close()
	e := InitRouter(db)

	e.Logger.Fatal(e.Start(conf.Server.Listen))
}

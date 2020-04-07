package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DbHandler struct {
	Conn *sql.DB
}

func InitDb(user string, password string, host string, port string, schema string) (*DbHandler, error) {

	cadena := fmt.Sprintf("%s:%s@tcp(%s:%s)/%", user, password, host, port, schema)

	conn, err := sql.Open("mysql", cadena)
	if err != nil {
		return nil, err
	}
	var output *[]byte
	conn.QueryRow("SHOW Table").Scan(&output)
	fmt.Println(output)

	return &DbHandler{conn}, nil
}

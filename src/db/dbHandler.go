package db

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DbHandler struct {
	Conn *sql.DB
}

// InitDb inicia la conexion con la base de datos
func InitDb(user string, password string, host string, port string, schema string) (*DbHandler, error) {

	cadena := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, schema)

	conn, err := sql.Open("mysql", cadena)
	if err != nil {
		return nil, err
	}

	return &DbHandler{conn}, nil
}

// CallSP realiza la llamada a un sp, retorna un puntero a un slice de bytes
func (h *DbHandler) CallSP(sp string, objeto interface{}) (*[]byte, error) {

	input, err := json.Marshal(objeto)
	if err != nil {
		return nil, err
	}

	var output *[]byte

	startTime := time.Now()

	if objeto != nil {
		query := fmt.Sprintf("CALL %s (?)", sp)

		err = h.Conn.QueryRow(query, string(input)).Scan(&output)
	} else {
		query := fmt.Sprintf("CALL %s ()", sp)
		// env.Logger.Logf("CALL %s ();", sp)
		err = h.Conn.QueryRow(query).Scan(&output)
	}

	called := fmt.Sprintf("CALL %s ('%s'); (%s)", sp, string(input), time.Since(startTime))

	//env.Logger.Log("[mysql] " + called)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		fmt.Println(err)
		return nil, errors.New(called + " " + err.Error())
	}

	return output, nil
}

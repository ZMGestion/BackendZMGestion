package models

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/structs"
	"encoding/json"
	"errors"

	"github.com/mitchellh/mapstructure"
)

type TelasService struct {
	DbHanlder *db.DbHandler
	Tela      *structs.Telas
}

func (ts *TelasService) DarAlta(token string) (*structs.Telas, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"Telas":           ts.Tela,
	}

	out, err := ts.DbHanlder.CallSP("zsp_tela_dar_alta", params)

	if err != nil {
		return nil, err
	}

	if out == nil {
		return nil, errors.New("ERROR_DEFAULT")
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	if err != nil {
		return nil, nil
	}
	var tela structs.Telas
	if response["Telas"] != nil {
		err = mapstructure.Decode(response["Telas"], &tela)
	} else {
		return nil, nil
	}
	if err != nil {
		return nil, nil
	}
	return &tela, nil
}

//Dar baja
func (ts *TelasService) DarBaja(token string) (*structs.Telas, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"Telas":           ts.Tela,
	}

	out, err := ts.DbHanlder.CallSP("zsp_tela_dar_baja", params)

	if err != nil {
		return nil, err
	}

	if out == nil {
		return nil, errors.New("ERROR_DEFAULT")
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	if err != nil {
		return nil, nil
	}
	var tela structs.Telas
	if response["Telas"] != nil {
		err = mapstructure.Decode(response["Telas"], &tela)
	} else {
		return nil, nil
	}
	if err != nil {
		return nil, nil
	}
	return &tela, nil
}

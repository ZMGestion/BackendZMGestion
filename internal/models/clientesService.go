package models

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/structs"
	"encoding/json"
	"errors"

	"github.com/mitchellh/mapstructure"
)

type ClientesService struct {
	DbHanlder *db.DbHandler
	Clientes  *structs.Clientes
}

func (cs *ClientesService) DarAlta(token string) (*structs.Clientes, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"Clientes":        cs.Clientes,
	}

	out, err := cs.DbHanlder.CallSP("zsp_cliente_dar_alta", params)

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
	var clientes structs.Clientes
	if response["Clientes"] != nil {
		err = mapstructure.Decode(response["Clientes"], &clientes)
	} else {
		return nil, nil
	}
	if err != nil {
		return nil, nil
	}
	return &clientes, nil
}

func (cs *ClientesService) DarBaja(token string) (*structs.Clientes, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"Clientes":        cs.Clientes,
	}

	out, err := cs.DbHanlder.CallSP("zsp_cliente_dar_baja", params)

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
	var clientes structs.Clientes
	if response["Clientes"] != nil {
		err = mapstructure.Decode(response["Clientes"], &clientes)
	} else {
		return nil, nil
	}
	if err != nil {
		return nil, nil
	}
	return &clientes, nil
}

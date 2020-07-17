package models

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/helpers"
	"BackendZMGestion/internal/structs"
	"encoding/json"
	"errors"

	"github.com/mitchellh/mapstructure"
)

type ClientesService struct {
	DbHanlder *db.DbHandler
	Cliente   *structs.Clientes
}

func (cs *ClientesService) DarAlta(token string) (*structs.Clientes, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"Clientes":        cs.Cliente,
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
		"Clientes":        cs.Cliente,
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

func (cs *ClientesService) ListarDomicilios() ([]*structs.Domicilios, error) {
	out, err := cs.DbHanlder.CallSP("zsp_cliente_listar_domicilios", helpers.GenerateJSONFromModels(cs.Cliente))

	if out == nil {
		return nil, errors.New("Not found")
	}

	if err != nil {
		return nil, err
	}

	var response []map[string]interface{}

	err = json.Unmarshal(*out, &response)

	if err != nil {
		return nil, err
	}

	var domicilios []*structs.Domicilios
	for _, el := range response {
		var domicilio structs.Domicilios
		if el["Domicilios"] != nil {
			err = mapstructure.Decode(el["Domicilios"], &domicilio)
			if err != nil {
				return nil, err
			}

		} else {
			return nil, nil
		}
		domicilios = append(domicilios, &domicilio)
	}

	return domicilios, nil
}

func (cs *ClientesService) Dame(token string) (*structs.Clientes, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"Clientes":        cs.Cliente,
	}

	out, err := cs.DbHanlder.CallSP("zsp_cliente_dame", params)

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

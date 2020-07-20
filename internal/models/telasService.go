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

//ModificarPrecio
func (ts *TelasService) ModificarPrecio(precio structs.Precios, token string) (interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"Telas":           ts.Tela,
		"Precios":         precio,
	}

	out, err := ts.DbHanlder.CallSP("zsp_tela_modificar_precio", params)

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
	var tela *structs.Telas
	var objetos map[string]interface{}
	if response["Telas"] != nil && response["Precios"] != nil {
		err = mapstructure.Decode(response["Telas"], &tela)
		if err != nil {
			return nil, err
		}
		err = mapstructure.Decode(response["Precios"], &precio)
		if err != nil {
			return nil, err
		}
		objetos = map[string]interface{}{
			"Telas":   tela,
			"Precios": precio,
		}
	}
	return &objetos, nil
}

func (ts *TelasService) ListarPrecios(token string) ([]map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"Telas":           ts.Tela,
	}

	out, err := ts.DbHanlder.CallSP("zsp_tela_listar_precios", params)

	if err != nil {
		return nil, err
	}

	if out == nil {
		return nil, errors.New("ERROR_DEFAULT")
	}

	var response []map[string]interface{}

	err = json.Unmarshal(*out, &response)

	if err != nil {
		return nil, nil
	}

	return response, nil
}

func (ts *TelasService) Dame() (interface{}, error) {

	params := map[string]interface{}{
		"Telas": ts.Tela,
	}

	out, err := ts.DbHanlder.CallSP("zsp_tela_dame", params)

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
	var tela *structs.Telas
	var precio *structs.Precios
	var objetos map[string]interface{}
	if response["Telas"] != nil && response["Precios"] != nil {
		err = mapstructure.Decode(response["Telas"], &tela)
		if err != nil {
			return nil, err
		}
		err = mapstructure.Decode(response["Precios"], &precio)
		if err != nil {
			return nil, err
		}
		objetos = map[string]interface{}{
			"Telas":   tela,
			"Precios": precio,
		}
	}
	return &objetos, nil
}

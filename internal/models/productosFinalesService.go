package models

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/structs"
	"encoding/json"
	"errors"

	"github.com/mitchellh/mapstructure"
)

type ProductosFinalesService struct {
	DbHandler     *db.DbHandler
	ProductoFinal *structs.ProductosFinales
}

func (pfs *ProductosFinalesService) ListarLustres(token string) ([]map[string]interface{}, error) {

	usuariosEjecuta := structs.Usuarios{
		Token: token,
	}

	param := map[string]interface{}{
		"UsuariosEjecuta": usuariosEjecuta,
	}

	out, err := pfs.DbHandler.CallSP("zsp_lustres_listar", param)

	if err != nil {
		return nil, err
	}

	if out == nil {
		return nil, nil
	}

	var response []map[string]interface{}

	err = json.Unmarshal(*out, &response)

	return response, err
}

func (pfs *ProductosFinalesService) Dame(token string) (map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta":  usuarioEjecuta,
		"ProductosFinales": pfs.ProductoFinal,
	}

	out, err := pfs.DbHandler.CallSP("zsp_productoFinal_dame", params)

	if err != nil {
		return nil, err
	}

	if out == nil {
		return nil, errors.New("ERROR_DEFAULT")
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (pfs *ProductosFinalesService) DarAlta(token string) (*structs.ProductosFinales, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta":  usuarioEjecuta,
		"ProductosFinales": pfs.ProductoFinal,
	}

	out, err := pfs.DbHandler.CallSP("zsp_productoFinal_dar_alta", params)

	if err != nil {
		return nil, err
	}

	if out == nil {
		return nil, errors.New("ERROR_DEFAULT")
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	if err != nil {
		return nil, err
	}

	var productoFinal structs.ProductosFinales
	if response["ProductosFinales"] != nil {
		err = mapstructure.Decode(response["ProductosFinales"], &productoFinal)
	} else {
		return nil, nil
	}

	return &productoFinal, err
}

func (pfs *ProductosFinalesService) DarBaja(token string) (*structs.ProductosFinales, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta":  usuarioEjecuta,
		"ProductosFinales": pfs.ProductoFinal,
	}

	out, err := pfs.DbHandler.CallSP("zsp_productoFinal_dar_baja", params)

	if err != nil {
		return nil, err
	}

	if out == nil {
		return nil, errors.New("ERROR_DEFAULT")
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	if err != nil {
		return nil, err
	}

	var productoFinal structs.ProductosFinales
	if response["ProductosFinales"] != nil {
		err = mapstructure.Decode(response["ProductosFinales"], &productoFinal)
	} else {
		return nil, nil
	}

	return &productoFinal, err
}

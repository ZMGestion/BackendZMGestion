package models

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/structs"
	"encoding/json"
	"errors"

	"github.com/mitchellh/mapstructure"
)

type GruposProductoService struct {
	DbHanlder     *db.DbHandler
	GrupoProducto *structs.GruposProducto
}

func (gps *GruposProductoService) DarAlta(token string) (*structs.GruposProducto, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"GruposProducto":  gps.GrupoProducto,
	}

	out, err := gps.DbHanlder.CallSP("zsp_grupoProducto_dar_alta", params)

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
	var grupoProducto structs.GruposProducto
	if response["GruposProducto"] != nil {
		err = mapstructure.Decode(response["GruposProducto"], &grupoProducto)
	} else {
		return nil, nil
	}
	if err != nil {
		return nil, nil
	}
	return &grupoProducto, nil
}

func (gps *GruposProductoService) DarBaja(token string) (*structs.GruposProducto, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"GruposProducto":  gps.GrupoProducto,
	}

	out, err := gps.DbHanlder.CallSP("zsp_grupoProducto_dar_baja", params)

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
	var grupoProducto structs.GruposProducto
	if response["GruposProducto"] != nil {
		err = mapstructure.Decode(response["GruposProducto"], &grupoProducto)
	} else {
		return nil, nil
	}
	if err != nil {
		return nil, nil
	}
	return &grupoProducto, nil
}

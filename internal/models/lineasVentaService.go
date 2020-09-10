package models

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/structs"
	"encoding/json"
	"errors"
)

//LineasVentaService LineasVentaService
type LineasVentaService struct {
	DbHandler      *db.DbHandler
	LineasProducto structs.LineasProducto
}

//Dame Dame
func (lvs *LineasVentaService) Dame(token string) (map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"LineasProducto":  lvs.LineasProducto,
	}

	out, err := lvs.DbHandler.CallSP("zsp_lineaVenta_dame", params)

	if err != nil {
		return nil, err
	}

	if out == nil {
		return nil, errors.New("ERROR_DEFAULT")
	}
	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	return response, err

}

//Cancelar Cancelar
func (lvs *LineasVentaService) Cancelar(token string) (map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"LineasProducto":  lvs.LineasProducto,
	}

	out, err := lvs.DbHandler.CallSP("zsp_lineaVenta_cancelar", params)

	if err != nil {
		return nil, err
	}

	if out == nil {
		return nil, errors.New("ERROR_DEFAULT")
	}
	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	return response, err

}

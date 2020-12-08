package models

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/structs"
	"encoding/json"
)

//ReportesService Servicio de reportes
type ReportesService struct {
	DbHandler *db.DbHandler
}

//Stock Funcion que permite conocer el stock de productos
func (rs *ReportesService) Stock(token string) ([]map[string]interface{}, error) {
	usuariosEjecuta := structs.Usuarios{
		Token: token,
	}

	param := map[string]interface{}{
		"UsuariosEjecuta": usuariosEjecuta,
	}

	out, err := rs.DbHandler.CallSP("zsp_reportes_stock", param)

	if err != nil {
		return nil, err
	}

	var response []map[string]interface{}

	err = json.Unmarshal(*out, &response)

	return response, err
}

//ListaPrecios Funcion que permite generar la lista de precios
func (rs *ReportesService) ListaPrecios(token string) ([]map[string]interface{}, error) {
	usuariosEjecuta := structs.Usuarios{
		Token: token,
	}

	param := map[string]interface{}{
		"UsuariosEjecuta": usuariosEjecuta,
	}

	out, err := rs.DbHandler.CallSP("zsp_reportes_listaPrecios", param)

	if err != nil {
		return nil, err
	}

	var response []map[string]interface{}

	err = json.Unmarshal(*out, &response)

	return response, err
}

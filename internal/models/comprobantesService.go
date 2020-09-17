package models

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/structs"
	"encoding/json"
)

//ComprobantesService ComprobantesService
type ComprobantesService struct {
	DbHandler   *db.DbHandler
	Comprobante *structs.Comprobantes
}

//DarAlta Funcion para dar de alta un comprobante
func (cs *ComprobantesService) DarAlta(token string) (map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Comprobantes":    cs.Comprobante,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	out, err := cs.DbHandler.CallSP("zsp_comprobante_dar_alta", params)

	if err != nil || out == nil {
		return nil, err
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	return response, err
}

//DarBaja Funcion para dar de baja un comprobante
func (cs *ComprobantesService) DarBaja(token string) (map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Comprobantes":    cs.Comprobante,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	out, err := cs.DbHandler.CallSP("zsp_comprobante_dar_baja", params)

	if err != nil || out == nil {
		return nil, err
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	return response, err
}

//Dame Funcion para instanciar un comprobante a partir de su Id
func (cs *ComprobantesService) Dame(token string) (map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Comprobantes":    cs.Comprobante,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	out, err := cs.DbHandler.CallSP("zsp_comprobante_dame", params)

	if err != nil || out == nil {
		return nil, err
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	return response, err
}

package models

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/structs"
	"encoding/json"
)

//RemitosService RemitosService
type RemitosService struct {
	DbHandler *db.DbHandler
	Remitos   structs.Remitos
}

//PasarACreado Funcion para pasar un remito a creado
func (rs *RemitosService) PasarACreado(token string) (map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Remitos":         rs.Remitos,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	out, err := rs.DbHandler.CallSP("zsp_remito_pasar_a_creado", params)

	if err != nil || out == nil {
		return nil, err
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	return response, err
}

//Cancelar Funcion para cancelar un remito
func (rs *RemitosService) Cancelar(token string) (map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Remitos":         rs.Remitos,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	out, err := rs.DbHandler.CallSP("zsp_remito_cancelar", params)

	if err != nil || out == nil {
		return nil, err
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	return response, err
}

//Descancelar Funcion para descancelar un remito
func (rs *RemitosService) Descancelar(token string) (map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Remitos":         rs.Remitos,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	out, err := rs.DbHandler.CallSP("zsp_remito_descancelar", params)

	if err != nil || out == nil {
		return nil, err
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	return response, err
}

//Entregar Funcion para entregar un remito
func (rs *RemitosService) Entregar(token string) (map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Remitos":         rs.Remitos,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	out, err := rs.DbHandler.CallSP("zsp_remito_entregar", params)

	if err != nil || out == nil {
		return nil, err
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	return response, err
}

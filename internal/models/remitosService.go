package models

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/structs"
	"encoding/json"
	"errors"
)

//RemitosService RemitosService
type RemitosService struct {
	DbHandler *db.DbHandler
	Remitos   structs.Remitos
}

//Dame Funcion para pasar instanciar un remito a partir de su Id
func (rs *RemitosService) Dame(token string) (map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Remitos":         rs.Remitos,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	out, err := rs.DbHandler.CallSP("zsp_remito_dame", params)

	if err != nil || out == nil {
		return nil, err
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	return response, err
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

//CrearLineaRemito Funcion qie permite crear una linea de remito
func (rs *RemitosService) CrearLineaRemito(lineaRemito structs.LineasProducto, productoFinal structs.ProductosFinales, token string) (map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta":  usuarioEjecuta,
		"LineasProducto":   lineaRemito,
		"ProductosFinales": productoFinal,
	}

	out, err := rs.DbHandler.CallSP("zsp_lineaRemito_crear", params)

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

//ModificarLineaRemito Funcion qie permite modificar una linea de remito
func (rs *RemitosService) ModificarLineaRemito(lineaRemito structs.LineasProducto, productoFinal structs.ProductosFinales, token string) (map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta":  usuarioEjecuta,
		"LineasProducto":   lineaRemito,
		"ProductosFinales": productoFinal,
	}

	out, err := rs.DbHandler.CallSP("zsp_lineaRemito_modificar", params)

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

//BorrarLineaRemito Funcion que permite borrar una linea de remito
func (rs *RemitosService) BorrarLineaRemito(lineaRemito structs.LineasProducto, token string) error {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"LineasProducto":  lineaRemito,
	}

	_, err := rs.DbHandler.CallSP("zsp_lineaRemito_borrar", params)

	return err
}

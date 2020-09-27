package models

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/structs"
	"encoding/json"
	"errors"
)

//VentasService VentasService
type VentasService struct {
	DbHandler *db.DbHandler
	Ventas    structs.Ventas
}

//Dame Dame
func (vs *VentasService) Dame(token string) (map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"Ventas":          vs.Ventas,
	}

	out, err := vs.DbHandler.CallSP("zsp_venta_dame", params)

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

//CrearLineaVenta CrearLineaVenta
func (vs *VentasService) CrearLineaVenta(lineaProducto structs.LineasProducto, productoFinal structs.ProductosFinales, token string) (map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta":  usuarioEjecuta,
		"LineasProducto":   lineaProducto,
		"ProductosFinales": productoFinal,
	}

	out, err := vs.DbHandler.CallSP("zsp_lineaVenta_crear", params)

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

//ModificarLineaVenta ModificarLineaVenta
func (vs *VentasService) ModificarLineaVenta(lineaProducto structs.LineasProducto, productoFinal structs.ProductosFinales, token string) (map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta":  usuarioEjecuta,
		"LineasProducto":   lineaProducto,
		"ProductosFinales": productoFinal,
	}

	out, err := vs.DbHandler.CallSP("zsp_lineaVenta_modificar", params)

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

//BorrarLineaVenta BorrarLineaVenta
func (vs *VentasService) BorrarLineaVenta(lineaProducto structs.LineasProducto, token string) error {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"LineasProducto":  lineaProducto,
	}

	_, err := vs.DbHandler.CallSP("zsp_lineaVenta_borrar", params)

	return err
}

//ChequearPrecios ChequearPrecios
func (vs *VentasService) ChequearPrecios(token string) (map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"Ventas":          vs.Ventas,
	}

	out, err := vs.DbHandler.CallSP("zsp_venta_chequearPrecios", params)

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

//Revisar Revisar
func (vs *VentasService) Revisar(token string) (map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"Ventas":          vs.Ventas,
	}

	out, err := vs.DbHandler.CallSP("zsp_venta_revisar", params)

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
func (vs *VentasService) Cancelar(token string) (map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"Ventas":          vs.Ventas,
	}

	out, err := vs.DbHandler.CallSP("zsp_venta_cancelar", params)

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

//CrearComprobante Funcion para crear un comprobante
func (vs *VentasService) CrearComprobante(comprobante structs.Comprobantes, token string) (map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Comprobantes":    comprobante,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	out, err := vs.DbHandler.CallSP("zsp_comprobante_crear", params)

	if err != nil || out == nil {
		return nil, err
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	return response, err
}

//ModificarComprobante Funcion para modificar un comprobante
func (vs *VentasService) ModificarComprobante(comprobante structs.Comprobantes, token string) (map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Comprobantes":    comprobante,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	out, err := vs.DbHandler.CallSP("zsp_comprobante_modificar", params)

	if err != nil || out == nil {
		return nil, err
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	return response, err
}

//BuscarComprobante Funcion para buscar un comprobante
func (vs *VentasService) BuscarComprobantes(comprobante structs.Comprobantes, paginacion structs.Paginaciones, token string) (*structs.RespuestaBusqueda, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Comprobantes":    comprobante,
		"UsuariosEjecuta": usuarioEjecuta,
		"Paginaciones":    paginacion,
	}

	out, err := vs.DbHandler.CallSP("zsp_comprobantes_buscar", params)

	if err != nil || out == nil {
		return nil, err
	}

	var respuestaBusqueda structs.RespuestaBusqueda

	err = json.Unmarshal(*out, &respuestaBusqueda)

	if err != nil {
		return nil, err
	}
	return &respuestaBusqueda, nil
}

//BorrarComprobante Funcion para borrar un comprobante
func (vs *VentasService) BorrarComprobante(comprobante structs.Comprobantes, token string) error {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Comprobantes":    comprobante,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	_, err := vs.DbHandler.CallSP("zsp_comprobante_borrar", params)

	return err
}

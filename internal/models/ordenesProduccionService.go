package models

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/structs"
	"encoding/json"
	"errors"
)

//OrdenesProduccionService OrdenesProduccionService
type OrdenesProduccionService struct {
	DbHanlder         *db.DbHandler
	OrdenesProduccion structs.OrdenesProduccion
}

//Dame Dame
func (ops *OrdenesProduccionService) Dame(token string) (map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta":   usuarioEjecuta,
		"OrdenesProduccion": ops.OrdenesProduccion,
	}

	out, err := ops.DbHanlder.CallSP("zsp_ordenProduccion_dame", params)

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

//PasarAPendiente PasarAPendiente
func (ops *OrdenesProduccionService) PasarAPendiente(token string) (map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta":   usuarioEjecuta,
		"OrdenesProduccion": ops.OrdenesProduccion,
	}

	out, err := ops.DbHanlder.CallSP("zsp_ordenProduccion_pasarAPendiente", params)

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

//CrearLineaOrdenProduccion CrearLineaOrdenProduccion
func (ops *OrdenesProduccionService) CrearLineaOrdenProduccion(lineaProducto structs.LineasProducto, productoFinal structs.ProductosFinales, token string) (map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta":  usuarioEjecuta,
		"LineasProducto":   lineaProducto,
		"ProductosFinales": productoFinal,
	}

	out, err := ops.DbHanlder.CallSP("zsp_lineaOrdenProduccion_crear", params)

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

//ModificarLineaOrdenProduccion ModificarLineaOrdenProduccion
func (ops *OrdenesProduccionService) ModificarLineaOrdenProduccion(lineaProducto structs.LineasProducto, productoFinal structs.ProductosFinales, token string) (map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta":  usuarioEjecuta,
		"LineasProducto":   lineaProducto,
		"ProductosFinales": productoFinal,
	}

	out, err := ops.DbHanlder.CallSP("zsp_lineaOrdenProduccion_modificar", params)

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

//BorrarLineaOrdenProduccion BorrarLineaOrdenProduccion
func (ops *OrdenesProduccionService) BorrarLineaOrdenProduccion(lineaProducto structs.LineasProducto, token string) error {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"LineasProducto":  lineaProducto,
	}

	_, err := ops.DbHanlder.CallSP("zsp_lineaOrdenProduccion_borrar", params)

	return err
}

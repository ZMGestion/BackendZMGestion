package models

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/structs"
	"encoding/json"
	"errors"
)

//LineasOrdenProduccionService LineasOrdenProduccionService
type LineasOrdenProduccionService struct {
	DbHandler      *db.DbHandler
	LineasProducto structs.LineasProducto
}

//Dame Dame
func (lps *LineasOrdenProduccionService) Dame(token string) (map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"LineasProducto":  lps.LineasProducto,
	}

	out, err := lps.DbHandler.CallSP("zsp_lineaOrdenProduccion_dame", params)

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

//ListarTareas ListarTareas
func (lps *LineasOrdenProduccionService) ListarTareas(lineaOrdenProduccion structs.LineasProducto, token string) ([]*map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"LineasProducto":  lineaOrdenProduccion,
	}

	out, err := lps.DbHandler.CallSP("zsp_lineaOrdenProduccion_listar_tareas", params)

	if err != nil {
		return nil, err
	}

	if out == nil {
		return nil, errors.New("ERROR_DEFAULT")
	}
	var response []*map[string]interface{}

	err = json.Unmarshal(*out, &response)

	return response, err
}

//CrearTarea CrearTarea
func (lps *LineasOrdenProduccionService) CrearTarea(tarea structs.Tareas, token string) (map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"Tareas":          tarea,
	}

	out, err := lps.DbHandler.CallSP("zsp_tareas_crear", params)

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

//BorrarTarea BorrarTarea
func (lps *LineasOrdenProduccionService) BorrarTarea(tarea structs.Tareas, token string) error {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"Tareas":          tarea,
	}

	_, err := lps.DbHandler.CallSP("zsp_tareas_borrar", params)

	return err
}

//EjecutarTarea EjecutarTarea
func (lps *LineasOrdenProduccionService) EjecutarTarea(tarea structs.Tareas, token string) (map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"Tareas":          tarea,
	}

	out, err := lps.DbHandler.CallSP("zsp_tareas_ejecutar", params)

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

//PausarTarea PausarTarea
func (lps *LineasOrdenProduccionService) PausarTarea(tarea structs.Tareas, token string) (map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"Tareas":          tarea,
	}

	out, err := lps.DbHandler.CallSP("zsp_tareas_pausar", params)

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

//ReanudarTarea ReanudarTarea
func (lps *LineasOrdenProduccionService) ReanudarTarea(tarea structs.Tareas, token string) (map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"Tareas":          tarea,
	}

	out, err := lps.DbHandler.CallSP("zsp_tareas_reanudar", params)

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

//FinalizarTarea FinalizarTarea
func (lps *LineasOrdenProduccionService) FinalizarTarea(tarea structs.Tareas, token string) (map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"Tareas":          tarea,
	}

	out, err := lps.DbHandler.CallSP("zsp_tareas_finalizar", params)

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

//VerificarTarea VerificarTarea
func (lps *LineasOrdenProduccionService) VerificarTarea(tarea structs.Tareas, token string) (map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"Tareas":          tarea,
	}

	out, err := lps.DbHandler.CallSP("zsp_tareas_verificar", params)

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

//CancelarTarea CancelarTarea
func (lps *LineasOrdenProduccionService) CancelarTarea(tarea structs.Tareas, token string) (map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"Tareas":          tarea,
	}

	out, err := lps.DbHandler.CallSP("zsp_tareas_cancelar", params)

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

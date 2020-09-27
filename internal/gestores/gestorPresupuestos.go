package gestores

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/structs"
	"encoding/json"
)

//GestorPresupuestos GestorPresupuestos
type GestorPresupuestos struct {
	DbHandler *db.DbHandler
}

//Crear Crear
func (gp *GestorPresupuestos) Crear(presupuesto structs.Presupuestos, token string) (*map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Presupuestos":    presupuesto,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	out, err := gp.DbHandler.CallSP("zsp_presupuesto_crear", params)

	if err != nil || out == nil {
		return nil, err
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	return &response, err
}

//Buscar Buscar
func (gp *GestorPresupuestos) Buscar(presupuesto structs.Presupuestos, productoFinal structs.ProductosFinales, parametrosBusqueda structs.ParametrosBusqueda, paginaciones structs.Paginaciones, token string) (*structs.RespuestaBusqueda, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta":    usuarioEjecuta,
		"Presupuestos":       presupuesto,
		"ProductosFinales":   productoFinal,
		"ParametrosBusqueda": parametrosBusqueda,
		"Paginaciones":       paginaciones,
	}

	out, err := gp.DbHandler.CallSP("zsp_presupuestos_buscar", params)

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

//Modificar Modificar
func (gp *GestorPresupuestos) Modificar(presupuesto structs.Presupuestos, token string) (*map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Presupuestos":    presupuesto,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	out, err := gp.DbHandler.CallSP("zsp_presupuesto_modificar", params)

	if err != nil || out == nil {
		return nil, err
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	return &response, err
}

//Borrar Borrar
func (gp *GestorPresupuestos) Borrar(presupuesto structs.Presupuestos, token string) error {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Presupuestos":    presupuesto,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	_, err := gp.DbHandler.CallSP("zsp_presupuesto_borrar", params)

	return err
}

//TransformarEnVenta TransformarEnVenta
func (gp *GestorPresupuestos) TransformarEnVenta(venta structs.Ventas, lineasPresupuesto []int, lineasVenta []map[string]interface{}, token string) (*map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Ventas":            venta,
		"LineasPresupuesto": lineasPresupuesto,
		"UsuariosEjecuta":   usuarioEjecuta,
		"LineasVenta":       lineasVenta,
	}

	out, err := gp.DbHandler.CallSP("zsp_presupuestos_transformar_venta", params)

	if err != nil || out == nil {
		return nil, err
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	return &response, err
}

//DamePresupuestos DamePresupuestos
func (gp *GestorPresupuestos) DamePresupuestos(presupuestos []int, token string) ([]map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Presupuestos":    presupuestos,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	out, err := gp.DbHandler.CallSP("zsp_presupuestos_dame_multiple", params)

	if err != nil || out == nil {
		return nil, err
	}

	var response []map[string]interface{}

	err = json.Unmarshal(*out, &response)

	return response, err
}

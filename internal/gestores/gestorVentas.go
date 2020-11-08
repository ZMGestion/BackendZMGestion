package gestores

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/structs"
	"encoding/json"
)

//GestorVentas GestorVentas
type GestorVentas struct {
	DbHandler *db.DbHandler
}

//Crear Funcion para crear ventas
func (gv *GestorVentas) Crear(venta structs.Ventas, token string) (*map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Ventas":          venta,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	out, err := gv.DbHandler.CallSP("zsp_venta_crear", params)

	if err != nil || out == nil {
		return nil, err
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	return &response, err
}

//Modificar Funcion para modificar una venta
func (gv *GestorVentas) Modificar(venta structs.Ventas, token string) (*map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Ventas":          venta,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	out, err := gv.DbHandler.CallSP("zsp_venta_modificar", params)

	if err != nil || out == nil {
		return nil, err
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	return &response, err
}

//Borrar Funcion para borrar una venta
func (gv *GestorVentas) Borrar(venta structs.Ventas, token string) error {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Ventas":          venta,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	_, err := gv.DbHandler.CallSP("zsp_venta_borrar", params)

	return err

}

//Buscar Funcion para buscar una venta
func (gv *GestorVentas) Buscar(venta structs.Ventas, productoFinal structs.ProductosFinales, parametrosBusqueda structs.ParametrosBusqueda, paginaciones structs.Paginaciones, token string) (*structs.RespuestaBusqueda, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta":    usuarioEjecuta,
		"Ventas":             venta,
		"ProductosFinales":   productoFinal,
		"ParametrosBusqueda": parametrosBusqueda,
		"Paginaciones":       paginaciones,
	}

	out, err := gv.DbHandler.CallSP("zsp_ventas_buscar", params)

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

//ModificarDomicilio Funcion para modificar el domicilio de una venta
func (gv *GestorVentas) ModificarDomicilio(venta structs.Ventas, token string) (*map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Ventas":          venta,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	out, err := gv.DbHandler.CallSP("zsp_venta_modificar_domicilio", params)

	if err != nil || out == nil {
		return nil, err
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	return &response, err
}

//DameVentas DameVentas
func (gv *GestorVentas) DameVentas(ventas []int, token string) ([]map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Ventas":          ventas,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	out, err := gv.DbHandler.CallSP("zsp_ventas_dame_multiple", params)

	if err != nil || out == nil {
		return nil, err
	}

	var response []map[string]interface{}

	err = json.Unmarshal(*out, &response)

	return response, err
}

//GenerarOrdenProduccion GenerarOrdenProduccion
func (gv *GestorVentas) GenerarOrdenProduccion(params map[string]interface{}, token string) (*map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}
	params["UsuariosEjecuta"] = usuarioEjecuta

	out, err := gv.DbHandler.CallSP("zsp_venta_generarOrdenProduccion", params)

	if err != nil || out == nil {
		return nil, err
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	return &response, err
}

package gestores

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/structs"
	"encoding/json"
)

type GestorPresupuestos struct {
	DbHandler *db.DbHandler
}

//Crear
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

//Buscar
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

//Modificar
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

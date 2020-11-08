package gestores

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/structs"
	"encoding/json"
)

//GestorOrdenesProduccion GestorOrdenesProduccion
type GestorOrdenesProduccion struct {
	DbHandler *db.DbHandler
}

//Crear Crear
func (gp *GestorOrdenesProduccion) Crear(ordenProduccion structs.OrdenesProduccion, token string) (*map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"OrdenesProduccion": ordenProduccion,
		"UsuariosEjecuta":   usuarioEjecuta,
	}

	out, err := gp.DbHandler.CallSP("zsp_ordenProduccion_crear", params)

	if err != nil || out == nil {
		return nil, err
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	return &response, err
}

//Buscar Buscar
func (gp *GestorOrdenesProduccion) Buscar(params map[string]interface{}, token string) (*structs.RespuestaBusqueda, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}
	params["UsuariosEjecuta"] = usuarioEjecuta

	out, err := gp.DbHandler.CallSP("zsp_ordenesProduccion_buscar", params)

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
func (gp *GestorOrdenesProduccion) Modificar(ordenProduccion structs.OrdenesProduccion, token string) (*map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"OrdenesProduccion": ordenProduccion,
		"UsuariosEjecuta":   usuarioEjecuta,
	}

	out, err := gp.DbHandler.CallSP("zsp_ordenProduccion_modificar", params)

	if err != nil || out == nil {
		return nil, err
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	return &response, err
}

//Borrar Borrar
func (gp *GestorOrdenesProduccion) Borrar(ordenProduccion structs.OrdenesProduccion, token string) error {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"OrdenesProduccion": ordenProduccion,
		"UsuariosEjecuta":   usuarioEjecuta,
	}

	_, err := gp.DbHandler.CallSP("zsp_ordenProduccion_borrar", params)

	return err
}

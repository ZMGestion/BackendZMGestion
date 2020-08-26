package gestores

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/structs"
	"encoding/json"

	"github.com/mitchellh/mapstructure"
)

type GestorGruposProducto struct {
	DbHandler *db.DbHandler
}

//Crear
func (ggp *GestorGruposProducto) Crear(grupoProducto structs.GruposProducto, token string) (*structs.GruposProducto, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"GruposProducto":  grupoProducto,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	out, err := ggp.DbHandler.CallSP("zsp_grupoProducto_crear", params)

	if err != nil || out == nil {
		return nil, err
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	if err == nil {
		if response["GruposProducto"] != nil {
			err = mapstructure.Decode(response["GruposProducto"], &grupoProducto)
		} else {
			return nil, nil
		}
	}

	return &grupoProducto, err
}

func (ggp *GestorGruposProducto) Modificar(grupoProducto structs.GruposProducto, token string) (*structs.GruposProducto, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"GruposProducto":  grupoProducto,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	out, err := ggp.DbHandler.CallSP("zsp_grupoProducto_modificar", params)

	if err != nil || out == nil {
		return nil, err
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	if err == nil {
		if response["GruposProducto"] != nil {
			err = mapstructure.Decode(response["GruposProducto"], &grupoProducto)
		} else {
			return nil, nil
		}
	}

	return &grupoProducto, err
}

func (ggp *GestorGruposProducto) ModificarPrecios(grupoProducto map[string]interface{}, token string) (*map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"GruposProducto":  grupoProducto,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	out, err := ggp.DbHandler.CallSP("zsp_grupoProducto_modificar_precios", params)

	if err != nil || out == nil {
		return nil, err
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	return &response, err
}

func (ggp *GestorGruposProducto) Borrar(grupoProducto structs.GruposProducto, token string) error {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"GruposProducto":  grupoProducto,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	_, err := ggp.DbHandler.CallSP("zsp_grupoProducto_borrar", params)

	return err
}

func (ggp *GestorGruposProducto) Buscar(grupoProducto structs.GruposProducto, paginacion structs.Paginaciones, token string) (*structs.RespuestaBusqueda, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"GruposProducto":  grupoProducto,
		"UsuariosEjecuta": usuarioEjecuta,
		"Paginaciones":    paginacion,
	}

	out, err := ggp.DbHandler.CallSP("zsp_gruposProducto_buscar", params)

	if err != nil {
		return nil, err
	}

	if out == nil {
		return nil, nil
	}

	var respuestaBusqueda structs.RespuestaBusqueda

	err = json.Unmarshal(*out, &respuestaBusqueda)

	if err != nil {
		return nil, err
	}

	return &respuestaBusqueda, nil
}

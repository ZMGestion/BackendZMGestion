package gestores

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/structs"
	"encoding/json"
)

//GestorComprobantes GestorComprobantes
type GestorComprobantes struct {
	DbHanlder *db.DbHandler
}

//Crear Funcion para crear un comprobante
func (gc *GestorComprobantes) Crear(comprobante structs.Comprobantes, token string) (map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Comprobantes":    comprobante,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	out, err := gc.DbHanlder.CallSP("zsp_comprobante_crear", params)

	if err != nil || out == nil {
		return nil, err
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	return response, err
}

//Modificar Funcion para modificar un comprobante
func (gc *GestorComprobantes) Modificar(comprobante structs.Comprobantes, token string) (map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Comprobantes":    comprobante,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	out, err := gc.DbHanlder.CallSP("zsp_comprobante_modificar", params)

	if err != nil || out == nil {
		return nil, err
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	return response, err
}

//Buscar Funcion para buscar un comprobante
func (gc *GestorComprobantes) Buscar(comprobante structs.Comprobantes, token string) (*structs.RespuestaBusqueda, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Comprobantes":    comprobante,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	out, err := gc.DbHanlder.CallSP("zsp_comprobantes_buscar", params)

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

package gestores

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/structs"
	"encoding/json"
)

type GestorPermisos struct {
	DbHandler *db.DbHandler
}

func (gp *GestorPermisos) Listar(paginacion structs.Paginaciones, token string) (*structs.RespuestaBusqueda, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}
	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"Paginaciones":    paginacion,
	}
	out, err := gp.DbHandler.CallSP("zsp_permisos_listar", params)

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
	// var response []map[string]interface{}

	// err = json.Unmarshal(*out, &response)

	// if err != nil {
	// 	return nil, err
	// }

	// var permisos []*structs.Permisos
	// for _, el := range response {
	// 	var permiso structs.Permisos
	// 	if el["Permisos"] != nil {
	// 		err = mapstructure.Decode(el["Permisos"], &permiso)
	// 		if err != nil {
	// 			return nil, err
	// 		}

	// 	} else {
	// 		return nil, nil
	// 	}
	// 	permisos = append(permisos, &permiso)
	// }

	return &respuestaBusqueda, nil

}

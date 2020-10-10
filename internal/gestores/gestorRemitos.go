package gestores

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/structs"
	"encoding/json"
)

//GestorRemitos GestorRemitos
type GestorRemitos struct {
	DbHandler *db.DbHandler
}

//Crear Funcion para crear un nuevo remito
func (gr *GestorRemitos) Crear(remito structs.Remitos, token string) (map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Remitos":         remito,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	out, err := gr.DbHandler.CallSP("zsp_remito_crear", params)

	if err != nil || out == nil {
		return nil, err
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	return response, err
}

//Borrar Funcion para borrar un remito
func (gr *GestorRemitos) Borrar(remito structs.Remitos, token string) error {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Remitos":         remito,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	_, err := gr.DbHandler.CallSP("zsp_remito_borrar", params)

	return err
}

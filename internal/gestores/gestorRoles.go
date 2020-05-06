package gestores

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/structs"
	"encoding/json"

	"github.com/mitchellh/mapstructure"
)

//GestorRoles commented
type GestorRoles struct {
	DbHandler *db.DbHandler
}

//Crear commented
func (gr *GestorRoles) Crear(rol structs.Roles, token string) (*structs.Roles, error) {

	usuario := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Roles":           rol,
		"UsuariosEjecuta": usuario,
	}

	out, err := gr.DbHandler.CallSP("zsp_rol_crear", params)

	if err != nil || out == nil {
		return nil, err
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	if err == nil {
		if response["Roles"] != nil {
			err = mapstructure.Decode(response["Roles"], &rol)
		} else {
			return nil, nil
		}
	}

	return &rol, err
}

//Listar commented
func (gr *GestorRoles) Listar() ([]*structs.Roles, error) {
	roles := []*structs.Roles{}
	out, err := gr.DbHandler.CallSP("zsp_roles_listar", nil)

	if err != nil || out == nil {
		return roles, err
	}

	err = json.Unmarshal(*out, &roles)

	return roles, err
}

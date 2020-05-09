package gestores

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/structs"
	"encoding/json"
	"errors"
	"fmt"

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
	out, err := gr.DbHandler.CallSP("zsp_roles_listar", nil)

	if out == nil {
		return nil, errors.New("Not found")
	}

	if err != nil {
		return nil, err
	}
	fmt.Println(string(*out))
	var response []map[string]interface{}

	err = json.Unmarshal(*out, &response)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	var roles []*structs.Roles
	for _, el := range response {
		var rol structs.Roles
		if el["Roles"] != nil {
			err = mapstructure.Decode(el["Roles"], &rol)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, nil
		}
		roles = append(roles, &rol)
	}

	return roles, nil
}

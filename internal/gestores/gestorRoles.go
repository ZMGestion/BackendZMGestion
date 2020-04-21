package gestores

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/structs"
	"encoding/json"
)

//GestorRoles commented
type GestorRoles struct {
	DbHandler *db.DbHandler
}

//Crear commented
func (gr *GestorRoles) Crear(rol structs.Roles) (*structs.Roles, error) {
	return &structs.Roles{}, nil
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

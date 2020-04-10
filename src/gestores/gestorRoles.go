package gestores

import (
	"BackendZMGestion/src/db"
	"BackendZMGestion/src/models"
	"encoding/json"
)

//GestorRoles commented
type GestorRoles struct {
	DbHandler *db.DbHandler
}

//Crear commented
func (gr *GestorRoles) Crear(rol models.Roles) (*models.Roles, error) {
	return &models.Roles{}, nil
}

//Listar commented
func (gr *GestorRoles) Listar() ([]*models.Roles, error) {
	roles := []*models.Roles{}
	out, err := gr.DbHandler.CallSP("zsp_roles_listar", nil)

	if err != nil || out == nil {
		return roles, err
	}

	err = json.Unmarshal(*out, &roles)

	return roles, err
}

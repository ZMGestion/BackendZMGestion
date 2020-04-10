package models

import (
	"BackendZMGestion/src/db"
	"BackendZMGestion/src/structs"
	"encoding/json"
)

type RolesService struct {
	Rol       *structs.Roles
	DbHandler *db.DbHandler
}

func (s *RolesService) Dame() error {

	out, err := s.DbHandler.CallSP("zsp_rol_dame", s.Rol)

	if err != nil || out == nil {
		return err
	}

	err = json.Unmarshal(*out, &s.Rol)

	return err
}

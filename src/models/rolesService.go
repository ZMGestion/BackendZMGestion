package models

import (
	"BackendZMGestion/src/db"
	"BackendZMGestion/src/helpers"
	"BackendZMGestion/src/structs"
	"encoding/json"
	"errors"
)

type RolesService struct {
	Rol       *structs.Roles
	Error     *structs.Errores
	DbHandler *db.DbHandler
}

func (s *RolesService) Dame() error {
	out, err := s.DbHandler.CallSP("zsp_rol_dame", helpers.GenerateJSONFromModels(s.Rol))

	if out == nil {
		return errors.New("Not found")
	}

	if err != nil {
		s.Rol = nil
		/*
			errores := structs.Errores{}

			errJSON := json.Unmarshal(*out, &errores)
			if errJSON != nil {
				s.Error = &structs.Errores{Error: "Error desconocido."}
			}
			s.Error = &errores
		*/
		return err
	}

	err = json.Unmarshal(*out, &s.Rol)

	return err
}

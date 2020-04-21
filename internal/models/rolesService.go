package models

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/helpers"
	"BackendZMGestion/internal/structs"
	"encoding/json"
	"errors"

	"github.com/mitchellh/mapstructure"
)

type RolesService struct {
	Rol       *structs.Roles
	DbHandler *db.DbHandler
}

func (s *RolesService) Dame() error {
	out, err := s.DbHandler.CallSP("zsp_rol_dame", helpers.GenerateJSONFromModels(s.Rol))

	if out == nil {
		return errors.New("Not found")
	}

	if err != nil {
		s.Rol = nil
		return err
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	if err == nil {
		if response["Roles"] != nil {
			err = mapstructure.Decode(response["Roles"], &s.Rol)
		} else {
			s.Rol = nil
		}
	}

	return err
}

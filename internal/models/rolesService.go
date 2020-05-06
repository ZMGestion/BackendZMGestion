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
	Rol       *structs.Roles //Es lo que se recibe de request
	DbHandler *db.DbHandler
}

func (s *RolesService) Dame() (*structs.Roles, error) {
	out, err := s.DbHandler.CallSP("zsp_rol_dame", helpers.GenerateJSONFromModels(s.Rol))

	if out == nil {
		return nil, errors.New("Not found")
	}

	if err != nil {
		return nil, err
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	var rol structs.Roles
	if err == nil {
		if response["Roles"] != nil {
			err = mapstructure.Decode(response["Roles"], &rol)
		} else {
			return nil, nil
		}
	}

	return &rol, err
}

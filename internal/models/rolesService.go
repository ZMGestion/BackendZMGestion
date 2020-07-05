package models

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/helpers"
	"BackendZMGestion/internal/structs"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type RolesService struct {
	Rol       *structs.Roles //Es lo que se recibe de request
	DbHandler *db.DbHandler
}

func (s *RolesService) Dame() (*structs.Roles, error) {
	out, err := s.DbHandler.CallSP("zsp_rol_dame", helpers.GenerateJSONFromModels(s.Rol))

	if out == nil {
		return nil, errors.New("ERROR_DEFAULT")
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

//ListarPermisos commented
func (s *RolesService) ListarPermisos() ([]*structs.Permisos, error) {
	out, err := s.DbHandler.CallSP("zsp_rol_listar_permisos", helpers.GenerateJSONFromModels(s.Rol))

	if err != nil {
		return nil, err
	}

	if out == nil {
		return nil, nil
	}
	var response []map[string]interface{}

	err = json.Unmarshal(*out, &response)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	var permisos []*structs.Permisos
	for _, el := range response {
		var permiso structs.Permisos
		if el["Permisos"] != nil {
			err = mapstructure.Decode(el["Permisos"], &permiso)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, nil
		}
		permisos = append(permisos, &permiso)
	}

	return permisos, nil
}

//AsignarPermisos commented
func (s *RolesService) AsignarPermisos(pPermisos []structs.Permisos, token string) error {

	usuario := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Roles":           s.Rol,
		"Permisos":        pPermisos,
		"UsuariosEjecuta": usuario,
	}

	_, err := s.DbHandler.CallSP("zsp_rol_asignar_permisos", params)

	return err
}

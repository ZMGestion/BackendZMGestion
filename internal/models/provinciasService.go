package models

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/helpers"
	"BackendZMGestion/internal/structs"
	"encoding/json"

	"github.com/mitchellh/mapstructure"
)

type ProvinciasService struct {
	DbHandler  *db.DbHandler
	Provincias *structs.Provincias
}

func (ps *ProvinciasService) ListarCiudades() ([]*structs.Ciudades, error) {
	out, err := ps.DbHandler.CallSP("zsp_ciudades_listar", helpers.GenerateJSONFromModels(ps.Provincias))

	if err != nil {
		return nil, err
	}

	if out == nil {
		return nil, nil
	}

	var response []map[string]interface{}

	err = json.Unmarshal(*out, &response)

	if err != nil {
		return nil, err
	}

	var ciudades []*structs.Ciudades
	for _, el := range response {
		var ciudad structs.Ciudades
		if el["Ciudades"] != nil {
			err = mapstructure.Decode(el["Ciudades"], &ciudad)
			if err != nil {
				return nil, err
			}

		} else {
			return nil, nil
		}
		ciudades = append(ciudades, &ciudad)
	}

	return ciudades, nil
}

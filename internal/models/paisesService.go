package models

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/helpers"
	"BackendZMGestion/internal/structs"
	"encoding/json"

	"github.com/mitchellh/mapstructure"
)

type PaisesService struct {
	DbHandler *db.DbHandler
	Paises    *structs.Paises
}

func (ps *PaisesService) ListarProvincias() ([]*structs.Provincias, error) {
	out, err := ps.DbHandler.CallSP("zsp_provincias_listar", helpers.GenerateJSONFromModels(ps.Paises))

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

	var provincias []*structs.Provincias
	for _, el := range response {
		var provincia structs.Provincias
		if el["Provincias"] != nil {
			err = mapstructure.Decode(el["Provincias"], &provincia)
			if err != nil {
				return nil, err
			}

		} else {
			return nil, nil
		}
		provincias = append(provincias, &provincia)
	}

	return provincias, nil
}

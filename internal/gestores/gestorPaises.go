package gestores

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/structs"
	"encoding/json"
	"errors"

	"github.com/mitchellh/mapstructure"
)

type GestorPaises struct {
	DbHandler *db.DbHandler
}

func (gp *GestorPaises) Listar() ([]*structs.Paises, error) {
	out, err := gp.DbHandler.CallSP("zsp_paises_listar", nil)

	if out == nil {
		return nil, errors.New("Not found")
	}

	if err != nil {
		return nil, err
	}

	var response []map[string]interface{}

	err = json.Unmarshal(*out, &response)

	if err != nil {
		return nil, err
	}

	var paises []*structs.Paises
	for _, el := range response {
		var pais structs.Paises
		if el["Paises"] != nil {
			err = mapstructure.Decode(el["Paises"], &pais)
			if err != nil {
				return nil, err
			}

		} else {
			return nil, nil
		}
		paises = append(paises, &pais)
	}

	return paises, nil

}

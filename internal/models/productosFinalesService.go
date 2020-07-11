package models

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/structs"
	"encoding/json"

	"github.com/mitchellh/mapstructure"
)

type ProductosFinalesService struct {
	DbHandler     *db.DbHandler
	ProductoFinal *structs.ProductosFinales
}

func (pfs *ProductosFinalesService) ListarLustres(token string) ([]*structs.Lustres, error) {

	usuariosEjecuta := structs.Usuarios{
		Token: token,
	}

	param := map[string]interface{}{
		"UsuariosEjecuta": usuariosEjecuta,
	}

	out, err := pfs.DbHandler.CallSP("zsp_lustres_listar", param)

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

	var lustres []*structs.Lustres
	for _, el := range response {
		if el["Lustres"] != nil {
			var lustre structs.Lustres
			err = mapstructure.Decode(el["Lustres"], &lustre)
			if err != nil {
				return nil, err
			}
			lustres = append(lustres, &lustre)
		} else {
			return nil, nil
		}
	}

	return lustres, nil

}

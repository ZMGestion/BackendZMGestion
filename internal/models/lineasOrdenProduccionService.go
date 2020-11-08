package models

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/structs"
	"encoding/json"
	"errors"
)

//LineasOrdenProduccionService LineasOrdenProduccionService
type LineasOrdenProduccionService struct {
	DbHandler      *db.DbHandler
	LineasProducto structs.LineasProducto
}

//Dame Dame
func (lps *LineasOrdenProduccionService) Dame(token string) (map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"LineasProducto":  lps.LineasProducto,
	}

	out, err := lps.DbHandler.CallSP("zsp_lineaOrdenProduccion_dame", params)

	if err != nil {
		return nil, err
	}

	if out == nil {
		return nil, errors.New("ERROR_DEFAULT")
	}
	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	return response, err

}

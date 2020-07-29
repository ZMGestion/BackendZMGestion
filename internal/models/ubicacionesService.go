package models

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/structs"
	"encoding/json"
	"errors"

	"github.com/mitchellh/mapstructure"
)

type UbicacionesService struct {
	DbHandler   *db.DbHandler
	Ubicaciones *structs.Ubicaciones
}

func (us *UbicacionesService) DarAlta(token string) (*structs.Ubicaciones, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"Ubicaciones":     us.Ubicaciones,
	}

	out, err := us.DbHandler.CallSP("zsp_ubicacion_dar_alta", params)

	if err != nil {
		return nil, err
	}

	if out == nil {
		return nil, errors.New("ERROR_DEFAULT")
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	if err != nil {
		return nil, nil
	}
	var ubicacion structs.Ubicaciones
	if response["Ubicaciones"] != nil {
		err = mapstructure.Decode(response["Ubicaciones"], &ubicacion)
	} else {
		return nil, nil
	}
	if err != nil {
		return nil, nil
	}
	return &ubicacion, nil
}

func (us *UbicacionesService) DarBaja(token string) (*structs.Ubicaciones, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"Ubicaciones":     us.Ubicaciones,
	}

	out, err := us.DbHandler.CallSP("zsp_ubicacion_dar_baja", params)

	if err != nil {
		return nil, err
	}

	if out == nil {
		return nil, errors.New("ERROR_DEFAULT")
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	if err != nil {
		return nil, nil
	}
	var ubicacion structs.Ubicaciones
	if response["Ubicaciones"] != nil {
		err = mapstructure.Decode(response["Ubicaciones"], &ubicacion)
	} else {
		return nil, nil
	}
	if err != nil {
		return nil, nil
	}
	return &ubicacion, nil
}

func (us *UbicacionesService) Dame(token string) (interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"Ubicaciones":     us.Ubicaciones,
	}

	out, err := us.DbHandler.CallSP("zsp_ubicacion_dame", params)

	if err != nil {
		return nil, err
	}

	if out == nil {
		return nil, errors.New("ERROR_DEFAULT")
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	if err != nil {
		return nil, nil
	}
	var ubicacion structs.Ubicaciones
	var domicilio structs.Domicilios

	if response["Ubicaciones"] != nil && response["Domicilios"] != nil {
		err = mapstructure.Decode(response["Ubicaciones"], &ubicacion)
		if err != nil {
			return nil, nil
		}
		err = mapstructure.Decode(response["Domicilios"], &domicilio)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, nil
	}

	respuesta := map[string]interface{}{
		"Ubicaciones": ubicacion,
		"Domicilios":  domicilio,
	}

	return &respuesta, nil
}

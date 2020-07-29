package gestores

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/structs"
	"encoding/json"
	"errors"

	"github.com/mitchellh/mapstructure"
)

type GestorUbicaciones struct {
	DbHandler *db.DbHandler
}

func (gu *GestorUbicaciones) Crear(ubicacion structs.Ubicaciones, domicilio structs.Domicilios, token string) (*structs.Ubicaciones, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Domicilios":      domicilio,
		"Ubicaciones":     ubicacion,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	out, err := gu.DbHandler.CallSP("zsp_ubicacion_crear", params)

	if out == nil {
		return nil, errors.New("Not found")
	}

	if err != nil {
		return nil, err
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	if err == nil {
		if response["Ubicaciones"] != nil {
			err = mapstructure.Decode(response["Ubicaciones"], &ubicacion)
		} else {
			return nil, nil
		}
	}

	return &ubicacion, err
}

func (gu *GestorUbicaciones) Borrar(ubicacion structs.Ubicaciones, token string) error {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Ubicaciones":     ubicacion,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	_, err := gu.DbHandler.CallSP("zsp_ubicacion_borrar", params)

	if err != nil {
		return err
	}

	return nil

}

func (gu *GestorUbicaciones) Modificar(ubicacion structs.Ubicaciones, domicilio structs.Domicilios, token string) (*structs.Ubicaciones, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"Domicilios":      domicilio,
		"Ubicaciones":     ubicacion,
	}

	out, err := gu.DbHandler.CallSP("zsp_ubicacion_modificar", params)

	if out == nil {
		return nil, errors.New("Not found")
	}

	if err != nil {
		return nil, err
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	if err == nil {
		if response["Ubicaciones"] != nil {
			err = mapstructure.Decode(response["Ubicaciones"], &ubicacion)
		} else {
			return nil, nil
		}
	}

	return &ubicacion, err
}

func (gu *GestorUbicaciones) Listar() ([]interface{}, error) {
	out, err := gu.DbHandler.CallSP("zsp_ubicaciones_listar", nil)

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

	var respuesta []interface{}

	for _, el := range response {
		var ubicacion structs.Ubicaciones
		if el["Ubicaciones"] != nil && el["Domicilios"] != nil {
			err = mapstructure.Decode(el["Ubicaciones"], &ubicacion)
			if err != nil {
				return nil, err
			}
			var domicilio structs.Domicilios
			err = mapstructure.Decode(el["Domicilios"], &domicilio)
			if err != nil {
				return nil, err
			}
			var ciudad structs.Ciudades
			if el["Ciudades"] != nil {
				err = mapstructure.Decode(el["Ciudades"], &ciudad)
				if err != nil {
					return nil, err
				}
			}
			var provincia structs.Provincias
			if el["Provincias"] != nil {
				err = mapstructure.Decode(el["Provincias"], &provincia)
				if err != nil {
					return nil, err
				}
			}
			var pais structs.Paises
			if el["Paises"] != nil {
				err = mapstructure.Decode(el["Paises"], &pais)
				if err != nil {
					return nil, err
				}
			}

			objetos := map[string]interface{}{
				"Ubicaciones": ubicacion,
				"Domicilios":  domicilio,
				"Ciudades":    ciudad,
				"Provincias":  provincia,
				"Paises":      pais,
			}
			respuesta = append(respuesta, objetos)
		} else {
			return nil, nil
		}
	}

	return respuesta, nil
}

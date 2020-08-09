package gestores

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/structs"
	"encoding/json"

	"github.com/mitchellh/mapstructure"
)

type GestorTelas struct {
	DbHandler *db.DbHandler
}

func (gt *GestorTelas) Crear(tela structs.Telas, precio structs.Precios, token string) (*structs.Telas, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Telas":           tela,
		"Precios":         precio,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	out, err := gt.DbHandler.CallSP("zsp_tela_crear", params)

	if err != nil || out == nil {
		return nil, err
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	if err == nil {
		if response["Telas"] != nil {
			err = mapstructure.Decode(response["Telas"], &tela)
		} else {
			return nil, nil
		}
	}

	return &tela, err
}

func (gt *GestorTelas) Modificar(tela structs.Telas, precio structs.Precios, token string) (*structs.Telas, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Telas":           tela,
		"Precios":         precio,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	out, err := gt.DbHandler.CallSP("zsp_tela_modificar", params)

	if err != nil || out == nil {
		return nil, err
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	if err == nil {
		if response["Telas"] != nil {
			err = mapstructure.Decode(response["Telas"], &tela)
		} else {
			return nil, nil
		}
	}

	return &tela, err
}

//Borrar Tela
func (gt *GestorTelas) Borrar(telas structs.Telas, token string) error {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Telas":           telas,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	_, err := gt.DbHandler.CallSP("zsp_tela_borrar", params)

	if err != nil {
		return err
	}
	return nil

}

func (gt *GestorTelas) Buscar(tela structs.Telas, paginacion structs.Paginaciones, token string) (*structs.RespuestaBusqueda, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Telas":           tela,
		"UsuariosEjecuta": usuarioEjecuta,
		"Paginaciones":    paginacion,
	}

	out, err := gt.DbHandler.CallSP("zsp_telas_buscar", params)

	if err != nil {
		return nil, err
	}

	if out == nil {
		return nil, nil
	}
	var respuestaBusqueda structs.RespuestaBusqueda

	err = json.Unmarshal(*out, &respuestaBusqueda)

	if err != nil {
		return nil, err

	}
	return &respuestaBusqueda, nil
	// var response []map[string]interface{}

	// err = json.Unmarshal(*out, &response)
	// var respuesta []interface{}
	// if err == nil {
	// 	for _, el := range response {
	// 		if el["Telas"] != nil && el["Precios"] != nil {
	// 			var tela structs.Telas
	// 			err = mapstructure.Decode(el["Telas"], &tela)
	// 			if err != nil {
	// 				return nil, err
	// 			}
	// 			var precio structs.Precios
	// 			err = mapstructure.Decode(el["Precios"], &precio)
	// 			if err != nil {
	// 				return nil, err
	// 			}
	// 			objeto := map[string]interface{}{
	// 				"Telas":   tela,
	// 				"Precios": precio,
	// 			}
	// 			respuesta = append(respuesta, objeto)
	// 		}
	// 	}
	// } else {
	// 	return nil, err
	// }

	// return respuesta, nil

}

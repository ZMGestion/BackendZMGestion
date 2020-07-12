package gestores

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/structs"
	"encoding/json"

	"github.com/mitchellh/mapstructure"
)

type GestorProductos struct {
	DbHandler *db.DbHandler
}

func (gp *GestorProductos) Crear(producto structs.Productos, precio structs.Precios, token string) (*structs.Productos, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Productos":       producto,
		"Precios":         precio,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	out, err := gp.DbHandler.CallSP("zsp_producto_crear", params)

	if err != nil || out == nil {
		return nil, err
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	if err == nil {
		if response["Productos"] != nil {
			err = mapstructure.Decode(response["Productos"], &producto)
		} else {
			return nil, nil
		}
	}

	return &producto, err
}

func (gp *GestorProductos) Modificar(producto structs.Productos, token string) (*structs.Productos, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Productos":       producto,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	out, err := gp.DbHandler.CallSP("zsp_producto_modificar", params)

	if err != nil || out == nil {
		return nil, err
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	if err == nil {
		if response["Productos"] != nil {
			err = mapstructure.Decode(response["Productos"], &producto)
		} else {
			return nil, nil
		}
	}

	return &producto, err
}

func (gp *GestorProductos) Borrar(producto structs.Productos, token string) (*structs.Productos, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Productos":       producto,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	_, err := gp.DbHandler.CallSP("zsp_producto_borrar", params)

	if err != nil {
		return nil, err
	}

	return nil, err
}

func (gp *GestorProductos) Buscar(producto structs.Productos, paginacion structs.Paginaciones, token string) (*structs.RespuestaBusqueda, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Productos":       producto,
		"UsuariosEjecuta": usuarioEjecuta,
		"Paginaciones":    paginacion,
	}

	out, err := gp.DbHandler.CallSP("zsp_productos_buscar", params)

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

	// var response []map[string]interface{}

	// err = json.Unmarshal(*out, &response)
	// var respuesta []interface{}
	// if err == nil {
	// 	for _, el := range response {
	// 		if el["Productos"] != nil && el["Precios"] != nil {
	// 			var producto structs.Productos
	// 			err = mapstructure.Decode(el["Productos"], &producto)
	// 			if err != nil {
	// 				return nil, err
	// 			}
	// 			var precio structs.Precios
	// 			err = mapstructure.Decode(el["Precios"], &precio)
	// 			if err != nil {
	// 				return nil, err
	// 			}
	// 			objeto := map[string]interface{}{
	// 				"Productos": producto,
	// 				"Precios":   precio,
	// 			}
	// 			respuesta = append(respuesta, objeto)
	// 		}
	// 	}
	// } else {
	// 	return nil, err
	// }

	return &respuestaBusqueda, nil
}

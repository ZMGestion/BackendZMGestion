package models

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/structs"
	"encoding/json"
	"errors"

	"github.com/mitchellh/mapstructure"
)

type ProductosService struct {
	DbHandler *db.DbHandler
	Producto  *structs.Productos
}

func (ps *ProductosService) DarAlta(token string) (*structs.Productos, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"Productos":       ps.Producto,
	}

	out, err := ps.DbHandler.CallSP("zsp_producto_dar_alta", params)

	if err != nil {
		return nil, err
	}

	if out == nil {
		return nil, errors.New("ERROR_DEFAULT")
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	if err != nil {
		return nil, err
	}
	var producto structs.Productos
	if response["Productos"] != nil {
		err = mapstructure.Decode(response["Productos"], &producto)
	} else {
		return nil, nil
	}
	if err != nil {
		return nil, nil
	}
	return &producto, nil
}

func (ps *ProductosService) DarBaja(token string) (*structs.Productos, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"Productos":       ps.Producto,
	}

	out, err := ps.DbHandler.CallSP("zsp_producto_dar_baja", params)

	if err != nil {
		return nil, err
	}

	if out == nil {
		return nil, errors.New("ERROR_DEFAULT")
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	if err != nil {
		return nil, err
	}
	var producto structs.Productos
	if response["Productos"] != nil {
		err = mapstructure.Decode(response["Productos"], &producto)
	} else {
		return nil, nil
	}
	if err != nil {
		return nil, nil
	}
	return &producto, nil
}

func (ps *ProductosService) ModificarPrecio(precio structs.Precios, token string) (interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"Productos":       ps.Producto,
		"Precios":         precio,
	}

	out, err := ps.DbHandler.CallSP("zsp_producto_modificar_precio", params)

	if err != nil {
		return nil, err
	}

	if out == nil {
		return nil, errors.New("ERROR_DEFAULT")
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	if err != nil {
		return nil, err
	}
	return response, nil
}

func (ps *ProductosService) ListarPrecios(token string) ([]*structs.Precios, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"Productos":       ps.Producto,
	}

	out, err := ps.DbHandler.CallSP("zsp_producto_listar_precios", params)

	if err != nil {
		return nil, err
	}

	if out == nil {
		return nil, errors.New("ERROR_DEFAULT")
	}
	var response []map[string]interface{}

	err = json.Unmarshal(*out, &response)

	if err != nil {
		return nil, err
	}
	var precios []*structs.Precios
	for _, el := range response {
		if el["Precios"] != nil {
			var precio structs.Precios
			err = mapstructure.Decode(el["Precios"], &precio)
			if err != nil {
				return nil, err
			}
			precios = append(precios, &precio)
		}
	}

	return precios, nil

}

func (ps *ProductosService) Dame(token string) (interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"Productos":       ps.Producto,
	}

	out, err := ps.DbHandler.CallSP("zsp_producto_dame", params)

	if err != nil {
		return nil, err
	}

	if out == nil {
		return nil, errors.New("ERROR_DEFAULT")
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)
	if err != nil {
		return nil, err
	}

	var producto structs.Productos
	var precio structs.Precios
	if response["Productos"] != nil && response["Precios"] != nil {
		err = mapstructure.Decode(response["Productos"], &producto)
		if err != nil {
			return nil, err
		}
		err = mapstructure.Decode(response["Precios"], &precio)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("ERROR_DEFAULT")
	}

	respuesta := map[string]interface{}{
		"Productos": producto,
		"Precios":   precio,
	}

	return respuesta, nil
}

func (ps *ProductosService) ListarCategorias(token string) ([]*structs.CategoriasProducto, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
	}

	out, err := ps.DbHandler.CallSP("zsp_categoriasProducto_listar", params)

	if err != nil {
		return nil, err
	}

	if out == nil {
		return nil, errors.New("ERROR_DEFAULT")
	}
	var response []map[string]interface{}

	err = json.Unmarshal(*out, &response)

	if err != nil {
		return nil, err
	}
	var categorias []*structs.CategoriasProducto
	for _, el := range response {
		if el["CategoriasProducto"] != nil {
			var categoria structs.CategoriasProducto
			err = mapstructure.Decode(el["CategoriasProducto"], &categoria)
			if err != nil {
				return nil, err
			}
			categorias = append(categorias, &categoria)
		}
	}
	return categorias, nil
}

func (ps *ProductosService) ListarTiposProducto(token string) ([]*structs.TiposProducto, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
	}

	out, err := ps.DbHandler.CallSP("zsp_tiposProducto_listar", params)

	if err != nil {
		return nil, err
	}

	if out == nil {
		return nil, errors.New("ERROR_DEFAULT")
	}
	var response []map[string]interface{}

	err = json.Unmarshal(*out, &response)

	if err != nil {
		return nil, err
	}
	var tiposProducto []*structs.TiposProducto
	for _, el := range response {
		if el["TiposProducto"] != nil {
			var tipoProducto structs.TiposProducto
			err = mapstructure.Decode(el["TiposProducto"], &tipoProducto)
			if err != nil {
				return nil, err
			}
			tiposProducto = append(tiposProducto, &tipoProducto)
		}
	}
	return tiposProducto, nil
}

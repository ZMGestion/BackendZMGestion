package gestores

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/structs"
	"encoding/json"
)

type GestorProductosFinales struct {
	DbHandler *db.DbHandler
}

func (gpf *GestorProductosFinales) Crear(productoFinal structs.ProductosFinales, token string) (*map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"ProductosFinales": productoFinal,
		"UsuariosEjecuta":  usuarioEjecuta,
	}

	out, err := gpf.DbHandler.CallSP("zsp_productoFinal_crear", params)

	if err != nil || out == nil {
		return nil, err
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	return &response, err
}

func (gpf *GestorProductosFinales) Modificar(productoFinal structs.ProductosFinales, token string) (*map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"ProductosFinales": productoFinal,
		"UsuariosEjecuta":  usuarioEjecuta,
	}

	out, err := gpf.DbHandler.CallSP("zsp_productoFinal_modificar", params)

	if err != nil || out == nil {
		return nil, err
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	return &response, err
}

func (gpf *GestorProductosFinales) Borrar(productoFinal structs.ProductosFinales, token string) error {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"ProductosFinales": productoFinal,
		"UsuariosEjecuta":  usuarioEjecuta,
	}

	_, err := gpf.DbHandler.CallSP("zsp_productoFinal_borrar", params)

	return err
}

func (gpf *GestorProductosFinales) Buscar(productoFinal structs.ProductosFinales, paginacion structs.Paginaciones, token string) (*structs.RespuestaBusqueda, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"ProductosFinales": productoFinal,
		"UsuariosEjecuta":  usuarioEjecuta,
		"Paginaciones":     paginacion,
	}

	out, err := gpf.DbHandler.CallSP("zsp_productosFinales_buscar", params)

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
}

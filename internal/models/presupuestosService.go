package models

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/structs"
	"encoding/json"
	"errors"
)

type PresupuestosService struct {
	DbHanlder    *db.DbHandler
	Presupuestos structs.Presupuestos
}

func (ps *PresupuestosService) PasarACreado(token string) (map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"Presupuestos":    ps.Presupuestos,
	}

	out, err := ps.DbHanlder.CallSP("zsp_presupuesto_pasar_a_creado", params)

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

func (ps *PresupuestosService) CrearLineaPresupuesto(lineaProducto structs.LineasProducto, productoFinal structs.ProductosFinales, token string) (map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta":  usuarioEjecuta,
		"LineasProducto":   lineaProducto,
		"ProductosFinales": productoFinal,
	}

	out, err := ps.DbHanlder.CallSP("zsp_lineaPresupuesto_crear", params)

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

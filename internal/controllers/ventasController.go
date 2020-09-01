package controllers

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/gestores"
	"BackendZMGestion/internal/helpers"
	"BackendZMGestion/internal/interfaces"
	"BackendZMGestion/internal/structs"
	"net/http"

	"github.com/labstack/echo"
	"github.com/mitchellh/mapstructure"
)

//VentasController VentasController
type VentasController struct {
	DbHandler *db.DbHandler
}

/**
 * @api {POST} /ventas/crear Crear Venta
 * @apiDescription Permite crear una venta
 * @apiGroup Ventas
 * @apiHeader {String} Authorization
 * @apiParam {Object} Ventas
 * @apiParam {int} Ventas.IdCliente
 * @apiParam {int} [Ventas.IdDomicilio]
 * @apiParam {int} Presupuestos.IdUbicacion
 * @apiParam {string} [Presupuestos.Observaciones]

  * @apiParamExample {json} Request-Example:
{
  "Ventas":{
    "IdCliente":3,
    "IdDomicilio":14,
    "IdUbicacion":1,
    "Observaciones":
  }
}
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":{
		"Ventas":{
			"Estado": "E",
			"FechaAlta": "2020-09-01 22:48:43.000000",
			"IdCliente": 3,
			"IdDomicilio": 14,
			"IdUbicacion": 1,
			"IdUsuario": 1,
			"IdVenta": 1,
			"Observaciones": null
		}
	}
}
* @apiErrorExample {json} Error-Response:
{
    "error": {
        "codigo": "ERROR_DEFAULT",
        "mensaje": "Ha ocurrido un error mientras se procesaba su petición."
    },
    "respuesta": null
}
*/
//Crear Crear
func (vc *VentasController) Crear(c echo.Context) error {

	venta := structs.Ventas{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Ventas"], &venta)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorVentas := gestores.GestorVentas{
		DbHandler: vc.DbHandler,
	}
	result, err := gestorVentas.Crear(venta, *token)

	if err != nil || result == nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error:     nil,
		Respuesta: result,
	}

	return c.JSON(http.StatusOK, response)
}

/**
 * @api {POST} /ventas/modificar Modificar Venta
 * @apiDescription Permite modificar una venta
 * @apiGroup Ventas
 * @apiHeader {String} Authorization
 * @apiParam {Object} Ventas
 * @apiParam {int} Ventas.IdVenta
 * @apiParam {int} Ventas.IdCliente
 * @apiParam {int} [Ventas.IdDomicilio]
 * @apiParam {int} Presupuestos.IdUbicacion
 * @apiParam {string} [Presupuestos.Observaciones]

  * @apiParamExample {json} Request-Example:
{
  "Ventas":{
    "IdVenta":1,
    "IdCliente":3,
    "IdDomicilio":14,
    "IdUbicacion":2,
    "Observaciones":""
  }
}
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":{
		"Ventas":{
			"Estado": "E",
			"FechaAlta": "2020-09-01 22:48:43.000000",
			"IdCliente": 3,
			"IdDomicilio": 14,
			"IdUbicacion": 2,
			"IdUsuario": 1,
			"IdVenta": 1,
			"Observaciones": null
		}
	}
}
* @apiErrorExample {json} Error-Response:
{
    "error": {
        "codigo": "ERROR_DEFAULT",
        "mensaje": "Ha ocurrido un error mientras se procesaba su petición."
    },
    "respuesta": null
}
*/
//Modificar Modificar
func (vc *VentasController) Modificar(c echo.Context) error {

	venta := structs.Ventas{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Ventas"], &venta)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorVentas := gestores.GestorVentas{
		DbHandler: vc.DbHandler,
	}
	result, err := gestorVentas.Modificar(venta, *token)

	if err != nil || result == nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error:     nil,
		Respuesta: result,
	}

	return c.JSON(http.StatusOK, response)
}

/**
 * @api {POST} /ventas/borrar Borrar Venta
 * @apiDescription Permite borrar una venta
 * @apiGroup Ventas
 * @apiHeader {String} Authorization
 * @apiParam {Object} Ventas
 * @apiParam {int} Ventas.IdVenta

 * @apiParamExample {json} Request-Example:
{
  "Ventas":{
    "IdVenta":1
  }
}
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta": null
}
* @apiErrorExample {json} Error-Response:
{
    "error": {
        "codigo": "ERROR_DEFAULT",
        "mensaje": "Ha ocurrido un error mientras se procesaba su petición."
    },
    "respuesta": null
}
*/
//Borrar Borrar
func (vc *VentasController) Borrar(c echo.Context) error {

	venta := structs.Ventas{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Ventas"], &venta)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorVentas := gestores.GestorVentas{
		DbHandler: vc.DbHandler,
	}
	err = gestorVentas.Borrar(venta, *token)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error:     nil,
		Respuesta: nil,
	}

	return c.JSON(http.StatusOK, response)
}

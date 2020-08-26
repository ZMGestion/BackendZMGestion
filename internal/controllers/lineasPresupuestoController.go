package controllers

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/helpers"
	"BackendZMGestion/internal/interfaces"
	"BackendZMGestion/internal/models"
	"BackendZMGestion/internal/structs"
	"net/http"

	"github.com/labstack/echo"
	"github.com/mitchellh/mapstructure"
)

type LineasPresupuestoController struct {
	DbHanlder *db.DbHandler
}

/**
 * @api {POST} /lineasPresupuestos/crear Crear linea de presupuesto
 * @apiDescription Permite crear una linea de presupuesto
 * @apiGroup Presupuestos
 * @apiHeader {String} Authorization
 * @apiParam {Object} LineasProducto
 * @apiParam {int} LineasProducto.IdPresupuesto
 * @apiParam {float} LineasProducto.PrecioUnitario
 * @apiParam {int} LineasProducto.Cantidad
 * @apiParam {Object} ProductosFinales
 * @apiParam {int} ProductosFinales.IdProducto
 * @apiParam {int} ProductosFinales.IdTela
 * @apiParam {int} ProductosFinales.IdLuste

 * @apiParamExample {json} Request-Example:
{
  "LineasProducto":{
    "IdReferencia":2,
    "Cantidad":3,
    "PrecioUnitario":108.21
  },
  "ProductosFinales":{
    "IdProducto":15,
    "IdTela":28,
    "IdLustre":2
  }
}
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":{
		"LineasProducto":{
			"Cantidad": 3,
			"Estado": "P",
			"FechaAlta": "2020-08-24 00:27:57.000000",
			"FechaCancelacion": null,
			"IdLineaProducto": 3,
			"IdLineaProductoPadre": null,
			"IdProductoFinal": 30,
			"IdReferencia": 2,
			"IdUbicacion": null,
			"PrecioUnitario": 1000,
			"Tipo": "P"
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
//Crear
func (lpc *LineasPresupuestoController) Crear(c echo.Context) error {

	lineaPresupuesto := structs.LineasProducto{}
	productoFinal := structs.ProductosFinales{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["LineasProducto"], &lineaPresupuesto)
	mapstructure.Decode(jsonMap["ProductosFinales"], &productoFinal)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnauthorized)
	}

	presupuestosService := models.PresupuestosService{
		DbHanlder: lpc.DbHanlder,
	}
	result, err := presupuestosService.CrearLineaPresupuesto(lineaPresupuesto, productoFinal, *token)

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
 * @api {POST} /lineasPresupuestos Listar lineas de presupuesto
 * @apiDescription Permite listar las lineas de presupuesto de un presupuesto dado
 * @apiGroup Presupuestos
 * @apiHeader {String} Authorization
 * @apiParam {Object} Presupuestos
 * @apiParam {int} Presupuestos.IdPresupuesto
 * @apiParamExample {json} Request-Example:
{
    "Presupuestos": {
		"IdPresupuestos": 1
	}
}
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":{
		"LineasProducto":[
			{
				"Cantidad": 4,
				"Estado": "P",
				"FechaAlta": "2020-08-17 20:13:28.000000",
				"FechaCancelacion": null,
				"IdLineaProducto": 1,
				"IdLineaProductoPadre": null,
				"IdProductoFinal": 1,
				"IdReferencia": 1,
				"IdUbicacion": null,
				"PrecioUnitario": 39000,
				"Tipo": "P"
			},
			{
				"Cantidad": 4,
				"Estado": "P",
				"FechaAlta": "2020-08-17 23:56:58.000000",
				"FechaCancelacion": null,
				"IdLineaProducto": 2,
				"IdLineaProductoPadre": null,
				"IdProductoFinal": 16,
				"IdReferencia": 1,
				"IdUbicacion": null,
				"PrecioUnitario": 39000,
				"Tipo": "P"
			}
		],
		"Presupuestos":{
			"Estado": "C",
			"FechaAlta": "2020-08-17 18:51:47.000000",
			"IdCliente": 3,
			"IdPresupuesto": 1,
			"IdUbicacion": 1,
			"IdUsuario": 1,
			"IdVenta": null,
			"Observaciones": null,
			"PeriodoValidez": 15
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
//Crear
func (lpc *LineasPresupuestoController) Listar(c echo.Context) error {

	presupuestos := structs.Presupuestos{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Presupuestos"], &presupuestos)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnauthorized)
	}

	presupuestosService := models.PresupuestosService{
		DbHanlder:    lpc.DbHanlder,
		Presupuestos: presupuestos,
	}
	result, err := presupuestosService.ListarLineasPresupuesto(*token)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error:     nil,
		Respuesta: result,
	}

	return c.JSON(http.StatusOK, response)
}

/**
 * @api {POST} /lineasPresupuestos/modificar Modificar linea de presupuesto
 * @apiDescription Permite modificar una linea de presupuesto
 * @apiGroup Presupuestos
 * @apiHeader {String} Authorization
 * @apiParam {Object} LineasProducto
 * @apiParam {int} LineasProducto.IdLineaProducto
 * @apiParam {int} LineasProducto.IdPresupuesto
 * @apiParam {float} LineasProducto.PrecioUnitario
 * @apiParam {int} LineasProducto.Cantidad
 * @apiParam {Object} ProductosFinales
 * @apiParam {int} ProductosFinales.IdProducto
 * @apiParam {int} ProductosFinales.IdTela
 * @apiParam {int} ProductosFinales.IdLuste

  * @apiParamExample {json} Request-Example:
{
  "LineasProducto":{
	"IdLineaProducto":3
    "IdReferencia":2,
    "Cantidad":3,
    "PrecioUnitario":108.21
  },
  "ProductosFinales":{
    "IdProducto":15,
    "IdTela":28,
    "IdLustre":1
  }
}

 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":{
		"LineasProducto":{
			"Cantidad": 3,
			"Estado": "P",
			"FechaAlta": "2020-08-24 00:27:57.000000",
			"FechaCancelacion": null,
			"IdLineaProducto": 3,
			"IdLineaProductoPadre": null,
			"IdProductoFinal": 30,
			"IdReferencia": 2,
			"IdUbicacion": null,
			"PrecioUnitario": 1000,
			"Tipo": "P"
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
func (lpc *LineasPresupuestoController) Modificar(c echo.Context) error {

	lineaPresupuesto := structs.LineasProducto{}
	productoFinal := structs.ProductosFinales{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["LineasProducto"], &lineaPresupuesto)
	mapstructure.Decode(jsonMap["ProductosFinales"], &productoFinal)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnauthorized)
	}

	presupuestosService := models.PresupuestosService{
		DbHanlder: lpc.DbHanlder,
	}
	result, err := presupuestosService.ModificarLineaPresupuesto(lineaPresupuesto, productoFinal, *token)

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
 * @api {POST} /lineasPresupuestos/borrar Borrar linea de presupuesto
 * @apiDescription Permite borrar una linea de presupuesto
 * @apiGroup Presupuestos
 * @apiHeader {String} Authorization
 * @apiParam {Object} LineasProducto
 * @apiParam {int} LineasProducto.IdLineaProducto
 * @apiParamExample {json} Request-Example:
{
  "LineasProducto":{
	"IdLineaProducto":3
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
func (lpc *LineasPresupuestoController) Borrar(c echo.Context) error {

	lineaPresupuesto := structs.LineasProducto{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["LineasProducto"], &lineaPresupuesto)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnauthorized)
	}

	presupuestosService := models.PresupuestosService{
		DbHanlder: lpc.DbHanlder,
	}
	err = presupuestosService.BorrarLineaPresupuesto(lineaPresupuesto, *token)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error:     nil,
		Respuesta: nil,
	}

	return c.JSON(http.StatusOK, response)
}

/**
 * @api {POST} /lineasPresupuestos/dame Dame linea de presupuesto
 * @apiDescription Permite instanciar una linea de presupuesto a partir de su Id
 * @apiGroup Presupuestos
 * @apiHeader {String} Authorization
 * @apiParam {Object} LineasProducto
 * @apiParam {int} LineasProducto.IdLineaProducto
 * @apiParamExample {json} Request-Example:
{
  "LineasProducto":{
	"IdLineaProducto":3
  }
}

 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":{
		"LineasProducto":{
			"Cantidad": 4,
			"IdLineaProducto": 1,
			"IdProductoFinal": 1,
			"PrecioUnitario": 39000
		},
		"Lustres": null,
			"Productos":{
			"IdProducto": 5,
			"Producto": "Silla comedor"
		},
		"ProductosFinales":{
			"FechaAlta": "2020-07-10 15:26:14.000000",
			"IdLustre": null,
			"IdProducto": 5,
			"IdProductoFinal": 1,
			"IdTela": 5
		},
		"Telas":{
			"IdTela": 5,
			"Tela": "Cuero de cebra"
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
//Dame Dame
func (lpc *LineasPresupuestoController) Dame(c echo.Context) error {

	lineaPresupuesto := structs.LineasProducto{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["LineasProducto"], &lineaPresupuesto)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnauthorized)
	}

	lineasPresupuestoService := models.LineasPresupuestoService{
		DbHandler:      lpc.DbHanlder,
		LineasProducto: lineaPresupuesto,
	}
	result, err := lineasPresupuestoService.Dame(*token)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error:     nil,
		Respuesta: result,
	}

	return c.JSON(http.StatusOK, response)
}

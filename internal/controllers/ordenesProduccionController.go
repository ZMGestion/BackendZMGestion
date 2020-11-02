package controllers

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/gestores"
	"BackendZMGestion/internal/helpers"
	"BackendZMGestion/internal/interfaces"
	"BackendZMGestion/internal/models"
	"BackendZMGestion/internal/structs"
	"net/http"

	"github.com/labstack/echo"
	"github.com/mitchellh/mapstructure"
)

//OrdenesProduccionController OrdenesProduccionController
type OrdenesProduccionController struct {
	DbHanlder *db.DbHandler
}

/**
 * @api {POST} /ordenesProduccion Buscar Órdenes de Producción
 * @apiDescription Permite buscar una orden de producción
 * @apiGroup OrdenesProduccion
 * @apiHeader {String} Authorization
 * @apiParam {Object} [OrdenesProduccion]
 * @apiParam {string} [OrdenesProduccion.Estado]
 * @apiParam {Object} [Tareas]
 * @apiParam {int} [Tareas.IdUsuarioFabricante]
 * @apiParam {int} [Tareas.IdUsuarioRevisor]
 * @apiParam {Object} [ProductosFinales]
 * @apiParam {int} [ProductosFinales.IdProducto]
 * @apiParam {int} [ProductosFinales.IdTela]
 * @apiParam {int} [ProductosFinales.IdLustre]
 * @apiParam {object} [ParametrosBusqueda]
 * @apiParam {string} [ParametrosBusqueda.FechaInicio]
 * @apiParam {string} [ParametrosBusqueda.FechaFin]
 * @apiParam {Object} [Paginaciones]
 * @apiParam {int} [Paginaciones.Pagina]
 * @apiParam {int} [Paginaciones.LongitudPagina]

  * @apiParamExample {json} Request-Example:
{
    "OrdenesProduccion": {
		"Estado": "C"
	},
	"Tareas":{
		"IdUsuarioFabricante":1,
		"IdUsuarioRevisor":0
	},
	"ProductosFinales":{
		"IdProducto":1,
		"IdTela":0,
		"IdLustre":0
	},
	"Paginaciones":{
		"Pagina":1,
		"LongitudPagina":10
	}
}
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":{
		"resultado":[

		],
		"Paginaciones":{
		"Pagina": 1,
		"LongitudPagina": 10,
		"CantidadTotal": 1
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
//Buscar Buscar
func (pc *OrdenesProduccionController) Buscar(c echo.Context) error {

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorOrdenesProduccion := gestores.GestorOrdenesProduccion{
		DbHandler: pc.DbHanlder,
	}
	result, err := gestorOrdenesProduccion.Buscar(jsonMap, *token)

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
 * @api {POST} /ordenesProduccion/crear Crear Orden Producción
 * @apiDescription Permite crear una orden de producción
 * @apiGroup OrdenesProduccion
 * @apiHeader {String} Authorization
 * @apiParam {Object} OrdenesProduccion
 * @apiParam {int} [OrdenesProduccion.IdVenta]
 * @apiParam {string} [OrdenesProduccion.Observaciones]

  * @apiParamExample {json} Request-Example:
{
    "OrdenesProduccion": {
		"IdVenta": 31,
		"Observaciones": ""
	}
}
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":{
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
func (pc *OrdenesProduccionController) Crear(c echo.Context) error {

	ordenProduccion := structs.OrdenesProduccion{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["OrdenesProduccion"], &ordenProduccion)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorOrdenesProduccion := gestores.GestorOrdenesProduccion{
		DbHandler: pc.DbHanlder,
	}
	result, err := gestorOrdenesProduccion.Crear(ordenProduccion, *token)

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
 * @api {POST} /ordenesProduccion/pasar-a-creado Confirmar OrdenProduccion
 * @apiDescription Permite confirmar la creación de un ordenProduccion
 * @apiGroup OrdenesProduccion
 * @apiHeader {String} Authorization
 * @apiParam {Object} OrdenesProduccion
 * @apiParam {int} OrdenesProduccion.IdOrdenProduccion

  * @apiParamExample {json} Request-Example:
{
    "OrdenesProduccion": {
		"IdOrdenProduccion": 2
	}
}
* @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":{
		"OrdenesProduccion":{
			"Estado": "C",
			"FechaAlta": "2020-08-22 20:02:10.000000",
			"IdCliente": 3,
			"IdOrdenProduccion": 2,
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
//PasarAPendiente PasarAPendiente
func (pc *OrdenesProduccionController) PasarAPendiente(c echo.Context) error {

	ordenProduccion := structs.OrdenesProduccion{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["OrdenesProduccion"], &ordenProduccion)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	ordenesProduccionService := models.OrdenesProduccionService{
		DbHanlder:         pc.DbHanlder,
		OrdenesProduccion: ordenProduccion,
	}
	result, err := ordenesProduccionService.PasarAPendiente(*token)

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
 * @api {POST} /ordenesProduccion/modificar Modificar OrdenProduccion
 * @apiDescription Permite modificar un ordenProduccion
 * @apiGroup OrdenesProduccion
 * @apiHeader {String} Authorization
 * @apiParam {Object} OrdenesProduccion
 * @apiParam {int} OrdenesProduccion.IdOrdenProduccion
 * @apiParam {int} OrdenesProduccion.IdCliente
 * @apiParam {int} OrdenesProduccion.IdUbicacion
 * @apiParam {string} [OrdenesProduccion.Observaciones]
 * @apiParamExample {json} Request-Example:
{
    "OrdenesProduccion": {
		"IdOrdenProduccion":1,
		"IdCliente": 12,
		"IdUbicacion": 13
	}
}
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":{
		"OrdenesProduccion":{
			"Estado": "E",
			"FechaAlta": "2020-08-22 20:02:10.000000",
			"IdCliente": 3,
			"IdOrdenProduccion": 2,
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
//Modificar Modificar
func (pc *OrdenesProduccionController) Modificar(c echo.Context) error {

	ordenProduccion := structs.OrdenesProduccion{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["OrdenesProduccion"], &ordenProduccion)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorOrdenesProduccion := gestores.GestorOrdenesProduccion{
		DbHandler: pc.DbHanlder,
	}
	result, err := gestorOrdenesProduccion.Modificar(ordenProduccion, *token)

	if err != nil || result == nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error:     nil,
		Respuesta: result,
	}

	return c.JSON(http.StatusOK, response)
}

/*
 * @api {POST} /ordenesProduccion/dame Dame OrdenProduccion
 * @apiDescription Permite instanciar un ordenProduccion a partir de su Id
 * @apiGroup OrdenesProduccion
 * @apiHeader {String} Authorization
 * @apiParam {Object} OrdenesProduccion
 * @apiParam {int} OrdenesProduccion.IdOrdenProduccion

  * @apiParamExample {json} Request-Example:
{
    "OrdenesProduccion": {
		"IdOrdenProduccion":1
	}
}
 * @apiSuccessExample {json} Success-Response:
{
    "error": null,
    "respuesta": {
        "OrdenesProduccion": {
            "Estado": "C",
            "IdVenta": null,
            "FechaAlta": "2020-08-17 18:51:47.000000",
            "IdCliente": 3,
            "IdUsuario": 1,
            "IdUbicacion": 1,
            "IdOrdenProduccion": 1,
            "Observaciones": null,
            "PeriodoValidez": 15
        },
        "LineasOrdenProduccion": [
            {
                "Telas": {
                    "Tela": "Cuero de cebra",
                    "IdTela": 5
                },
                "Lustres": null,
                "Productos": {
                    "Producto": "Silla comedor",
                    "IdProducto": 5
                },
                "LineasProducto": {
                    "Cantidad": 4,
                    "PrecioUnitario": 39000.0,
                    "IdLineaProducto": 1,
                    "IdProductoFinal": 1
                },
                "ProductosFinales": {
                    "IdTela": 5,
                    "IdLustre": null,
                    "FechaAlta": "2020-07-10 15:26:14.000000",
                    "IdProducto": 5,
                    "IdProductoFinal": 1
                }
            }
        ]
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
func (pc *OrdenesProduccionController) Dame(c echo.Context) error {

	ordenProduccion := structs.OrdenesProduccion{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["OrdenesProduccion"], &ordenProduccion)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	ordenesProduccionService := models.OrdenesProduccionService{
		DbHanlder:         pc.DbHanlder,
		OrdenesProduccion: ordenProduccion,
	}
	result, err := ordenesProduccionService.Dame(*token)

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
 * @api {POST} /ordenesProduccion/lineasOrdenProduccion/dame Dame Linea de OrdenProduccion
 * @apiDescription Permite instanciar una linea de ordenProduccion a partir de su Id
 * @apiGroup OrdenesProduccion
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
func (pc *OrdenesProduccionController) DameLineaOrdenProduccion(c echo.Context) error {

	lineaOrdenProduccion := structs.LineasProducto{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["LineasProducto"], &lineaOrdenProduccion)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnauthorized)
	}

	lineasOrdenProduccionService := models.LineasOrdenProduccionService{
		DbHandler:      pc.DbHanlder,
		LineasProducto: lineaOrdenProduccion,
	}
	result, err := lineasOrdenProduccionService.Dame(*token)

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
 * @api {POST} /ordenesProduccion/lineasOrdenProduccion/borrar Borrar Linea de OrdenProduccion
 * @apiDescription Permite borrar una linea de ordenProduccion
 * @apiGroup OrdenesProduccion
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
//BorrarLineaOrdenesProduccion BorrarLineaOrdenesProduccion
func (pc *OrdenesProduccionController) BorrarLineaOrdenesProduccion(c echo.Context) error {

	lineaOrdenProduccion := structs.LineasProducto{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["LineasProducto"], &lineaOrdenProduccion)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnauthorized)
	}

	ordenesProduccionService := models.OrdenesProduccionService{
		DbHanlder: pc.DbHanlder,
	}
	err = ordenesProduccionService.BorrarLineaOrdenProduccion(lineaOrdenProduccion, *token)

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
 * @api {POST} /ordenesProduccion/lineasOrdenProduccion/modificar Modificar Linea de OrdenProduccion
 * @apiDescription Permite modificar una linea de ordenProduccion
 * @apiGroup OrdenesProduccion
 * @apiHeader {String} Authorization
 * @apiParam {Object} LineasProducto
 * @apiParam {int} LineasProducto.IdLineaProducto
 * @apiParam {int} LineasProducto.IdOrdenProduccion
 * @apiParam {float} LineasProducto.PrecioUnitario
 * @apiParam {int} LineasProducto.Cantidad
 * @apiParam {Object} ProductosFinales
 * @apiParam {int} ProductosFinales.IdProducto
 * @apiParam {int} [ProductosFinales.IdTela]
 * @apiParam {int} [ProductosFinales.IdLuste]

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
//ModificarLineaOrdenProduccion ModificarLineaOrdenProduccion
func (pc *OrdenesProduccionController) ModificarLineaOrdenProduccion(c echo.Context) error {

	lineaOrdenProduccion := structs.LineasProducto{}
	productoFinal := structs.ProductosFinales{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["LineasProducto"], &lineaOrdenProduccion)
	mapstructure.Decode(jsonMap["ProductosFinales"], &productoFinal)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnauthorized)
	}

	ordenesProduccionService := models.OrdenesProduccionService{
		DbHanlder: pc.DbHanlder,
	}
	result, err := ordenesProduccionService.ModificarLineaOrdenProduccion(lineaOrdenProduccion, productoFinal, *token)

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
 * @api {POST} /ordenesProduccion/lineasOrdenProduccion/crear Crear Linea de OrdenProduccion
 * @apiDescription Permite crear una linea de ordenProduccion
 * @apiGroup OrdenesProduccion
 * @apiHeader {String} Authorization
 * @apiParam {Object} LineasProducto
 * @apiParam {int} LineasProducto.IdOrdenProduccion
 * @apiParam {float} LineasProducto.PrecioUnitario
 * @apiParam {int} LineasProducto.Cantidad
 * @apiParam {Object} ProductosFinales
 * @apiParam {int} ProductosFinales.IdProducto
 * @apiParam {int} [ProductosFinales.IdTela]
 * @apiParam {int} [ProductosFinales.IdLuste]

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
//CrearLineaOrdenProduccion CrearLineaOrdenProduccion
func (pc *OrdenesProduccionController) CrearLineaOrdenProduccion(c echo.Context) error {

	lineaOrdenProduccion := structs.LineasProducto{}
	productoFinal := structs.ProductosFinales{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["LineasProducto"], &lineaOrdenProduccion)
	mapstructure.Decode(jsonMap["ProductosFinales"], &productoFinal)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnauthorized)
	}

	ordenesProduccionService := models.OrdenesProduccionService{
		DbHanlder: pc.DbHanlder,
	}
	result, err := ordenesProduccionService.CrearLineaOrdenProduccion(lineaOrdenProduccion, productoFinal, *token)

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
 * @api {POST} /ordenesProduccion/borrar Borrar OrdenProduccion
 * @apiDescription Permite borrar un ordenProduccion
 * @apiGroup OrdenesProduccion
 * @apiHeader {String} Authorization
 * @apiParam {Object} OrdenesProduccion
 * @apiParam {int} OrdenesProduccion.IdOrdenProduccion

  * @apiParamExample {json} Request-Example:
{
    "OrdenesProduccion": {
		"IdOrdenProduccion": 2
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
func (pc *OrdenesProduccionController) Borrar(c echo.Context) error {

	ordenProduccion := structs.OrdenesProduccion{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["OrdenesProduccion"], &ordenProduccion)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorOrdenesProduccion := gestores.GestorOrdenesProduccion{
		DbHandler: pc.DbHanlder,
	}
	err = gestorOrdenesProduccion.Borrar(ordenProduccion, *token)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error:     nil,
		Respuesta: nil,
	}

	return c.JSON(http.StatusOK, response)
}

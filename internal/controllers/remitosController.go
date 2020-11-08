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

//RemitosController RemitosController
type RemitosController struct {
	DbHandler *db.DbHandler
}

/**
 * @api {POST} /remitos/crear Crear Remito
 * @apiDescription Permite crear un remito
 * @apiGroup Remitos
 * @apiHeader {String} Authorization
 * @apiParam {Object} Remitos
 * @apiParam {int} [Remitos.IdUbicacion]
 * @apiParam {int} Remitos.Tipo
 * @apiParam {string} [Remitos.Observaciones]

  * @apiParamExample {json} Request-Example:
{
    "Remitos": {
		"IdDomicilio": 12,
		"IdUbicacion": 13,
		"Tipo":"E"
	}
}
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":{
		"Remitos":{
			"IdRemito": 2,
			"IdUbicacion": 1,
			"IdDomicilio": 2,
			"IdUsuario": 2,
			"FechaEntrega": "2020-09-22 20:02:10.000000",
			"FechaAlta": "2020-08-22 20:02:10.000000",
			"Estado":"C",
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
func (rc *RemitosController) Crear(c echo.Context) error {

	remito := structs.Remitos{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Remitos"], &remito)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorRemitos := gestores.GestorRemitos{
		DbHandler: rc.DbHandler,
	}
	result, err := gestorRemitos.Crear(remito, *token)

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
 * @api {POST} /remitos/borrar Borrar Remito
 * @apiDescription Permite borrar un remito
 * @apiGroup Remitos
 * @apiHeader {String} Authorization
 * @apiParam {Object} Remitos
 * @apiParam {int} Remitos.IdRemito

  * @apiParamExample {json} Request-Example:
{
    "Remitos": {
		"IdRemito": 1
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
//Crear Crear
func (rc *RemitosController) Borrar(c echo.Context) error {

	remito := structs.Remitos{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Remitos"], &remito)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorRemitos := gestores.GestorRemitos{
		DbHandler: rc.DbHandler,
	}
	err = gestorRemitos.Borrar(remito, *token)

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
 * @api {POST} /remitos/cancelar Cancelar Remito
 * @apiDescription Permite cancelar un remito
 * @apiGroup Remitos
 * @apiHeader {String} Authorization
 * @apiParam {Object} Remitos
 * @apiParam {int} Remitos.IdRemito

  * @apiParamExample {json} Request-Example:
{
    "Remitos": {
		"IdRemito": 2
	}
}
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":{
		"Remitos":{
			"IdRemito": 2,
			"IdUbicacion": 1,
			"IdDomicilio": 2,
			"IdUsuario": 2,
			"FechaEntrega": "2020-09-22 20:02:10.000000",
			"FechaAlta": "2020-08-22 20:02:10.000000",
			"Estado":"C",
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
//Cancelar Cancelar
func (rc *RemitosController) Cancelar(c echo.Context) error {

	remito := structs.Remitos{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Remitos"], &remito)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	remitosService := models.RemitosService{
		DbHandler: rc.DbHandler,
		Remitos:   remito,
	}
	result, err := remitosService.Cancelar(*token)

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
 * @api {POST} /remitos/descancelar Descancelar Remito
 * @apiDescription Permite descancelar un remito
 * @apiGroup Remitos
 * @apiHeader {String} Authorization
 * @apiParam {Object} Remitos
 * @apiParam {int} Remitos.IdRemito

  * @apiParamExample {json} Request-Example:
{
    "Remitos": {
		"IdRemito": 2
	}
}
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":{
		"Remitos":{
			"IdRemito": 2,
			"IdUbicacion": 1,
			"IdDomicilio": 2,
			"IdUsuario": 2,
			"FechaEntrega": "2020-09-22 20:02:10.000000",
			"FechaAlta": "2020-08-22 20:02:10.000000",
			"Estado":"C",
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
//Descancelar Descancelar
func (rc *RemitosController) Descancelar(c echo.Context) error {

	remito := structs.Remitos{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Remitos"], &remito)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	remitosService := models.RemitosService{
		DbHandler: rc.DbHandler,
		Remitos:   remito,
	}
	result, err := remitosService.Descancelar(*token)

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
 * @api {POST} /remitos/entregar Entregar Remito
 * @apiDescription Permite entregar un remito
 * @apiGroup Remitos
 * @apiHeader {String} Authorization
 * @apiParam {Object} Remitos
 * @apiParam {int} Remitos.IdRemito
 * @apiParam {string} [Remitos.FechaEntrega]

  * @apiParamExample {json} Request-Example:
{
    "Remitos": {
		"IdRemito": 2,
		"FechaEntrega":"2020-09-22"
	}
}
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":{
		"Remitos":{
			"IdRemito": 2,
			"IdUbicacion": 1,
			"IdDomicilio": 2,
			"IdUsuario": 2,
			"FechaEntrega": "2020-09-22 20:02:10.000000",
			"FechaAlta": "2020-08-22 20:02:10.000000",
			"Estado":"C",
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
//Entregar Entregar
func (rc *RemitosController) Entregar(c echo.Context) error {

	remito := structs.Remitos{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Remitos"], &remito)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	remitosService := models.RemitosService{
		DbHandler: rc.DbHandler,
		Remitos:   remito,
	}
	result, err := remitosService.Entregar(*token)

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
 * @api {POST} /remitos/pasarACreado Pasar Remito a creado
 * @apiDescription Permite pasar a creado un remito
 * @apiGroup Remitos
 * @apiHeader {String} Authorization
 * @apiParam {Object} Remitos
 * @apiParam {int} Remitos.IdRemito
 * @apiParam {string} Remitos.FechaEntrega

  * @apiParamExample {json} Request-Example:
{
    "Remitos": {
		"IdRemito": 2
	}
}
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":{
		"Remitos":{
			"IdRemito": 2,
			"IdUbicacion": 1,
			"IdDomicilio": 2,
			"IdUsuario": 2,
			"FechaEntrega": "2020-09-22 20:02:10.000000",
			"FechaAlta": "2020-08-22 20:02:10.000000",
			"Estado":"C",
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
//Entregar Entregar
func (rc *RemitosController) PasarACreado(c echo.Context) error {

	remito := structs.Remitos{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Remitos"], &remito)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	remitosService := models.RemitosService{
		DbHandler: rc.DbHandler,
		Remitos:   remito,
	}
	result, err := remitosService.PasarACreado(*token)

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
 * @api {POST} /remitos Buscar Remitos
 * @apiDescription Permite buscar un remito
 * @apiGroup Remitos
 * @apiHeader {String} Authorization
 * @apiParam {Object} [Remitos]
 * @apiParam {int} [Remitos.IdUsuario]
 * @apiParam {int} [Remitos.IdUbicacion]
 * @apiParam {string} [Presupuestos.Tipo]
 * @apiParam {string} [Presupuestos.Estado]
 * @apiParam {Object} [ProductosFinales]
 * @apiParam {int} [ProductosFinales.IdProducto]
 * @apiParam {int} [ProductosFinales.IdTela]
 * @apiParam {int} [ProductosFinales.IdLustre]
 * @apiParam {object} [ParametrosBusqueda]
 * @apiParam {string} [ParametrosBusqueda.FechaEntregaDesde]
 * @apiParam {string} [ParametrosBusqueda.FechaEntregaHasta]
 * @apiParam {Object} [Paginaciones]
 * @apiParam {int} [Paginaciones.Pagina]
 * @apiParam {int} [Paginaciones.LongitudPagina]

  * @apiParamExample {json} Request-Example:
{
    "Remitos": {
		"IdUbicacion": 13,
		"IdUsuario":1,
		"Tipo":"E",
		"Estado":"T"
	},
	"ProductosFinales":{
		"IdProducto":1,
		"IdTela":0,
		"IdLustre"
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
			{
				"LineasPresupuesto":[
					{
						"LineasProducto":{
							"Cantidad": 3,
							"IdLineaProducto": 3,
							"IdProductoFinal": 28,
							"PrecioUnitario": 108.21
						},
						"Lustres":{
							"IdLustre": 1,
							"Lustre": "Lustre1"
						},
						"Productos":{
							"IdProducto": 12,
							"Producto": "Wilkinson"
						},
						"ProductosFinales":{
							"FechaAlta": "2020-08-24 00:15:46.000000",
							"IdLustre": 1,
							"IdProducto": 12,
							"IdProductoFinal": 28,
							"IdTela": 5
						},
						"Telas":{
							"IdTela": 5,
							"Tela": "Cuero de cebra"
						}
					}
				],
				"Remitos":{
					"IdRemito": 2,
					"IdUbicacion": 1,
					"IdDomicilio": 2,
					"IdUsuario": 2,
					"FechaEntrega": "2020-09-22 20:02:10.000000",
					"FechaAlta": "2020-08-22 20:02:10.000000",
					"Estado":"C",
					"Observaciones": null
				}
			}
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
func (rc *RemitosController) Buscar(c echo.Context) error {

	remito := structs.Remitos{}
	productoFinal := structs.ProductosFinales{}
	paginacion := structs.Paginaciones{}
	parametrosBusqueda := structs.ParametrosBusqueda{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Remitos"], &remito)
	mapstructure.Decode(jsonMap["ProductosFinales"], &productoFinal)
	mapstructure.Decode(jsonMap["Paginaciones"], &paginacion)
	mapstructure.Decode(jsonMap["ParametrosBusqueda"], &parametrosBusqueda)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorRemitos := gestores.GestorRemitos{
		DbHandler: rc.DbHandler,
	}
	result, err := gestorRemitos.Buscar(remito, productoFinal, parametrosBusqueda, paginacion, *token)

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
 * @api {POST} /remitos/dame Dame Remitos
 * @apiDescription Permite instanciar un remito a partir de su Id
 * @apiGroup Remitos
 * @apiHeader {String} Authorization
 * @apiParam {Object} Remitos
 * @apiParam {int} Remitos.IdRemito
 * @apiParamExample {json} Request-Example:
{
    "Remitos": {
		"idRemito": 13
	}
}
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":{
		"resultado":
		{
			"LineasPresupuesto":[
				{
					"LineasProducto":{
						"Cantidad": 3,
						"IdLineaProducto": 3,
						"IdProductoFinal": 28,
						"PrecioUnitario": 108.21
					},
					"Lustres":{
						"IdLustre": 1,
						"Lustre": "Lustre1"
					},
					"Productos":{
						"IdProducto": 12,
						"Producto": "Wilkinson"
					},
					"ProductosFinales":{
						"FechaAlta": "2020-08-24 00:15:46.000000",
						"IdLustre": 1,
						"IdProducto": 12,
						"IdProductoFinal": 28,
						"IdTela": 5
					},
					"Telas":{
						"IdTela": 5,
						"Tela": "Cuero de cebra"
					}
				}
			],
			"Remitos":{
				"IdRemito": 2,
				"IdUbicacion": 1,
				"IdDomicilio": 2,
				"IdUsuario": 2,
				"FechaEntrega": "2020-09-22 20:02:10.000000",
				"FechaAlta": "2020-08-22 20:02:10.000000",
				"Estado":"C",
				"Observaciones": null
			}
		},
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
//Dame Dame
func (rc *RemitosController) Dame(c echo.Context) error {

	remito := structs.Remitos{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Remitos"], &remito)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	remitosService := models.RemitosService{
		DbHandler: rc.DbHandler,
		Remitos:   remito,
	}
	result, err := remitosService.Dame(*token)

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
 * @api {POST} /remitos/lineasRemito/crear Crear Linea de Remito
 * @apiDescription Permite crear una linea de remito
 * @apiGroup LineasRemito
 * @apiHeader {String} Authorization
 * @apiParam {Object} LineasProducto
 * @apiParam {int} LineasProducto.IdReferencia
 * @apiParam {int} LineasProducto.Cantidad
 * @apiParam {int} [LineasProducto.IdUbicacion]
 * @apiParam {Object} ProductosFinales
 * @apiParam {int} ProductosFinales.IdProducto
 * @apiParam {int} [ProductosFinales.IdTela]
 * @apiParam {int} [ProductosFinales.IdLuste]

 * @apiParamExample {json} Request-Example:
{
  "LineasProducto":{
    "IdReferencia":2,
	"Cantidad":3,
	"IdUbicacion": 1
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
//CrearLineaRemito CrearLineaRemito
func (rc *RemitosController) CrearLineaRemito(c echo.Context) error {

	lineaRemito := structs.LineasProducto{}
	productoFinal := structs.ProductosFinales{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["LineasProducto"], &lineaRemito)
	mapstructure.Decode(jsonMap["ProductosFinales"], &productoFinal)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnauthorized)
	}

	remitosService := models.RemitosService{
		DbHandler: rc.DbHandler,
	}
	result, err := remitosService.CrearLineaRemito(lineaRemito, productoFinal, *token)

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
 * @api {POST} /remitos/lineasRemito/modificar Modificar Linea de Remito
 * @apiDescription Permite modificar una linea de remito
 * @apiGroup LineasRemito
 * @apiHeader {String} Authorization
 * @apiParam {Object} LineasProducto
 * @apiParam {int} LineasProducto.IdLineaProducto
 * @apiParam {int} LineasProducto.IdReferencia
 * @apiParam {int} LineasProducto.Cantidad
 * @apiParam {int} [LineasProducto.IdUbicacion]
 * @apiParam {Object} ProductosFinales
 * @apiParam {int} ProductosFinales.IdProducto
 * @apiParam {int} [ProductosFinales.IdTela]
 * @apiParam {int} [ProductosFinales.IdLuste]

 * @apiParamExample {json} Request-Example:
{
  "LineasProducto":{
	"IdLineaProducto": 4,
    "IdReferencia":2,
	"Cantidad":3,
	"IdUbicacion": 1
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
//ModificarLineaRemito ModificarLineaRemito
func (rc *RemitosController) ModificarLineaRemito(c echo.Context) error {

	lineaRemito := structs.LineasProducto{}
	productoFinal := structs.ProductosFinales{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["LineasProducto"], &lineaRemito)
	mapstructure.Decode(jsonMap["ProductosFinales"], &productoFinal)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnauthorized)
	}

	remitosService := models.RemitosService{
		DbHandler: rc.DbHandler,
	}
	result, err := remitosService.ModificarLineaRemito(lineaRemito, productoFinal, *token)

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
 * @api {POST} /remitos/lineasRemito/borrar Borrar Linea de Remito
 * @apiDescription Permite borrar una linea de remito
 * @apiGroup LineasRemito
 * @apiHeader {String} Authorization
 * @apiParam {Object} LineasProducto
 * @apiParam {int} LineasProducto.IdLineaProducto

 * @apiParamExample {json} Request-Example:
{
  "LineasProducto":{
	"IdLineaProducto": 4,
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
//BorrarLineaRemito BorrarLineaRemito
func (rc *RemitosController) BorrarLineaRemito(c echo.Context) error {

	lineaRemito := structs.LineasProducto{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["LineasProducto"], &lineaRemito)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnauthorized)
	}

	remitosService := models.RemitosService{
		DbHandler: rc.DbHandler,
	}
	err = remitosService.BorrarLineaRemito(lineaRemito, *token)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error:     nil,
		Respuesta: nil,
	}

	return c.JSON(http.StatusOK, response)
}

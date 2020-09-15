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

//PresupuestosController PresupuestosController
type PresupuestosController struct {
	DbHanlder *db.DbHandler
}

/**
 * @api {POST} /presupuestos Buscar Presupuestos
 * @apiDescription Permite buscar un presupuesto
 * @apiGroup Presupuestos
 * @apiHeader {String} Authorization
 * @apiParam {Object} [Presupuestos]
 * @apiParam {int} [Presupuestos.IdCliente]
 * @apiParam {int} [Presupuestos.IdUbicacion]
 * @apiParam {int} [Presupuestos.IdUsuario]
 * @apiParam {string} [Presupuestos.Estado]
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
    "Presupuestos": {
		"IdCliente": 12,
		"IdUbicacion": 13,
		"IdUsuario":1
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
				"Presupuestos":{
					"Estado": "C",
					"FechaAlta": "2020-08-22 20:02:10.000000",
					"IdCliente": 3,
					"IdPresupuesto": 2,
					"IdUbicacion": 1,
					"IdUsuario": 1,
					"IdVenta": null,
					"Observaciones": null,
					"PeriodoValidez": 15,
					"_PrecioTotal": 324.63
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
func (pc *PresupuestosController) Buscar(c echo.Context) error {

	presupuesto := structs.Presupuestos{}
	productoFinal := structs.ProductosFinales{}
	paginacion := structs.Paginaciones{}
	parametrosBusqueda := structs.ParametrosBusqueda{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Presupuestos"], &presupuesto)
	mapstructure.Decode(jsonMap["ProductosFinales"], &productoFinal)
	mapstructure.Decode(jsonMap["Paginaciones"], &paginacion)
	mapstructure.Decode(jsonMap["ParametrosBusqueda"], &parametrosBusqueda)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorPresupuestos := gestores.GestorPresupuestos{
		DbHandler: pc.DbHanlder,
	}
	result, err := gestorPresupuestos.Buscar(presupuesto, productoFinal, parametrosBusqueda, paginacion, *token)

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
 * @api {POST} /presupuestos/crear Crear Presupuesto
 * @apiDescription Permite crear un presupuesto
 * @apiGroup Presupuestos
 * @apiHeader {String} Authorization
 * @apiParam {Object} Presupuestos
 * @apiParam {int} Presupuestos.IdCliente
 * @apiParam {int} Presupuestos.IdUbicacion
 * @apiParam {string} [Presupuestos.Observaciones]

  * @apiParamExample {json} Request-Example:
{
    "Presupuestos": {
		"IdCliente": 12,
		"IdUbicacion": 13
	}
}
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":{
		"Presupuestos":{
			"Estado": "E",
			"FechaAlta": "2020-08-22 20:02:10.000000",
			"IdCliente": 3,
			"IdPresupuesto": 2,
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
//Crear Crear
func (pc *PresupuestosController) Crear(c echo.Context) error {

	presupuesto := structs.Presupuestos{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Presupuestos"], &presupuesto)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorPresupuestos := gestores.GestorPresupuestos{
		DbHandler: pc.DbHanlder,
	}
	result, err := gestorPresupuestos.Crear(presupuesto, *token)

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
 * @api {POST} /presupuestos/pasar-a-creado Confirmar Presupuesto
 * @apiDescription Permite confirmar la creación de un presupuesto
 * @apiGroup Presupuestos
 * @apiHeader {String} Authorization
 * @apiParam {Object} Presupuestos
 * @apiParam {int} Presupuestos.IdPresupuesto

  * @apiParamExample {json} Request-Example:
{
    "Presupuestos": {
		"IdPresupuesto": 2
	}
}
* @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":{
		"Presupuestos":{
			"Estado": "C",
			"FechaAlta": "2020-08-22 20:02:10.000000",
			"IdCliente": 3,
			"IdPresupuesto": 2,
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
//PasarACreado PasarACreado
func (pc *PresupuestosController) PasarACreado(c echo.Context) error {

	presupuesto := structs.Presupuestos{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Presupuestos"], &presupuesto)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	presupuestosService := models.PresupuestosService{
		DbHanlder:    pc.DbHanlder,
		Presupuestos: presupuesto,
	}
	result, err := presupuestosService.PasarACreado(*token)

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
 * @api {POST} /presupuestos/modificar Modificar Presupuesto
 * @apiDescription Permite modificar un presupuesto
 * @apiGroup Presupuestos
 * @apiHeader {String} Authorization
 * @apiParam {Object} Presupuestos
 * @apiParam {int} Presupuestos.IdPresupuesto
 * @apiParam {int} Presupuestos.IdCliente
 * @apiParam {int} Presupuestos.IdUbicacion
 * @apiParam {string} [Presupuestos.Observaciones]
 * @apiParamExample {json} Request-Example:
{
    "Presupuestos": {
		"IdPresupuesto":1,
		"IdCliente": 12,
		"IdUbicacion": 13
	}
}
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":{
		"Presupuestos":{
			"Estado": "E",
			"FechaAlta": "2020-08-22 20:02:10.000000",
			"IdCliente": 3,
			"IdPresupuesto": 2,
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
func (pc *PresupuestosController) Modificar(c echo.Context) error {

	presupuesto := structs.Presupuestos{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Presupuestos"], &presupuesto)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorPresupuestos := gestores.GestorPresupuestos{
		DbHandler: pc.DbHanlder,
	}
	result, err := gestorPresupuestos.Modificar(presupuesto, *token)

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
 * @api {POST} /presupuestos/dame Dame Presupuesto
 * @apiDescription Permite instanciar un presupuesto a partir de su Id
 * @apiGroup Presupuestos
 * @apiHeader {String} Authorization
 * @apiParam {Object} Presupuestos
 * @apiParam {int} Presupuestos.IdPresupuesto

  * @apiParamExample {json} Request-Example:
{
    "Presupuestos": {
		"IdPresupuesto":1
	}
}
 * @apiSuccessExample {json} Success-Response:
{
    "error": null,
    "respuesta": {
        "Presupuestos": {
            "Estado": "C",
            "IdVenta": null,
            "FechaAlta": "2020-08-17 18:51:47.000000",
            "IdCliente": 3,
            "IdUsuario": 1,
            "IdUbicacion": 1,
            "IdPresupuesto": 1,
            "Observaciones": null,
            "PeriodoValidez": 15
        },
        "LineasPresupuesto": [
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
func (pc *PresupuestosController) Dame(c echo.Context) error {

	presupuesto := structs.Presupuestos{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Presupuestos"], &presupuesto)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	presupuestosService := models.PresupuestosService{
		DbHanlder:    pc.DbHanlder,
		Presupuestos: presupuesto,
	}
	result, err := presupuestosService.Dame(*token)

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
 * @api {POST} /presupuestos/lineasPresupuesto/dame Dame Linea de Presupuesto
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
func (pc *PresupuestosController) DameLineaPresupuesto(c echo.Context) error {

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
		DbHandler:      pc.DbHanlder,
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

/**
 * @api {POST} /presupuestos/lineasPresupuesto/borrar Borrar Linea de Presupuesto
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
//BorrarLineaPresupuestos BorrarLineaPresupuestos
func (pc *PresupuestosController) BorrarLineaPresupuestos(c echo.Context) error {

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
		DbHanlder: pc.DbHanlder,
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
 * @api {POST} /presupuestos/lineasPresupuesto/modificar Modificar Linea de Presupuesto
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
//ModificarLineaPresupuesto ModificarLineaPresupuesto
func (pc *PresupuestosController) ModificarLineaPresupuesto(c echo.Context) error {

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
		DbHanlder: pc.DbHanlder,
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
 * @api {POST} /presupuestos/lineasPresupuesto/crear Crear Linea de Presupuesto
 * @apiDescription Permite crear una linea de presupuesto
 * @apiGroup Presupuestos
 * @apiHeader {String} Authorization
 * @apiParam {Object} LineasProducto
 * @apiParam {int} LineasProducto.IdPresupuesto
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
//CrearLineaPresupuesto CrearLineaPresupuesto
func (pc *PresupuestosController) CrearLineaPresupuesto(c echo.Context) error {

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
		DbHanlder: pc.DbHanlder,
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
 * @api {POST} /presupuestos/borrar Borrar Presupuesto
 * @apiDescription Permite borrar un presupuesto
 * @apiGroup Presupuestos
 * @apiHeader {String} Authorization
 * @apiParam {Object} Presupuestos
 * @apiParam {int} Presupuestos.IdPresupuesto

  * @apiParamExample {json} Request-Example:
{
    "Presupuestos": {
		"IdPresupuesto": 2
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
func (pc *PresupuestosController) Borrar(c echo.Context) error {

	presupuesto := structs.Presupuestos{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Presupuestos"], &presupuesto)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorPresupuestos := gestores.GestorPresupuestos{
		DbHandler: pc.DbHanlder,
	}
	err = gestorPresupuestos.Borrar(presupuesto, *token)

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
 * @api {POST} /presupuestos/transformarEnVenta Transformar presupuestos en venta
 * @apiDescription Permite transformar presupuestos en ventas
 * @apiGroup Presupuestos
 * @apiHeader {String} Authorization
 * @apiParam {Object} Ventas
 * @apiParam {int} Ventas.IdDomicilio
 * @apiParam {int} Ventas.IdUbicacion
 * @apiParam {String} [Ventas.Observaciones]
 * @apiParam {int[]} LineasPresupuesto
 * @apiParam {Object[]} LineasVenta
 * @apiParamExample {json} Request-Example:
{
  "Ventas":{
    "IdDomicilio":1,
    "IdUbicacion":1
  },
  "LineasPresupuesto":[1, 2, 3],
  "LineasVenta":[
	  {
		  "ProductosFinales":{
			  "IdProducto":1,
			  "IdTela":1,
			  "IdLustre":1
		  },
		  "LineasProducto":{
			  "Cantidad":1,
			  "PrecioUnitario":100.00
		  }
	  }
  ]
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
func (pc *PresupuestosController) TransformarEnVenta(c echo.Context) error {
	venta := structs.Ventas{}
	var LineasPresupuesto []int
	var lineasVenta []map[string]interface{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	mapstructure.Decode(jsonMap["Ventas"], &venta)
	mapstructure.Decode(jsonMap["LineasPresupuesto"], &LineasPresupuesto)
	mapstructure.Decode(jsonMap["LineasVenta"], &lineasVenta)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorPresupuestos := gestores.GestorPresupuestos{
		DbHandler: pc.DbHanlder,
	}
	result, err := gestorPresupuestos.TransformarEnVenta(venta, LineasPresupuesto, lineasVenta, *token)

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
 * @api {POST} /presupuestos/dameMultiples Dame multiples presupuestos
 * @apiDescription Permite instanciar mas de un presupuesto a partir de su Id
 * @apiGroup Presupuestos
 * @apiHeader {String} Authorization
 * @apiParam {int[]} Presupuestos
 * @apiParamExample {json} Request-Example:
{
  "Presupuestos":[1, 2, 3]
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
//DamePresupuestos DamePresupuestos
func (pc *PresupuestosController) DamePresupuestos(c echo.Context) error {
	var presupuestos []int

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Presupuestos"], &presupuestos)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorPresupuestos := gestores.GestorPresupuestos{
		DbHandler: pc.DbHanlder,
	}
	result, err := gestorPresupuestos.DamePresupuestos(presupuestos, *token)

	if err != nil || result == nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error:     nil,
		Respuesta: result,
	}

	return c.JSON(http.StatusOK, response)

}

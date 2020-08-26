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
 * @apiParam {Object} Presupuestos
 * @apiParam {int} Presupuestos.IdCliente
 * @apiParam {int} Presupuestos.IdUbicacion
 * @apiParam {int} Presupuestos.IdUsuario
 * @apiParam {Object} ProductosFinales
 * @apiParam {int} ProductosFinales.IdProducto
 * @apiParam {int} ProductosFinales.IdTela
 * @apiParam {int} ProductosFinales.IdLustre
 * @apiParam {Object} Paginaciones
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
	mapstructure.Decode(jsonMap["ProductosFianles"], &productoFinal)
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
 * @api {POST} /presupuestos/pasar-a-creado Confirmar la creación de un presupuesto
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

/*
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

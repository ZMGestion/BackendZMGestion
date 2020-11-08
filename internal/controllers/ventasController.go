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
 * @apiParam {int} Ventas.IdUbicacion
 * @apiParam {string} [Ventas.Observaciones]

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
 * @apiParam {int} Ventas.IdUbicacion
 * @apiParam {string} [Ventas.Observaciones]

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

/**
 * @api {POST} /ventas/lineasVentas/crear Crear linea de venta
 * @apiDescription Permite crear una linea de venta
 * @apiGroup Ventas
 * @apiHeader {String} Authorization
 * @apiParam {Object} LineasProducto
 * @apiParam {int} LineasProducto.IdReferencia
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
//CrearLineaVenta CrearLineaVenta
func (vc *VentasController) CrearLineaVenta(c echo.Context) error {

	lineaVenta := structs.LineasProducto{}
	productoFinal := structs.ProductosFinales{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["LineasProducto"], &lineaVenta)
	mapstructure.Decode(jsonMap["ProductosFinales"], &productoFinal)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnauthorized)
	}

	ventasService := models.VentasService{
		DbHandler: vc.DbHandler,
	}
	result, err := ventasService.CrearLineaVenta(lineaVenta, productoFinal, *token)

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
 * @api {POST} /presupuestos/lineasVenta/modificar Modificar linea de venta
 * @apiDescription Permite modificar una linea de venta
 * @apiGroup Ventas
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
//ModificarLineaVenta ModificarLineaVenta
func (vc *VentasController) ModificarLineaVenta(c echo.Context) error {

	lineaVenta := structs.LineasProducto{}
	productoFinal := structs.ProductosFinales{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["LineasProducto"], &lineaVenta)
	mapstructure.Decode(jsonMap["ProductosFinales"], &productoFinal)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnauthorized)
	}

	ventasService := models.VentasService{
		DbHandler: vc.DbHandler,
	}
	result, err := ventasService.ModificarLineaVenta(lineaVenta, productoFinal, *token)

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
 * @api {POST} /ventas/lineasVenta/borrar Borrar linea de venta
 * @apiDescription Permite borrar una linea de venta
 * @apiGroup Ventas
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
//BorrarLineaVenta BorrarLineaVenta
func (vc *VentasController) BorrarLineaVenta(c echo.Context) error {

	lineaVenta := structs.LineasProducto{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["LineasProducto"], &lineaVenta)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnauthorized)
	}

	ventasService := models.VentasService{
		DbHandler: vc.DbHandler,
	}
	err = ventasService.BorrarLineaVenta(lineaVenta, *token)

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
 * @api {POST} /ventas/lineasVenta/dame Dame linea de venta
 * @apiDescription Permite instanciar una linea de venta a partir de su Id.
 * @apiGroup Ventas
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
			"Cantidad": 3,
			"IdLineaProducto": 51,
			"IdProductoFinal": 16,
			"PrecioUnitario": 108.21
		},
		"Lustres":{
			"IdLustre": 2,
			"Lustre": "Lustre2"
		},
		"Productos":{
			"IdProducto": 15,
			"Producto": "Becker"
		},
		"ProductosFinales":{
			"FechaAlta": "2020-08-17 21:11:41.000000",
			"IdLustre": 2,
			"IdProducto": 15,
			"IdProductoFinal": 16,
			"IdTela": 28
		},
		"Telas":{
			"IdTela": 28,
			"Tela": "Brownie"
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
//DameLineaVenta DameLineaVenta
func (vc *VentasController) DameLineaVenta(c echo.Context) error {

	lineaVenta := structs.LineasProducto{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["LineasProducto"], &lineaVenta)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnauthorized)
	}

	lineasVentaService := models.LineasVentaService{
		DbHandler:      vc.DbHandler,
		LineasProducto: lineaVenta,
	}
	result, err := lineasVentaService.Dame(*token)

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
 * @api {POST} /ventas/dame Dame Venta
 * @apiDescription Permite instanciar una venta a partir de su Id
 * @apiGroup Ventas
 * @apiHeader {String} Authorization
 * @apiParam {Object} Ventas
 * @apiParam {int} Ventas.IdVenta

  * @apiParamExample {json} Request-Example:
{
    "Vemtas": {
		"IdVenta":1
	}
}
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":{
		"Clientes":{
			"Apellidos": "Choua",
			"Nombres": "Loik",
			"RazonSocial": null
		},
		"Domicilios":{
			"Domicilio": "El Tipal Lote 13"
		},
		"LineasVenta":[],
		"Ubicaciones":{
			"Ubicacion": "Casa Central Tucumán"
		},
		"Usuarios":{
			"Apellidos": "Super Admin",
			"Nombres": "Adam"
		},
		"Ventas":{
			"Estado": "E",
			"FechaAlta": "2020-09-05 20:48:50.000000",
			"IdCliente": 3,
			"IdDomicilio": 14,
			"IdUbicacion": 1,
			"IdUsuario": 1,
			"IdVenta": 2,
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
//Dame Dame
func (vc *VentasController) Dame(c echo.Context) error {

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

	ventasService := models.VentasService{
		DbHandler: vc.DbHandler,
		Ventas:    venta,
	}
	result, err := ventasService.Dame(*token)

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
 * @api {POST} /ventas Buscar Ventas
 * @apiDescription Permite buscar una venta
 * @apiGroup Ventas
 * @apiHeader {String} Authorization
 * @apiParam {Object} [Ventas]
 * @apiParam {int} [Ventas.IdCliente]
 * @apiParam {int} [Ventas.IdUbicacion]
 * @apiParam {int} [Ventas.IdUsuario]
 * @apiParam {string} [Ventas.Estado]
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
    "Ventas": {
		"IdCliente": 12,
		"IdUbicacion": 13,
		"IdUsuario":1,
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
func (vc *VentasController) Buscar(c echo.Context) error {

	venta := structs.Ventas{}
	productoFinal := structs.ProductosFinales{}
	paginacion := structs.Paginaciones{}
	parametrosBusqueda := structs.ParametrosBusqueda{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Ventas"], &venta)
	mapstructure.Decode(jsonMap["ProductosFianles"], &productoFinal)
	mapstructure.Decode(jsonMap["Paginaciones"], &paginacion)
	mapstructure.Decode(jsonMap["ParametrosBusqueda"], &parametrosBusqueda)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorVentas := gestores.GestorVentas{
		DbHandler: vc.DbHandler,
	}
	result, err := gestorVentas.Buscar(venta, productoFinal, parametrosBusqueda, paginacion, *token)

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
 * @api {POST} /ventas/chequearPrecios Chequear Venta
 * @apiDescription Permite controlar si los precios de las lineas de ventas son los actuales
 * @apiGroup Ventas
 * @apiHeader {String} Authorization
 * @apiParam {Object} Ventas
 * @apiParam {int} Ventas.IdVenta
 * @apiParamExample {json} Request-Example:
{
    "Vemtas": {
		"IdVenta":1
	}
}
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":{
		"Clientes":{
			"Apellidos": "Choua",
			"Nombres": "Loik",
			"RazonSocial": null
		},
		"Domicilios":{
			"Domicilio": "El Tipal Lote 13"
		},
		"LineasVenta":[],
		"Ubicaciones":{
			"Ubicacion": "Casa Central Tucumán"
		},
		"Usuarios":{
			"Apellidos": "Super Admin",
			"Nombres": "Adam"
		},
		"Ventas":{
			"Estado": "E",
			"FechaAlta": "2020-09-05 20:48:50.000000",
			"IdCliente": 3,
			"IdDomicilio": 14,
			"IdUbicacion": 1,
			"IdUsuario": 1,
			"IdVenta": 2,
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
//ChequearPrecios ChequearPrecios
func (vc *VentasController) ChequearPrecios(c echo.Context) error {
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

	ventasService := models.VentasService{
		DbHandler: vc.DbHandler,
		Ventas:    venta,
	}

	result, err := ventasService.ChequearPrecios(*token)

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
 * @api {POST} /ventas/revisar Revisar Venta
 * @apiDescription Permite aceptar una venta donde los precios de las lineas de ventas son diferente a los actuales
 * @apiGroup Ventas
 * @apiHeader {String} Authorization
 * @apiParam {Object} Ventas
 * @apiParam {int} Ventas.IdVenta
 * @apiParamExample {json} Request-Example:
{
    "Vemtas": {
		"IdVenta":1
	}
}
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":{
		"Clientes":{
			"Apellidos": "Choua",
			"Nombres": "Loik",
			"RazonSocial": null
		},
		"Domicilios":{
			"Domicilio": "El Tipal Lote 13"
		},
		"LineasVenta":[],
		"Ubicaciones":{
			"Ubicacion": "Casa Central Tucumán"
		},
		"Usuarios":{
			"Apellidos": "Super Admin",
			"Nombres": "Adam"
		},
		"Ventas":{
			"Estado": "E",
			"FechaAlta": "2020-09-05 20:48:50.000000",
			"IdCliente": 3,
			"IdDomicilio": 14,
			"IdUbicacion": 1,
			"IdUsuario": 1,
			"IdVenta": 2,
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
//Revisar Revisar
func (vc *VentasController) Revisar(c echo.Context) error {
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

	ventasService := models.VentasService{
		DbHandler: vc.DbHandler,
		Ventas:    venta,
	}

	result, err := ventasService.Revisar(*token)

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
 * @api {POST} /ventas/lineasVenta/cancelar Cancelar linea de venta
 * @apiDescription Permite cancelar una linea de venta.
 * @apiGroup Ventas
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
//CancelarLineaVenta CancelarLineaVenta
func (vc *VentasController) CancelarLineaVenta(c echo.Context) error {
	lineaVenta := structs.LineasProducto{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["LineasProducto"], &lineaVenta)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnauthorized)
	}

	lineasVentaService := models.LineasVentaService{
		DbHandler:      vc.DbHandler,
		LineasProducto: lineaVenta,
	}
	result, err := lineasVentaService.Cancelar(*token)

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
 * @api {POST} /ventas/cancelar Cancelar Venta
 * @apiDescription Permite cancelar una venta
 * @apiGroup Ventas
 * @apiHeader {String} Authorization
 * @apiParam {Object} Ventas
 * @apiParam {int} Ventas.IdVenta
 * @apiParamExample {json} Request-Example:
{
    "Vemtas": {
		"IdVenta":1
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
//Revisar Revisar
func (vc *VentasController) Cancelar(c echo.Context) error {
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

	ventasService := models.VentasService{
		DbHandler: vc.DbHandler,
		Ventas:    venta,
	}

	result, err := ventasService.Cancelar(*token)

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
 * @api {POST} /ventas/comprobantes/crear Crear Comprobante
 * @apiDescription Permite crear un comprobante
 * @apiGroup Comprobantes
 * @apiHeader {String} Authorization
 * @apiParam {Object} Comprobantes
 * @apiParam {int} Comprobantes.IdVenta
 * @apiParam {string} Comprobantes.Tipo
 * @apiParam {int} Comprobantes.NumeroComprobante
 * @apiParam {double} Comprobantes.Monto
 * @apiParam {string} [Comprobantes.Observaciones]

  * @apiParamExample {json} Request-Example:
{
    "Comprobantes": {
		"IdVenta": 12,
		"Tipo": 13,
		"NumeroComprobante": 12,
		"Monto":"201.20"
	}
}
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":{
		"Comprobantes":{
			"IdComprobante": 1,
			"IdVenta": 1,
			"IdUsuario": 2,
			"Tipo": A,
			"NumeroComprobante": 1,
			"Monto": 201.20,
			"FechaAlta": "2020-08-22 20:02:10.000000",
			"FechaBaja": null,
			"Observaciones": "",
			"Estado":"A"
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
//CrearComprobante CrearComprobante
func (vc *VentasController) CrearComprobante(c echo.Context) error {

	comprobante := structs.Comprobantes{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Comprobantes"], &comprobante)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	ventasService := models.VentasService{
		DbHandler: vc.DbHandler,
	}
	result, err := ventasService.CrearComprobante(comprobante, *token)

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
 * @api {POST} /ventas/comprobantes/modificar Modificar Comprobante
 * @apiDescription Permite modificar un comprobante
 * @apiGroup Comprobantes
 * @apiHeader {String} Authorization
 * @apiParam {Object} Comprobantes
 * @apiParam {int} Comprobantes.IdComprobante
 * @apiParam {int} Comprobantes.IdVenta
 * @apiParam {string} Comprobantes.Tipo
 * @apiParam {int} Comprobantes.NumeroComprobante
 * @apiParam {double} Comprobantes.Monto
 * @apiParam {string} [Comprobantes.Observaciones]

  * @apiParamExample {json} Request-Example:
{
    "Comprobantes": {
		"IdComprobante":1,
		"IdVenta": 12,
		"Tipo": 13,
		"NumeroComprobante": 12,
		"Monto":"201.20"
	}
}
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":{
		"Comprobantes":{
			"IdComprobante": 1,
			"IdVenta": 1,
			"IdUsuario": 2,
			"Tipo": A,
			"NumeroComprobante": 1,
			"Monto": 201.20,
			"FechaAlta": "2020-08-22 20:02:10.000000",
			"FechaBaja": null,
			"Observaciones": "",
			"Estado":"A"
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
//ModificarComprobante ModificarComprobante
func (vc *VentasController) ModificarComprobante(c echo.Context) error {

	comprobante := structs.Comprobantes{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Comprobantes"], &comprobante)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	ventasService := models.VentasService{
		DbHandler: vc.DbHandler,
	}
	result, err := ventasService.ModificarComprobante(comprobante, *token)

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
 * @api {POST} /ventas/comprobantes/buscar Buscar Comprobante
 * @apiDescription Permite buscar un comprobante
 * @apiGroup Comprobantes
 * @apiHeader {String} Authorization
 * @apiParam {Object} Comprobantes
 * @apiParam {int} [Comprobantes.IdVenta]
 * @apiParam {int} [Comprobantes.IdUsuario]
 * @apiParam {string} [Comprobantes.Tipo]
 * @apiParam {int} [Comprobantes.NumeroComprobante]

  * @apiParamExample {json} Request-Example:
{
    "Comprobantes": {
		"IdUsuario": 1,
		"IdVenta": 12,
		"Tipo": 13,
		"NumeroComprobante": 12,
	}
}
* @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":{
		"Comprobantes":[
			{
				"IdComprobante": 1,
				"IdVenta": 1,
				"IdUsuario": 2,
				"Tipo": A,
				"NumeroComprobante": 1,
				"Monto": 201.20,
				"FechaAlta": "2020-08-22 20:02:10.000000",
				"FechaBaja": null,
				"Observaciones": "",
				"Estado":"A"
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
//BuscarComprobante BuscarComprobante
func (vc *VentasController) BuscarComprobantes(c echo.Context) error {

	comprobante := structs.Comprobantes{}
	paginacion := structs.Paginaciones{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Comprobantes"], &comprobante)
	mapstructure.Decode(jsonMap["Paginaciones"], &paginacion)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	ventasService := models.VentasService{
		DbHandler: vc.DbHandler,
	}
	result, err := ventasService.BuscarComprobantes(comprobante, paginacion, *token)

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
 * @api {POST} /ventas/comprobantes/dame Dame Comprobante
 * @apiDescription Permite instanciar un comprobante a partir de su Id
 * @apiGroup Comprobantes
 * @apiParam {Object} Comprobantes
 * @apiParam {int} Comprobantes.IdComprobante
 * @apiHeader {String} Authorization
 * @apiParam {Object} Comprobantes
 * @apiParamExample {json} Request-Example:
{
    "Comprobantes": {
		"IdComprobante":1
	}
}
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":{
		"Comprobantes":{
			"IdComprobante": 1,
			"IdVenta": 1,
			"IdUsuario": 2,
			"Tipo": A,
			"NumeroComprobante": 1,
			"Monto": 201.20,
			"FechaAlta": "2020-08-22 20:02:10.000000",
			"FechaBaja": null,
			"Observaciones": "",
			"Estado":"A"
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
//DameComprobante DameComprobante
func (vc *VentasController) DameComprobante(c echo.Context) error {

	comprobante := structs.Comprobantes{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Comprobantes"], &comprobante)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	comprobantesService := models.ComprobantesService{
		DbHandler:   vc.DbHandler,
		Comprobante: &comprobante,
	}
	result, err := comprobantesService.Dame(*token)

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
 * @api {POST} /ventas/comprobantes/darAlta Dar alta Comprobante
 * @apiDescription Permite dar de alta un comprobante a partir de su Id
 * @apiGroup Comprobantes
 * @apiParam {Object} Comprobantes
 * @apiParam {int} Comprobantes.IdComprobante
 * @apiHeader {String} Authorization
 * @apiParam {Object} Comprobantes
 * @apiParamExample {json} Request-Example:
{
    "Comprobantes": {
		"IdComprobante":1
	}
}
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":{
		"Comprobantes":{
			"IdComprobante": 1,
			"IdVenta": 1,
			"IdUsuario": 2,
			"Tipo": A,
			"NumeroComprobante": 1,
			"Monto": 201.20,
			"FechaAlta": "2020-08-22 20:02:10.000000",
			"FechaBaja": null,
			"Observaciones": "",
			"Estado":"A"
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
//DarAlta DarAlta
func (vc *VentasController) DarAltaComprobante(c echo.Context) error {

	comprobante := structs.Comprobantes{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Comprobantes"], &comprobante)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	comprobantesService := models.ComprobantesService{
		DbHandler:   vc.DbHandler,
		Comprobante: &comprobante,
	}
	result, err := comprobantesService.DarAlta(*token)

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
 * @api {POST} /ventas/comprobantes/darbaja Dar baja Comprobante
 * @apiDescription Permite dar de baja un comprobante a partir de su Id
 * @apiGroup Comprobantes
 * @apiParam {Object} Comprobantes
 * @apiParam {int} Comprobantes.IdComprobante
 * @apiHeader {String} Authorization
 * @apiParam {Object} Comprobantes
 * @apiParamExample {json} Request-Example:
{
    "Comprobantes": {
		"IdComprobante":1
	}
}
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":{
		"Comprobantes":{
			"IdComprobante": 1,
			"IdVenta": 1,
			"IdUsuario": 2,
			"Tipo": B,
			"NumeroComprobante": 1,
			"Monto": 201.20,
			"FechaAlta": "2020-08-22 20:02:10.000000",
			"FechaBaja": "2020-08-22 21:02:10.000000",
			"Observaciones": "",
			"Estado":"A"
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
//DarBajaComprobante DarBajaComprobante
func (vc *VentasController) DarBajaComprobante(c echo.Context) error {

	comprobante := structs.Comprobantes{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Comprobantes"], &comprobante)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	comprobantesService := models.ComprobantesService{
		DbHandler:   vc.DbHandler,
		Comprobante: &comprobante,
	}
	result, err := comprobantesService.DarBaja(*token)

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
 * @api {POST} /ventas/comprobantes/borrar Borrar Comprobante
 * @apiDescription Permite modificar un comprobante
 * @apiGroup Comprobantes
 * @apiHeader {String} Authorization
 * @apiParam {Object} Comprobantes
 * @apiParam {int} Comprobantes.IdComprobante
 * @apiParamExample {json} Request-Example:
{
    "Comprobantes": {
		"IdComprobante":1
	}
}
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta": null
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
//BorrarComprobante BorrarComprobante
func (vc *VentasController) BorrarComprobante(c echo.Context) error {

	comprobante := structs.Comprobantes{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Comprobantes"], &comprobante)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	ventasService := models.VentasService{
		DbHandler: vc.DbHandler,
	}
	err = ventasService.BorrarComprobante(comprobante, *token)

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
 * @api {POST} /ventas/dameMultiple Dame multiple ventas
 * @apiDescription Permite instanciar mas de una venta a partir de su Id
 * @apiGroup Ventas
 * @apiHeader {String} Authorization
 * @apiParam {int[]} Ventas
 * @apiParamExample {json} Request-Example:
{
  "Ventas":[1, 2, 3]
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
//DameVentas DameVentas
func (vc *VentasController) DameVentas(c echo.Context) error {
	var ventas []int

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Ventas"], &ventas)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorVentas := gestores.GestorVentas{
		DbHandler: vc.DbHandler,
	}
	result, err := gestorVentas.DameVentas(ventas, *token)

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
 * @api {POST} /ventas/generarOrdenProduccion Generar orden de producción para una o más ventas
 * @apiDescription Permite genear una orden de producción a partir de una o más lineas de ventas pertenecientes a una o más ventas
 * @apiGroup Ventas
 * @apiHeader {String} Authorization
 * @apiParam {Object} OrdenesProduccion
 * @apiParam {String} [OrdenesProduccion.Observaciones]
 * @apiParam {Object[]} LineasVenta
 * @apiParam {int} LineasVenta.Cantidad
 * @apiParam {int} LineasVenta.IdProductoFinal
 * @apiParam {int[]} LineasVenta.IdLineasPadre
 * @apiParam {Object[]} LineasOrdenProduccion
 * @apiParamExample {json} Request-Example:
{
  "OrdenesProduccion":{
    "Observaciones": "",
  },
  "LineasVenta":[
	  {
		  "Cantidad": 7,
		  "IdProductoFinal": 59,
		  "IdLineasPadre": [237,408,531]
	  }
  ],
  "LineasOrdenProduccion":[
	  {
		  "ProductosFinales":{
			  "IdProducto":1,
			  "IdTela":1,
			  "IdLustre":1
		  },
		  "LineasProducto":{
			  "Cantidad":1
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
func (vc *VentasController) GenerarOrdenProduccion(c echo.Context) error {
	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorVentas := gestores.GestorVentas{
		DbHandler: vc.DbHandler,
	}
	result, err := gestorVentas.GenerarOrdenProduccion(jsonMap, *token)

	if err != nil || result == nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error:     nil,
		Respuesta: result,
	}

	return c.JSON(http.StatusOK, response)
}

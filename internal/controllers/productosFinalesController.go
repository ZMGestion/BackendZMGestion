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

type ProductosFinalesController struct {
	DbHandler *db.DbHandler
}

/**
 * @api {GET} /productosFinales/lustres Listar Lustres
 * @apiDescription Devuelve una lista de lustres
 * @apiGroup ProductosFinales
 * @apiSuccessExample {json} Success-Response:
 {
    "error": null,
    "respuesta": [
		{
			"Roles":{
				"IdRol": 1,
				"Rol": "Administradores",
				"FechaAlta": "2020-04-09 15:01:35.000000",
				"Observaciones": ""
			}
		},
		{
			"Roles":{
				"IdRol": 2,
				"Rol": "Vendedores",
				"FechaAlta": "2020-04-09 15:01:35.000000",
				"Observaciones": ""
			}
		}

    ]
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
func (lc *ProductosFinalesController) ListarLustres(c echo.Context) error {

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	productosFinalesService := models.ProductosFinalesService{
		DbHandler: lc.DbHandler,
	}

	result, err := productosFinalesService.ListarLustres(*token)

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
 * @api {POST} /productosFinales/crear Crear ProductoFinal Final
 * @apiDescription Permite crear un prodcuto final
 * @apiGroup ProductosFinales
 * @apiHeader {String} Authorization
 * @apiParam {Object} ProductosFinales
 * @apiParam {int} ProductosFinales.IdProducto
 * @apiParam {int} ProductosFinales.IdTela
 * @apiParam {string} ProductosFinales.IdLustre

  * @apiParamExample {json} Request-Example:
{
    "ProductosFinales": {
		"IdProducto": 12,
		"IdTela": 13,
		"IdLustre": 1
	}
}
 * @apiSuccessExample {json} Success-Response:
{
    "error": null,
    "respuesta": {
        "ProductosFinales": {
            "Estado": "A",
            "IdTela": 13,
            "IdLustre": 1,
            "FechaAlta": "2020-08-16 21:44:18.000000",
            "FechaBaja": null,
            "IdProducto": 12,
            "IdProductoFinal": 14
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
func (pfc *ProductosFinalesController) Crear(c echo.Context) error {

	productoFinal := structs.ProductosFinales{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["ProductosFinales"], &productoFinal)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorProductosFinales := gestores.GestorProductosFinales{
		DbHandler: pfc.DbHandler,
	}
	result, err := gestorProductosFinales.Crear(productoFinal, *token)

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
 * @api {POST} /productosFinales/modificar Modificar ProductoFinal Final
 * @apiDescription Permite modificar un prodcuto final
 * @apiGroup ProductosFinales
 * @apiHeader {String} Authorization
 * @apiParam {Object} ProductosFinales
 * @apiParam {int} ProductosFinales.IdProductoFinal
 * @apiParam {int} ProductosFinales.IdProducto
 * @apiParam {int} ProductosFinales.IdTela
 * @apiParam {string} ProductosFinales.IdLustre

  * @apiParamExample {json} Request-Example:
{
    "ProductosFinales": {
		"IdProductoFinal": 32,
		"IdProducto": 12,
		"IdTela": 13,
		"IdLustre": 1
	}
}
 * @apiSuccessExample {json} Success-Response:
{
    "error": null,
    "respuesta": {
        "ProductosFinales": {
            "IdProductoFinal": 32,
            "Estado": "A",
            "IdTela": 13,
            "IdLustre": 1,
            "FechaAlta": "2020-08-16 21:44:18.000000",
            "FechaBaja": null,
            "IdProducto": 12
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
//Modificar
func (pfc *ProductosFinalesController) Modificar(c echo.Context) error {

	productoFinal := structs.ProductosFinales{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["ProductosFinales"], &productoFinal)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorProductosFinales := gestores.GestorProductosFinales{
		DbHandler: pfc.DbHandler,
	}
	result, err := gestorProductosFinales.Modificar(productoFinal, *token)

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
 * @api {POST} /productosFinales/modificar Modificar ProductoFinal Final
 * @apiDescription Permite modificar un prodcuto final
 * @apiGroup ProductosFinales
 * @apiHeader {String} Authorization
 * @apiParam {Object} ProductosFinales
 * @apiParam {int} ProductosFinales.IdProductoFinal

  * @apiParamExample {json} Request-Example:
{
    "ProductosFinales": {
		"IdProductoFinal": 32
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
//Borrar
func (pfc *ProductosFinalesController) Borrar(c echo.Context) error {

	productoFinal := structs.ProductosFinales{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["ProductosFinales"], &productoFinal)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorProductosFinales := gestores.GestorProductosFinales{
		DbHandler: pfc.DbHandler,
	}
	err = gestorProductosFinales.Borrar(productoFinal, *token)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error: nil,
	}

	return c.JSON(http.StatusOK, response)
}

//DarAlta
func (pfc *ProductosFinalesController) DarAlta(c echo.Context) error {

	productoFinal := structs.ProductosFinales{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["ProductosFinales"], &productoFinal)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	productosFinalesService := models.ProductosFinalesService{
		ProductoFinal: &productoFinal,
		DbHandler:     pfc.DbHandler,
	}

	result, err := productosFinalesService.DarAlta(*token)

	if err != nil || result == nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error: nil,
	}

	response.AddModels(result)

	return c.JSON(http.StatusOK, response)
}

//DarBaja
func (pfc *ProductosFinalesController) DarBaja(c echo.Context) error {

	productoFinal := structs.ProductosFinales{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["ProductosFinales"], &productoFinal)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	productosFinalesService := models.ProductosFinalesService{
		ProductoFinal: &productoFinal,
		DbHandler:     pfc.DbHandler,
	}

	result, err := productosFinalesService.DarBaja(*token)

	if err != nil || result == nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error: nil,
	}

	response.AddModels(result)

	return c.JSON(http.StatusOK, response)
}

//Buscar
func (pfc *ProductosFinalesController) Buscar(c echo.Context) error {

	productoFinal := structs.ProductosFinales{}
	paginacion := structs.Paginaciones{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["ProductosFinales"], &productoFinal)
	mapstructure.Decode(jsonMap["Paginaciones"], &paginacion)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorProductosFinales := gestores.GestorProductosFinales{
		DbHandler: pfc.DbHandler,
	}

	result, err := gestorProductosFinales.Buscar(productoFinal, paginacion, *token)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error:     nil,
		Respuesta: result,
	}

	return c.JSON(http.StatusOK, response)
}

//Dame
func (pfc *ProductosFinalesController) Dame(c echo.Context) error {

	productoFinal := structs.ProductosFinales{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["ProductosFinales"], &productoFinal)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	productosFinalesService := models.ProductosFinalesService{
		ProductoFinal: &productoFinal,
		DbHandler:     pfc.DbHandler,
	}

	result, err := productosFinalesService.Dame(*token)

	if err != nil || result == nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error:     nil,
		Respuesta: result,
	}

	return c.JSON(http.StatusOK, response)
}

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

type GruposProductoController struct {
	DbHandler *db.DbHandler
}

/**
 * @api {POST} /gruposProducto/dame Dame Grupo de producto
 * @apiDescription Permite instanciar un grupo de producto a partir de su Id
 * @apiGroup GruposProducto
 * @apiHeader {String} Authorization
 * @apiParam {Object} GruposProducto
 * @apiParam {int} GruposProducto.IdGrupoProducto


  * @apiParamExample {json} Request-Example:
{
	 "GruposProducto": {
            "IdGrupoProducto":6
        }
 }
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":{
        "GruposProducto": {
			"IdGrupoProducto": 6,
			"Grupo": "Grupo de prueba 2",
			"FechaAlta": "2020-07-04 21:39:47.000000",
			"FechaBaja": "",
			"Descripcion": "",
            "Estado": "A",
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
func (gpc *GruposProductoController) Dame(c echo.Context) error {

	grupoProducto := structs.GruposProducto{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["GruposProducto"], &grupoProducto)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gruposProductoService := models.GruposProductoService{
		GrupoProducto: &grupoProducto,
		DbHandler:     gpc.DbHandler,
	}

	result, err := gruposProductoService.Dame(*token)

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
 * @api {POST} /gruposProducto/crear Crear Grupo de Producto
 * @apiDescription Permite crear un grupo de producto
 * @apiGroup GruposProducto
 * @apiHeader {String} Authorization
 * @apiParam {Object} GruposProducto
 * @apiParam {string} GruposProducto.Grupo
 * @apiParam {string} GruposProducto.Descripcion

 * @apiParamExample {json} Request-Example:
{
	"GruposProducto":{
        "Grupo":"Grupo1"
    }
}
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":{
		"GruposProducto":{
			"IdGrupoProducto": 3,
			"Grupo": "Grupo1",
			"FechaAlta": "2020-07-04 21:39:47.000000",
			"FechaBaja": "",
			"Descripcion": "",
			"Estado": "A"
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

func (gpc *GruposProductoController) Crear(c echo.Context) error {

	grupoProducto := structs.GruposProducto{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["GruposProducto"], &grupoProducto)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorGruposProducto := gestores.GestorGruposProducto{
		DbHandler: gpc.DbHandler,
	}
	result, err := gestorGruposProducto.Crear(grupoProducto, *token)

	if err != nil || result == nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error: nil,
	}

	response.AddModels(result)

	return c.JSON(http.StatusOK, response)
}

/**
 * @api {POST} /gruposProducto/modificar Modificar Grupo de Producto
 * @apiDescription Permite modificar un grupo de producto
 * @apiGroup GruposProducto
 * @apiHeader {String} Authorization
 * @apiParam {Object} GruposProducto
 * @apiParam {int} GruposProducto.IdGrupoProducto
 * @apiParam {string} GruposProducto.Grupo
 * @apiParam {string} GruposProducto.Descripcion

 * @apiParamExample {json} Request-Example:
{
	"GruposProducto":{
		"IdGrupoProducto":3,
        "Grupo":"Grupo1Modificado"
    }
}
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":{
		"GruposProducto":{
			"IdGrupoProducto": 3,
			"Grupo": "Grupo1Modificado",
			"FechaAlta": "2020-07-04 21:39:47.000000",
			"FechaBaja": "",
			"Descripcion": "",
			"Estado": "A"
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

func (gpc *GruposProductoController) Modificar(c echo.Context) error {

	grupoProducto := structs.GruposProducto{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["GruposProducto"], &grupoProducto)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorGruposProducto := gestores.GestorGruposProducto{
		DbHandler: gpc.DbHandler,
	}
	result, err := gestorGruposProducto.Modificar(grupoProducto, *token)

	if err != nil || result == nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error: nil,
	}

	response.AddModels(result)

	return c.JSON(http.StatusOK, response)
}

/**
 * @api {POST} /gruposProducto/borrar Borrar Grupo de Producto
 * @apiDescription Permite borrar un grupo de producto
 * @apiGroup GruposProducto
 * @apiHeader {String} Authorization
 * @apiParam {Object} GruposProducto
 * @apiParam {int} GruposProducto.IdGrupoProducto

 * @apiParamExample {json} Request-Example:
{
	"GruposProducto":{
		"IdGrupoProducto":3
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

func (gpc *GruposProductoController) Borrar(c echo.Context) error {

	grupoProducto := structs.GruposProducto{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["GruposProducto"], &grupoProducto)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorGruposProducto := gestores.GestorGruposProducto{
		DbHandler: gpc.DbHandler,
	}
	err = gestorGruposProducto.Borrar(grupoProducto, *token)

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
 * @api {POST} /gruposProducto/darAlta Dar alta GrupoProducto
 * @apiDescription Permite dar de alta un grupo de producto
 * @apiGroup GruposProducto
 * @apiHeader {String} Authorization
 * @apiParam {Object} GruposProducto
 * @apiParam {int} GruposProducto.IdGrupoProducto


  * @apiParamExample {json} Request-Example:
{
	 "GruposProducto": {
            "IdGrupoProducto":4,
        }
 }
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":{
		"GruposProducto":{
			"IdGrupoProducto": 5,
			"Grupo": "Grupo1",
			"FechaAlta": "2020-07-04 22:15:38.000000",
			"FechaBaja": "2020-07-04 22:15:44.000000",
			"Descripcion": "",
			"Estado": "A"
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
func (gpc *GruposProductoController) DarAlta(c echo.Context) error {

	grupoProducto := structs.GruposProducto{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["GruposProducto"], &grupoProducto)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gruposProductoService := models.GruposProductoService{
		DbHandler:     gpc.DbHandler,
		GrupoProducto: &grupoProducto,
	}

	result, err := gruposProductoService.DarAlta(*token)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error: nil,
	}

	response.AddModels(result)

	return c.JSON(http.StatusOK, response)
}

/**
 * @api {POST} /gruposProducto/darBaja Dar baja GrupoProducto
 * @apiDescription Permite dar de baja un grupo de producto
 * @apiGroup GruposProducto
 * @apiHeader {String} Authorization
 * @apiParam {Object} GruposProducto
 * @apiParam {int} GruposProducto.IdGrupoProducto


  * @apiParamExample {json} Request-Example:
{
	 "GruposProducto": {
            "IdGrupoProducto":4,
        }
 }
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":{
		"GruposProducto":{
			"IdGrupoProducto": 5,
			"Grupo": "Grupo1",
			"FechaAlta": "2020-07-04 22:15:38.000000",
			"FechaBaja": "2020-07-04 22:15:44.000000",
			"Descripcion": "",
			"Estado": "B"
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
func (gpc *GruposProductoController) DarBaja(c echo.Context) error {

	grupoProducto := structs.GruposProducto{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["GruposProducto"], &grupoProducto)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gruposProductoService := models.GruposProductoService{
		DbHandler:     gpc.DbHandler,
		GrupoProducto: &grupoProducto,
	}

	result, err := gruposProductoService.DarBaja(*token)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error: nil,
	}

	response.AddModels(result)

	return c.JSON(http.StatusOK, response)
}

/**
 * @api {POST} /gruposProducto Buscar Grupos de Producto
 * @apiDescription Permite buscar un grupo de producto
 * @apiGroup GruposProducto
 * @apiHeader {String} Authorization
 * @apiParam {Object} GruposProducto
 * @apiParam {string} GruposProducto.Grupo
 * @apiParam {string} GruposProducto.Estado

 * @apiParamExample {json} Request-Example:
{
	"GruposProducto":{
    }
}
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":{
		"GruposProducto":[
			{
				"IdGrupoProducto": 2,
				"Grupo": "Grupo de prueba 2",
				"FechaAlta": "2020-07-04 16:38:50.000000",
				"FechaBaja": "",
				"Descripcion": "",
				"Estado": "A"
			},
			{
				"IdGrupoProducto": 5,
				"Grupo": "Grupo1",
				"FechaAlta": "2020-07-04 22:15:38.000000",
				"FechaBaja": "2020-07-04 22:15:44.000000",
				"Descripcion": "",
				"Estado": "A"
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
//Crear

func (gpc *GruposProductoController) Buscar(c echo.Context) error {

	grupoProducto := structs.GruposProducto{}
	paginacion := structs.Paginaciones{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["GruposProducto"], &grupoProducto)
	mapstructure.Decode(jsonMap["Paginaciones"], &paginacion)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorGruposProducto := gestores.GestorGruposProducto{
		DbHandler: gpc.DbHandler,
	}
	result, err := gestorGruposProducto.Buscar(grupoProducto, paginacion, *token)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error:     nil,
		Respuesta: result,
	}

	return c.JSON(http.StatusOK, response)
}

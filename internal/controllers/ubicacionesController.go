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

type UbicacionesController struct {
	DbHandler *db.DbHandler
}

/**
 * @api {POST} /ubicaciones/crear Crear Ubicación
 * @apiDescription Permite crear una ubicación
 * @apiGroup Ubicaciones
 * @apiHeader {String} Authorization
 * @apiParam {Object} Ubicaciones
 * @apiParam {int} Ubicaciones.Ubicacion
 * @apiParam {int} Domicilios.Observaciones
 * @apiParam {Object} Domicilios
 * @apiParam {int} Domicilios.IdCiudad
 * @apiParam {int} Domicilios.IdProvincia
 * @apiParam {int} Domicilios.IdPais
 * @apiParam {string} Domicilios.Domicilio
 * @apiParam {string} Domicilios.CodigoPostal
 * @apiParam {string} Domicilios.Observaciones

  * @apiParamExample {json} Request-Example:
{
    "Ubicaciones":{
        "Ubicacion": "Prueba"
    },
    "Domicilios":{
        "IdCiudad":1,
        "IdProvincia":1,
        "IdPais":"AR",
        "Domicilio":"Domicilio de prueba",
        "CodigoPostal":"PRUEBA",
        "Observaciones":""
    }
}
 * @apiSuccessExample {json} Success-Response:
{
    "error": null,
    "respuesta": {
        "Ubicaciones": {
            "IdUbicacion": 5,
            "IdDomicilio": 19,
            "Ubicacion": "Prueba",
            "FechaAlta": "2020-06-13 13:38:59.000000",
            "Estado": "A"
        }
    }
}
* @apiErrorExample {json} Error-Response:
{
    "error": {
        "codigo": "ERROR_EXISTE_UBICACION",
        "mensaje": "La ubicacion ingresada ya existe."
    },
    "respuesta": null
}
*/

func (uc *UbicacionesController) Crear(c echo.Context) error {

	domicilios := structs.Domicilios{}
	ubicaciones := structs.Ubicaciones{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Domicilios"], &domicilios)
	mapstructure.Decode(jsonMap["Ubicaciones"], &ubicaciones)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorUbicaciones := gestores.GestorUbicaciones{
		DbHandler: uc.DbHandler,
	}
	result, err := gestorUbicaciones.Crear(ubicaciones, domicilios, *token)

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
 * @api {POST} /ubicaciones/borar Borrar Ubicación
 * @apiDescription Permite borrar una ubicación existente
 * @apiGroup Ubicaciones
 * @apiHeader {String} Authorization
 * @apiParam {Object} Ubicaciones
 * @apiParam {int} Ubicaciones.IdUbicacion
 * @apiParamExample {json} Request-Example:
{
	"Ubicaciones":{
		"IdUbicacion":5
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
        "codigo": "ERROR_SIN_PERMISOS",
        "mensaje": "No cuenta con los permisos para ejecutar esta accion."
    },
    "respuesta": null
}
*/
func (uc *UbicacionesController) Borrar(c echo.Context) error {

	ubicacion := structs.Ubicaciones{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Ubicaciones"], &ubicacion)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorUbicaciones := gestores.GestorUbicaciones{
		DbHandler: uc.DbHandler,
	}

	err = gestorUbicaciones.Borrar(ubicacion, *token)

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
 * @api {POST} /ubicaciones/modificar Modificar Ubicación
 * @apiDescription Permite modificar una ubicación existente
 * @apiGroup Ubicaciones
 * @apiHeader {String} Authorization
 * @apiParam {Object} Ubicaciones
 * @apiParam {int} Ubicaciones.IdUbicacion
 * @apiParamExample {json} Request-Example:
{
    "Ubicaciones":{
        "IdUbicacion": 8,
        "Ubicacion":"Modificar prueba"
    },
    "Domicilios":{
        "IdDomicilio":8,
        "IdCiudad":1,
        "IdProvincia":1,
        "IdPais":"AR",
        "Domicilio":"Domicilio de prueba",
        "CodigoPostal":"PRUEBA",
        "Observaciones":""
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
        "codigo": "ERROR_SIN_PERMISOS",
        "mensaje": "No cuenta con los permisos para ejecutar esta accion."
    },
    "respuesta": null
}
*/
func (uc *UbicacionesController) Modificar(c echo.Context) error {

	ubicacion := structs.Ubicaciones{}
	domicilio := structs.Domicilios{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Ubicaciones"], &ubicacion)
	mapstructure.Decode(jsonMap["Domicilios"], &domicilio)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorUbicaciones := gestores.GestorUbicaciones{
		DbHandler: uc.DbHandler,
	}

	result, err := gestorUbicaciones.Modificar(ubicacion, domicilio, *token)

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
 * @api {GET} /ubicaciones Listar Ubicaciones
 * @apiName Listar Ubicaciones
 * @apiDescription Devuelve una lista con la ubicación y su dirección.
 * @apiGroup Ubicaciones
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":[
		{
			"Ciudades":{
				"IdCiudad": 2,
				"IdProvincia": 2,
				"IdPais": "AR",
				"Ciudad": "Salta"
			},
				"Domicilios":{
				"IdDomicilio": 3,
				"IdCiudad": 2,
				"IdProvincia": 2,
				"IdPais": "AR",
				"Domicilio": "España 109",
				"CodigoPostal": "4400",
				"FechaAlta": "2020-06-13 13:52:16.000000",
				"Observaciones": "Domicilio sucursal Salta"
			},
			"Paises":{
				"IdPais": "AR",
				"Pais": "Argentina"
			},
			"Provincias":{
				"IdPais": "AR",
				"IdProvincia": 2,
				"Provincia": "Salta"
			},
			"Ubicaciones":{
				"IdUbicacion": 3,
				"IdDomicilio": 3,
				"Ubicacion": "Sucursal Salta",
				"FechaAlta": "2020-06-13 13:52:17.000000",
				"Estado": "A"
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
func (uc *UbicacionesController) Listar(c echo.Context) error {

	gestorUbicaciones := gestores.GestorUbicaciones{
		DbHandler: uc.DbHandler,
	}

	result, err := gestorUbicaciones.Listar()

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
 * @api {POST} /ubicaciones/darAlta Dar alta Ubicación
 * @apiDescription Permite dar de alta una ubicación
 * @apiGroup Ubicaciones
 * @apiHeader {String} Authorization
 * @apiParam {Object} Ubicaciones
 * @apiParam {int} Ubicaciones.IdUbicacion


  * @apiParamExample {json} Request-Example:
{
    "Ubicaciones":{
        "IdUbicacion": 8
    }
}
 * @apiSuccessExample {json} Success-Response:
{
    "error": null,
    "respuesta": {
        "Ubicaciones": {
            "IdUbicacion": 8,
            "IdDomicilio": 8,
            "Ubicacion": "Modificar prueba",
            "FechaAlta": "2020-06-13 15:43:20.000000",
            "FechaBaja": "2020-06-13 15:42:10.000000",
            "Estado": "A"
        }
    }
}
* @apiErrorExample {json} Error-Response:
{
    "error": {
        "codigo": "ERROR_UBICACION_ESTA_ALTA",
        "mensaje": "La ubicación no existe o ya está en estado de 'Alta'."
    },
    "respuesta": null
}
*/
func (uc *UbicacionesController) DarAlta(c echo.Context) error {

	ubicacion := structs.Ubicaciones{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Ubicaciones"], &ubicacion)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	ubicacionesService := models.UbicacionesService{
		DbHandler:   uc.DbHandler,
		Ubicaciones: &ubicacion,
	}

	result, err := ubicacionesService.DarAlta(*token)

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
 * @api {POST} /ubicaciones/dame Dame Ubicación
 * @apiDescription Permite instanciar una ubicación a partir de su Id
 * @apiGroup Ubicaciones
 * @apiHeader {String} Authorization
 * @apiParam {Object} Ubicaciones
 * @apiParam {int} Ubicaciones.IdUbicacion


  * @apiParamExample {json} Request-Example:
{
    "Ubicaciones":{
        "IdUbicacion": 8
    }
}
 * @apiSuccessExample {json} Success-Response:
{
    "error": null,
    "respuesta": {
        "Ubicaciones": {
            "IdUbicacion": 8,
            "IdDomicilio": 8,
            "Ubicacion": "Modificar prueba",
            "FechaAlta": "2020-06-13 15:43:20.000000",
            "FechaBaja": "2020-06-13 15:42:10.000000",
            "Estado": "A"
        }
    }
}
* @apiErrorExample {json} Error-Response:
{
    "error": {
        "codigo": "ERROR_UBICACION_ESTA_ALTA",
        "mensaje": "La ubicación no existe o ya está en estado de 'Alta'."
    },
    "respuesta": null
}
*/
func (uc *UbicacionesController) Dame(c echo.Context) error {

	ubicacion := structs.Ubicaciones{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Ubicaciones"], &ubicacion)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	ubicacionesService := models.UbicacionesService{
		DbHandler:   uc.DbHandler,
		Ubicaciones: &ubicacion,
	}

	result, err := ubicacionesService.Dame(*token)

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
 * @api {POST} /ubicaciones/darBaja Dar baja Ubicación
 * @apiDescription Permite dar de baja una ubicación
 * @apiGroup Ubicaciones
 * @apiHeader {String} Authorization
 * @apiParam {Object} Ubicaciones
 * @apiParam {int} Ubicaciones.IdUbicacion
 * @apiParamExample {json} Request-Example:
{
	 "Ubicaciones": {
            "IdUbicacion":8,
        }
 }
 * @apiSuccessExample {json} Success-Response:
{
    "error": null,
    "respuesta": {
        "Ubicaciones": {
            "IdUbicacion": 8,
            "IdDomicilio": 8,
            "Ubicacion": "Modificar prueba",
            "FechaAlta": "2020-06-13 14:55:33.000000",
            "FechaBaja": "2020-06-13 15:42:10.000000",
            "Estado": "B"
        }
    }
}
* @apiErrorExample {json} Error-Response:
{
    "error": {
        "codigo": "ERROR_UBICACION_ESTA_BAJA",
        "mensaje": "La ubicación no existe o ya está en estado de 'Baja'."
    },
    "respuesta": null
}
*/
func (uc *UbicacionesController) DarBaja(c echo.Context) error {

	ubicacion := structs.Ubicaciones{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Ubicaciones"], &ubicacion)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	ubicacionesService := models.UbicacionesService{
		DbHandler:   uc.DbHandler,
		Ubicaciones: &ubicacion,
	}

	result, err := ubicacionesService.DarBaja(*token)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error: nil,
	}

	response.AddModels(result)

	return c.JSON(http.StatusOK, response)
}

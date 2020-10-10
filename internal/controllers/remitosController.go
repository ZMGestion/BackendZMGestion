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
 * @apiParam {int} Remitos.IdDomicilio
 * @apiParam {int} Remitos.IdUbicacion
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

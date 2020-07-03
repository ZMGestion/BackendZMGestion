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

type TelasController struct {
	DbHandler *db.DbHandler
}

/**
 * @api {POST} /telas/crear Crear Tela
 * @apiDescription Permite crear una tela
 * @apiGroup Telas
 * @apiHeader {String} Authorization
 * @apiParam {Object} Telas
 * @apiParam {int} Telas.Tela
 * @apiParam {int} Telas.Observaciones
 * @apiParam {Object} Precios
 * @apiParam {int} Precios.Precio

  * @apiParamExample {json} Request-Example:
{
    "Telas":{
        "Tela": "Prueba3"
    },
    "Precios":{
        "Precio":1.20
    }
}
 * @apiSuccessExample {json} Success-Response:
{
    "error": null,
    "respuesta": {
        "Telas": {
            "IdTela": 4,
            "Tela": "Prueba5",
            "FechaAlta": "2020-06-30 23:39:57.000000",
            "FechaBaja": "",
            "Observaciones": "",
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

func (tc *TelasController) Crear(c echo.Context) error {

	telas := structs.Telas{}
	precios := structs.Precios{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Telas"], &telas)
	mapstructure.Decode(jsonMap["Precios"], &precios)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorTelas := gestores.GestorTelas{
		DbHandler: tc.DbHandler,
	}
	result, err := gestorTelas.Crear(telas, precios, *token)

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
 * @api {POST} /telas/modificar Modificar Tela
 * @apiDescription Permite modificar una tela
 * @apiGroup Telas
 * @apiHeader {String} Authorization
 * @apiParam {Object} Telas
 * @apiParam {int} Telas.IdTela
 * @apiParam {int} Telas.Tela
 * @apiParam {int} Telas.Observaciones

  * @apiParamExample {json} Request-Example:
{
    "Telas":{
		"IdTela":1
        "Tela": "Prueba13"
    }
}
 * @apiSuccessExample {json} Success-Response:
{
    "error": null,
    "respuesta": {
        "Telas": {
            "IdTela": 4,
            "Tela": "Prueba5",
            "FechaAlta": "2020-06-30 23:39:57.000000",
            "FechaBaja": "",
            "Observaciones": "",
            "Estado": "A"
        }
    }
}
* @apiErrorExample {json} Error-Response:
{
    "error": {
        "codigo": "ERROR_DEFAULT",
        "mensaje": "Ha ocurrido un error mientras se procesaba su peticion."
    },
    "respuesta": null
}
*/

func (tc *TelasController) Modificar(c echo.Context) error {

	telas := structs.Telas{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Telas"], &telas)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorTelas := gestores.GestorTelas{
		DbHandler: tc.DbHandler,
	}
	result, err := gestorTelas.Modificar(telas, *token)

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
 * @api {POST} /telas/darAlta Dar alta tela
 * @apiDescription Permite dar de alta una tela
 * @apiGroup Telas
 * @apiHeader {String} Authorization
 * @apiParam {Object} Telas
 * @apiParam {int} Telas.IdTela


  * @apiParamExample {json} Request-Example:
{
	 "Telas": {
            "IdTela":3,
        }
 }
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":{
		"Telas":{
			"IdTela": 1,
			"Tela": "Prueba65",
			"FechaAlta": "2020-06-29 22:56:45.000000",
			"FechaBaja": "2020-07-03 12:16:53.000000",
			"Observaciones": "",
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
func (tc *TelasController) DarAlta(c echo.Context) error {

	tela := structs.Telas{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Telas"], &tela)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	telasService := models.TelasService{
		DbHanlder: tc.DbHandler,
		Tela:      &tela,
	}

	result, err := telasService.DarAlta(*token)

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
 * @api {POST} /telas/darBaja Dar baja tela
 * @apiDescription Permite dar de baja una tela
 * @apiGroup Telas
 * @apiHeader {String} Authorization
 * @apiParam {Object} Telas
 * @apiParam {int} Telas.IdTela


  * @apiParamExample {json} Request-Example:
{
	 "Telas": {
            "IdTela":3,
        }
 }
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":{
		"Telas":{
			"IdTela": 1,
			"Tela": "Prueba65",
			"FechaAlta": "2020-06-29 22:56:45.000000",
			"FechaBaja": "2020-07-03 12:16:53.000000",
			"Observaciones": "",
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
func (tc *TelasController) DarBaja(c echo.Context) error {

	tela := structs.Telas{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Telas"], &tela)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	telasService := models.TelasService{
		DbHanlder: tc.DbHandler,
		Tela:      &tela,
	}

	result, err := telasService.DarBaja(*token)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error: nil,
	}

	response.AddModels(result)

	return c.JSON(http.StatusOK, response)
}

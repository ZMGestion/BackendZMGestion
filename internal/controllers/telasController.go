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
 * @apiParam {string} Telas.Tela
 * @apiParam {string} Telas.Observaciones
 * @apiParam {Object} Precios
 * @apiParam {double} Precios.Precio

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
        "codigo": "ERROR_DEFAULT",
        "mensaje": "Ha ocurrido un error mientras se procesaba su petición."
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
		"IdTela":1,
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
 * @api {POST} /telas/darAlta Dar alta Tela
 * @apiDescription Permite dar de alta una tela
 * @apiGroup Telas
 * @apiHeader {String} Authorization
 * @apiParam {Object} Telas
 * @apiParam {int} Telas.IdTela


  * @apiParamExample {json} Request-Example:
{
	 "Telas": {
            "IdTela":3
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
 * @api {POST} /telas/darBaja Dar baja Tela
 * @apiDescription Permite dar de baja una tela
 * @apiGroup Telas
 * @apiHeader {String} Authorization
 * @apiParam {Object} Telas
 * @apiParam {int} Telas.IdTela


  * @apiParamExample {json} Request-Example:
{
	 "Telas": {
            "IdTela":3
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

/**
 * @api {POST} /telas/borar Borrar Tela
 * @apiDescription Permite borrar una tela
 * @apiGroup Telas
 * @apiHeader {String} Authorization
 * @apiParam {Object} Telas
 * @apiParam {int} Telas.IdTela


  * @apiParamExample {json} Request-Example:
{
	 "Telas": {
            "IdTela":2
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
func (tc *TelasController) Borrar(c echo.Context) error {

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

	gestorTelas := gestores.GestorTelas{
		DbHandler: tc.DbHandler,
	}
	err = gestorTelas.Borrar(tela, *token)

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
 * @api {POST} /telas/precios/modificar Modificar Precio Tela
 * @apiDescription Permite modificar el precio de una tela
 * @apiGroup Telas
 * @apiHeader {String} Authorization
 * @apiParam {Object} Telas
 * @apiParam {int} Telas.IdTela
 * @apiParam {Object} Precios
 * @apiParam {int} Precios.IdPrecio
 * @apiParamExample {json} Request-Example:
{
	"Telas": {
		"IdTela":3
	},
	"Precios":{
		"Precio":1.21
	}
 }
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":{
		"Precios":{
			"IdPrecio": 17,
			"Precio": 1.22,
			"Tipo": "",
			"IdReferencia": 0,
			"FechaAlta": "2020-07-03 20:19:15.000000"
		},
		"Telas":{
			"IdTela": 5,
			"Tela": "Prueba5",
			"FechaAlta": "2020-07-03 19:57:18.000000",
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
        "mensaje": "Ha ocurrido un error mientras se procesaba su petición."
    },
    "respuesta": null
}
*/
func (tc *TelasController) ModificarPrecio(c echo.Context) error {

	tela := structs.Telas{}
	precio := structs.Precios{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Telas"], &tela)
	mapstructure.Decode(jsonMap["Precios"], &precio)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	telasService := models.TelasService{
		DbHanlder: tc.DbHandler,
		Tela:      &tela,
	}

	result, err := telasService.ModificarPrecio(precio, *token)

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
 * @api {POST} /telas/precios Listar Precios Tela
 * @apiDescription Permite listar el historico de precios de una tela
 * @apiGroup Telas
 * @apiHeader {String} Authorization
 * @apiParam {Object} Telas
 * @apiParam {int} Telas.IdTela


  * @apiParamExample {json} Request-Example:
{
	 "Telas": {
            "IdTela":3
        }
 }
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":{
		"Precios":[
			{"IdPrecio": 14, "Precio": 1.2, "Tipo": "", "IdReferencia": 0, "FechaAlta": "2020-07-03 19:57:18.000000"…},
			{"IdPrecio": 16, "Precio": 1.21, "Tipo": "", "IdReferencia": 0, "FechaAlta": "2020-07-03 20:15:10.000000"…},
			{"IdPrecio": 17, "Precio": 1.22, "Tipo": "", "IdReferencia": 0, "FechaAlta": "2020-07-03 20:19:15.000000"…},
			{"IdPrecio": 18, "Precio": 1.23, "Tipo": "", "IdReferencia": 0, "FechaAlta": "2020-07-03 22:29:53.000000"…}
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
func (tc *TelasController) ListarPrecios(c echo.Context) error {

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

	result, err := telasService.ListarPrecios(*token)

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
 * @api {POST} /telas Buscar Telas
 * @apiDescription Permite buscar una tela
 * @apiGroup Telas
 * @apiHeader {String} Authorization
 * @apiParam {Object} Telas
 * @apiParam {string} Telas.Tela
 * @apiParam {string} Telas.Estado

  * @apiParamExample {json} Request-Example:
{
    "Telas":{
		"IdTela":1,
        "Tela": "Prueba13"
    }
}
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":[
		{
			"Precios":{
				"IdPrecio": 29,
				"Precio": 1.2,
				"Tipo": "",
				"IdReferencia": 0,
				"FechaAlta": ""
			},
			"Telas":{
				"IdTela": 5,
				"Tela": "Prueba5",
				"FechaAlta": "2020-07-03 19:57:18.000000",
				"FechaBaja": "",
				"Observaciones": "",
				"Estado": "A"
			}
		}
	]
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
//Buscar
func (tc *TelasController) Buscar(c echo.Context) error {

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
	result, err := gestorTelas.Buscar(telas, *token)

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
 * @api {POST} /telas/dame Dame Tela
 * @apiDescription Permite instanciar una tela a partir de su Id
 * @apiGroup Telas
 * @apiHeader {String} Authorization
 * @apiParam {Object} Telas
 * @apiParam {int} Telas.IdTela

  * @apiParamExample {json} Request-Example:
{
    "Telas":{
		"IdTela":5
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
func (tc *TelasController) Dame(c echo.Context) error {

	tela := structs.Telas{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Telas"], &tela)

	//headerToken := c.Request().Header.Get("Authorization")
	//token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	telasService := models.TelasService{
		DbHanlder: tc.DbHandler,
		Tela:      &tela,
	}

	result, err := telasService.Dame()

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error:     nil,
		Respuesta: result,
	}

	return c.JSON(http.StatusOK, response)
}

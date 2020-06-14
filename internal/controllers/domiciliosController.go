package controllers

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/gestores"
	"BackendZMGestion/internal/helpers"
	"BackendZMGestion/internal/interfaces"
	"BackendZMGestion/internal/structs"
	"net/http"

	"github.com/labstack/echo"
	"github.com/mitchellh/mapstructure"
)

//DomiciliosController
type DomiciliosController struct {
	DbHandler *db.DbHandler
}

/**
 * @api {POST} /domicilios/crear Crear Domicilio
 * @apiDescription Permite crear un domicilio para un cliente
 * @apiGroup Domicilios
 * @apiHeader {String} Authorization
 * @apiParam {Object} Domicilios
 * @apiParam {int} Domicilios.IdCiudad
 * @apiParam {int} Domicilios.IdProvincia
 * @apiParam {int} Domicilios.IdPais
 * @apiParam {string} Domicilios.Domicilio
 * @apiParam {string} Domicilio.CodigoPostal
 * @apiParam {string} Domicilios.Observaciones
 * @apiParam {Object} Clientes
 * @apiParam {int} Clientes.IdCliente


  * @apiParamExample {json} Request-Example:
{
    "error": {
        "codigo": "ERROR_NOEXISTE_CLIENTE",
        "mensaje": "No existe el cliente ingresado."
    },
    "respuesta": null
}
 * @apiSuccessExample {json} Success-Response:
{

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

func (dc *DomiciliosController) Crear(c echo.Context) error {

	domicilios := structs.Domicilios{}
	clientes := structs.Clientes{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Domicilios"], &domicilios)
	mapstructure.Decode(jsonMap["Clientes"], &clientes)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorDomicilios := gestores.GestorDomicilios{
		DbHanlder: dc.DbHandler,
	}
	result, err := gestorDomicilios.Crear(domicilios, clientes, *token)

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
 * @api {POST} /domicilios/borar Borrar Domicilio
 * @apiDescription Permite borrar un domicilio de un cliente
 * @apiGroup Domicilios
 * @apiHeader {String} Authorization
 * @apiParam {Object} Domicilios
 * @apiParam {int} Domicilios.IdDomicilio
 * @apiParam {Object} Clientes
 * @apiParam {int} Clientes.IdCliente



  * @apiParamExample {json} Request-Example:
{
	 "Domicilios":{
		 "IdDomicilio":1
	 },
	 "Cliente":{
		 "IdCliente":1
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
func (dc *DomiciliosController) Borrar(c echo.Context) error {

	domicilio := structs.Domicilios{}
	cliente := structs.Clientes{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Domicilios"], &domicilio)
	mapstructure.Decode(jsonMap["Clientes"], &cliente)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorDomicilio := gestores.GestorDomicilios{
		DbHanlder: dc.DbHandler,
	}

	err = gestorDomicilio.Borrar(domicilio, cliente, *token)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error:     nil,
		Respuesta: nil,
	}

	return c.JSON(http.StatusOK, response)
}

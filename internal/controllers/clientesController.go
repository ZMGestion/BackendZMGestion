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

type ClientesController struct {
	DbHanlder *db.DbHandler
}

/**
 * @api {POST} /clientes/crear Crear Cliente
 * @apiDescription Permite crear un cliente
 * @apiGroup Clientes
 * @apiHeader {String} Authorization
 * @apiParam {Object} Clientes
 * @apiParam {string} Clientes.IdPais
 * @apiParam {int} Clientes.IdTipoDocumento
 * @apiParam {string} Clientes.Documento
 * @apiParam {string} Clientes.Nombres
 * @apiParam {string} Clientes.Apellidos
 * @apiParam {string} Clientes.RazonSocial
 * @apiParam {string} Clientes.Tipo
 * @apiParam {string} Clientes.Email
 * @apiParam {string} Clientes.Telefono
 * @apiParam {string} Clientes.FechaNacimiento
 * @apiParam {Object} Domicilios
 * @apiParam {string} Domicilios.Domicilio
 * @apiParam {string} Domicilios.CodigoPostal
 * @apiParam {string} Domicilios.IdPais
 * @apiParam {int} Domicilios.IdProvincia
 * @apiParam {int} Domicilios.IdCiudad
 * @apiParamExample {json} Request-Example:
{  "Clientes":{
        "IdPais": "AR",
        "IdTipoDocumento": 1,
        "Documento": "41144069",
        "Tipo":"F",
        "FechaNacimiento":"1998-05-27",
        "Nombres":"Loik",
        "Apellidos":"Choua",
        "Email":"loikchoua4@gmail.com",
        "Telefono":"3815483777"
    },
    "Domicilios":{
        "IdCiudad":1,
        "IdProvincia":1,
        "IdPais":"AR",
        "Domicilio":"Domicilio",
        "CodigoPostal":"4000"
    }
}
 * @apiSuccessExample {json} Success-Response:
{
    "error": null,
    "respuesta": {
        "Clientes": {
            "IdCliente": 3,
            "IdPais": "AR",
            "IdTipoDocumento": 1,
            "Documento": "41144069",
            "Tipo": "F",
            "FechaNacimiento": "1998-05-27",
            "Nombres": "Loik",
            "Apellidos": "Choua",
            "RazonSocial": "",
            "Email": "loikchoua4@gmail.com",
            "Telefono": "3815483777",
            "FechaAlta": "2020-06-24 15:32:47.000000",
            "FechaBaja": "",
            "Estado": "Choua"
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

func (cc *ClientesController) Crear(c echo.Context) error {

	cliente := structs.Clientes{}
	domicilio := structs.Domicilios{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Clientes"], &cliente)
	mapstructure.Decode(jsonMap["Domicilios"], &domicilio)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorCliente := gestores.GestorClientes{
		DbHandler: cc.DbHanlder,
	}
	result, err := gestorCliente.Crear(cliente, domicilio, *token)

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
 * @api {POST} /clientes/modificar Modificar Cliente
 * @apiDescription Permite modificar un cliente
 * @apiGroup Cliente
 * @apiHeader {String} Authorization
 * @apiParam {Object} Clientes
 * @apiParam {string} Clientes.IdPais
 * @apiParam {int} Clientes.IdTipoDocumento
 * @apiParam {string} Clientes.Documento
 * @apiParam {string} Clientes.Nombres
 * @apiParam {string} Clientes.Apellidos
 * @apiParam {string} Clientes.RazonSocial
 * @apiParam {string} Clientes.Tipo
 * @apiParam {string} Clientes.Email
 * @apiParam {string} Clientes.Telefono
 * @apiParam {string} Clientes.FechaNacimiento
 * @apiParamExample {json} Request-Example:
{
    "Clientes":{
        "IdPais": "AR",
        "IdTipoDocumento": 1,
        "Documento": "41144069",
        "Tipo":"F",
        "FechaNacimiento":"1998-05-27",
        "Nombres":"Loik",
        "Apellidos":"Choua",
        "Email":"loikchoua4@gmail.com",
        "Telefono":"3815483777"
    }
}
 * @apiSuccessExample {json} Success-Response:
{
    "error": null,
    "respuesta": {
        "Clientes": {
            "IdCliente": 3,
            "IdPais": "AR",
            "IdTipoDocumento": 1,
            "Documento": "41144069",
            "Tipo": "F",
            "FechaNacimiento": "1998-05-27",
            "Nombres": "Loik",
            "Apellidos": "Choua",
            "RazonSocial": "",
            "Email": "loikchoua4@gmail.com",
            "Telefono": "+543815483777",
            "FechaAlta": "2020-06-24 15:32:47.000000",
            "FechaBaja": "",
            "Estado": "Choua"
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
func (cc *ClientesController) Modificar(c echo.Context) error {

	cliente := structs.Clientes{}
	domicilio := structs.Domicilios{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Clientes"], &cliente)
	mapstructure.Decode(jsonMap["Domicilios"], &domicilio)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorCliente := gestores.GestorClientes{
		DbHandler: cc.DbHanlder,
	}
	result, err := gestorCliente.Modificar(cliente, *token)

	if err != nil || result == nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error: nil,
	}

	response.AddModels(result)

	return c.JSON(http.StatusOK, response)
}

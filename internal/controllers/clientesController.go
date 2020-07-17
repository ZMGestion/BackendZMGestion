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

type ClientesController struct {
	DbHanlder *db.DbHandler
}

/**
 * @api {POST} /clientes/dame Dame Cliente
 * @apiDescription Permite instanciar un cliente a partir de su Id
 * @apiGroup Clientes
 * @apiHeader {String} Authorization
 * @apiParam {Object} Clientes
 * @apiParam {int} Clientes.IdCliente
 * @apiParamExample {json} Request-Example:
{
    "Clientes":{
        "IdCliente": 3
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
            "FechaNacimiento": "",
            "Nombres": "Loik",
            "Apellidos": "Choua",
            "RazonSocial": "",
            "Email": "loikchoua4@gmail.com",
            "Telefono": "+543815483777",
            "FechaAlta": "2020-06-24 15:32:47.000000",
            "FechaBaja": "",
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
func (cc *ClientesController) Dame(c echo.Context) error {

	cliente := structs.Clientes{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Clientes"], &cliente)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	clientesService := models.ClientesService{
		DbHanlder: cc.DbHanlder,
		Cliente:   &cliente,
	}
	result, err := clientesService.Dame(*token)

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
 * @api {POST} /clientes/crear Crear Cliente
 * @apiDescription Permite crear un cliente
 * @apiGroup Clientes
 * @apiHeader {String} Authorization
 * @apiParam {Object} Clientes
 * @apiParam {string} Clientes.IdPais
 * @apiParam {int} Clientes.IdTipoDocumento
 * @apiParam {string} Clientes.Documento
 * @apiParam {string} [Clientes.Nombres] Si Tipo = 'F', es Obligatorio
 * @apiParam {string} [Clientes.Apellidos] Si Tipo = 'F', es Obligatorio
 * @apiParam {string} [Clientes.RazonSocial] Si Tipo = 'J', es Obligatorio
 * @apiParam {string} Clientes.Tipo
 * @apiParam {string} Clientes.Email
 * @apiParam {string} Clientes.Telefono
 * @apiParam {string} Clientes.FechaNacimiento
 * @apiParam {Object} [Domicilios]
 * @apiParam {string} [Domicilios.Domicilio]
 * @apiParam {string} [Domicilios.CodigoPostal]
 * @apiParam {string} [Domicilios.IdPais]
 * @apiParam {int} [Domicilios.IdProvincia]
 * @apiParam {int} [Domicilios.IdCiudad]
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
 * @apiGroup Clientes
 * @apiHeader {String} Authorization
 * @apiParam {Object} Clientes
 * @apiParam {string} Clientes.IdPais
 * @apiParam {int} Clientes.IdTipoDocumento
 * @apiParam {string} Clientes.Documento
 * @apiParam {string} [Clientes.Nombres] Si Tipo = 'F', es Obligatorio
 * @apiParam {string} [Clientes.Apellidos] Si Tipo = 'F', es Obligatorio
 * @apiParam {string} [Clientes.RazonSocial] Si Tipo = 'J', es Obligatorio
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

/**
 * @api {POST} /clientes/darAlta Dar alta Cliente
 * @apiDescription Permite dar de alta un cliente
 * @apiGroup Clientes
 * @apiHeader {String} Authorization
 * @apiParam {Object} Clientes
 * @apiParam {int} Clientes.IdCliente
 * @apiParamExample {json} Request-Example:
{
    "Clientes":{
        "IdCliente": 3
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
            "FechaNacimiento": "",
            "Nombres": "Loik",
            "Apellidos": "Choua",
            "RazonSocial": "",
            "Email": "loikchoua4@gmail.com",
            "Telefono": "+543815483777",
            "FechaAlta": "2020-06-24 15:32:47.000000",
            "FechaBaja": "",
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
func (cc *ClientesController) DarAlta(c echo.Context) error {

	cliente := structs.Clientes{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Clientes"], &cliente)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	clientesService := models.ClientesService{
		DbHanlder: cc.DbHanlder,
		Cliente:   &cliente,
	}
	result, err := clientesService.DarAlta(*token)

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
 * @api {POST} /clientes/darBaja Dar baja Cliente
 * @apiDescription Permite dar de baja un cliente
 * @apiGroup Clientes
 * @apiHeader {String} Authorization
 * @apiParam {Object} Clientes
 * @apiParam {int} Clientes.IdCliente
 * @apiParamExample {json} Request-Example:
{
    "Clientes":{
        "IdCliente": 3
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
            "FechaNacimiento": "",
            "Nombres": "Loik",
            "Apellidos": "Choua",
            "RazonSocial": "",
            "Email": "loikchoua4@gmail.com",
            "Telefono": "+543815483777",
            "FechaAlta": "2020-06-24 15:32:47.000000",
            "FechaBaja": "",
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
func (cc *ClientesController) DarBaja(c echo.Context) error {

	cliente := structs.Clientes{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Clientes"], &cliente)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	clientesService := models.ClientesService{
		DbHanlder: cc.DbHanlder,
		Cliente:   &cliente,
	}
	result, err := clientesService.DarBaja(*token)

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
 * @api {POST} /clientes/borrar Borrar Cliente
 * @apiDescription Permite borrar un cliente
 * @apiGroup Clientes
 * @apiHeader {String} Authorization
 * @apiParam {Object} Clientes
 * @apiParam {int} Clientes.IdCliente
 * @apiParamExample {json} Request-Example:
{
    "Clientes":{
        "IdCliente": 3
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
func (cc *ClientesController) Borrar(c echo.Context) error {

	cliente := structs.Clientes{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Clientes"], &cliente)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorClientes := gestores.GestorClientes{
		DbHandler: cc.DbHanlder,
	}
	err = gestorClientes.Borrar(cliente, *token)

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
 * @api {POST} /clientes/domicilios Listar Domicilios
 * @apiDescription Permite listar los domicilios de un cliente
 * @apiGroup Clientes
 * @apiParam {Object} Clientes
 * @apiParam {int} Clientes.IdCliente
  * @apiParamExample {json} Request-Example:
{
	"Clientes":{
		"IdCliente":1
	}
}
 * @apiSuccessExample {json} Success-Response:
{
    "error": null,
    "respuesta": {
        "Domicilios": [
            {
                "IdDomicilio": 14,
                "IdCiudad": 1,
                "IdProvincia": 1,
                "IdPais": "AR",
                "Domicilio": "El Tipal Lote 13",
                "CodigoPostal": "El Tipal Lote 13",
                "FechaAlta": "2020-06-28 23:04:26.000000",
                "Observaciones": ""
            },
            {
                "IdDomicilio": 15,
                "IdCiudad": 1,
                "IdProvincia": 1,
                "IdPais": "AR",
                "Domicilio": "El Tipal Lote 14",
                "CodigoPostal": "El Tipal Lote 14",
                "FechaAlta": "2020-06-28 23:04:49.000000",
                "Observaciones": ""
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

func (cc *ClientesController) ListarDomicilios(c echo.Context) error {

	clientes := structs.Clientes{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Clientes"], &clientes)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	clientesService := models.ClientesService{
		DbHanlder: cc.DbHanlder,
		Cliente:   &clientes,
	}
	result, err := clientesService.ListarDomicilios()

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
 * @api {POST} /clientes/domicilios/agregar Agregar Domicilio
 * @apiDescription Permite crear un domicilio para un cliente
 * @apiGroup Clientes
 * @apiHeader {String} Authorization
 * @apiParam {Object} Domicilios
 * @apiParam {int} Domicilios.IdCiudad
 * @apiParam {int} Domicilios.IdProvincia
 * @apiParam {int} Domicilios.IdPais
 * @apiParam {string} Domicilios.Domicilio
 * @apiParam {string} Domicilio.CodigoPostal
 * @apiParam {string} [Domicilios.Observaciones]
 * @apiParam {Object} Clientes
 * @apiParam {int} Clientes.IdCliente
  * @apiParamExample {json} Request-Example:
{
	"Clientes":{
		"IdCliente":1
	},
	"Domicilios":{
		"Domicilio":"El Tipal Lote 13",
		"IdCiudad":1,
		"IdProvincia":1,
		"IdPais":"AR",
		"CodigoPostal":"4107"
	}
}
 * @apiSuccessExample {json} Success-Response:
{
    "error": null,
    "respuesta": {
        "Domicilios": {
            "IdDomicilio": 14,
            "IdCiudad": 1,
            "IdProvincia": 1,
            "IdPais": "AR",
            "Domicilio": "El Tipal Lote 13",
            "CodigoPostal": "4107",
            "FechaAlta": "2020-06-28 23:04:26.000000",
            "Observaciones": ""
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

func (cc *ClientesController) AgregarDomicilio(c echo.Context) error {

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
		DbHanlder: cc.DbHanlder,
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
 * @api {POST} /clientes/domicilios/quitar Quitar Domicilio
 * @apiDescription Permite borrar un domicilio de un cliente
 * @apiGroup Clientes
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
	 "Clientes":{
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
func (cc *ClientesController) QuitarDomicilio(c echo.Context) error {

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
		DbHanlder: cc.DbHanlder,
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

/**
 * @api {POST} /clientes Buscar Clientes
 * @apiDescription Permite buscar clientes
 * @apiGroup Clientes
 * @apiHeader {String} Authorization
 * @apiParam {Object} Clientes
 * @apiParam {string} [Clientes.IdPais]
 * @apiParam {string} [Clientes.Documento]
 * @apiParam {string} [Clientes.Nombres]
 * @apiParam {string} [Clientes.Apellidos]
 * @apiParam {string} [Clientes.RazonSocial]
 * @apiParam {string} [Clientes.Tipo]
 * @apiParam {string} [Clientes.Email]
 * @apiParam {string} [Clientes.Telefono]
 * @apiParam {Object} Paginaciones
 * @apiParam {int} [Paginaciones.Pagina]
 * @apiParam {int} [Paginaciones.LongitudPagina]
 * @apiParamExample {json} Request-Example:
{
    "Clientes":{
        "IdPais": "AR",
        "Documento": "41144069",
        "Tipo":"F",
        "Nombres":"Loik",
        "Apellidos":"Choua",
        "Email":"loikchoua4@gmail.com",
        "Telefono":"3815483777"
    },
    "Paginaciones":{
      "Pagina":1,
      "LongitudPagina":1
    }
}
 * @apiSuccessExample {json} Success-Response:
{
    "error": null,
    "respuesta":{
        "resultado":[
            {
                "Clientes":{
                    "Apellidos": "Choua",
                    "Documento": "41144069",
                    "Email": "loikchoua4@gmail.com",
                    "Estado": "A",
                    "FechaAlta": "2020-06-24 15:32:47.000000",
                    "FechaBaja": null,
                    "FechaNacimiento": "1998-05-27",
                    "IdCliente": 3,
                    "IdPais": "AR",
                    "IdTipoDocumento": 1,
                    "Nombres": "Loik",
                    "RazonSocial": null,
                    "Telefono": "+543815483777",
                    "Tipo": "F"
                }
            }
        ],
        "Paginaciones":{
            "Pagina": 1,
            "LongitudPagina": 1,
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
func (cc *ClientesController) Buscar(c echo.Context) error {

	cliente := structs.Clientes{}
	paginacion := structs.Paginaciones{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Clientes"], &cliente)
	mapstructure.Decode(jsonMap["Paginaciones"], &paginacion)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorCliente := gestores.GestorClientes{
		DbHandler: cc.DbHanlder,
	}
	result, err := gestorCliente.Buscar(cliente, paginacion, *token)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error:     nil,
		Respuesta: result,
	}

	return c.JSON(http.StatusOK, response)
}

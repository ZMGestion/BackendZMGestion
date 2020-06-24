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
 * @api {POST} /usuarios/crear Crear Usuario
 * @apiDescription Permite crear un usuario
 * @apiGroup Usuarios
 * @apiHeader {String} Authorization
 * @apiParam {Object} Usuarios
 * @apiParam {int} Usuarios.IdRol
 * @apiParam {int} Usuarios.IdUbicacion
 * @apiParam {int} Usuarios.IdTipoDocumento
 * @apiParam {string} Usuarios.Documento
 * @apiParam {string} Usuarios.Nombres
 * @apiParam {string} Usuarios.Apellidos
 * @apiParam {string} Usuarios.EstadoCivil
 * @apiParam {string} Usuarios.Telefono
 * @apiParam {string} Usuarios.Email
 * @apiParam {int} Usuarios.CantidadHijos
 * @apiParam {string} Usuarios.Usuario
 * @apiParam {string} Usuarios.Password
 * @apiParam {string} Usuarios.FechaNacimiento
 * @apiParam {string} Usuarios.FechaInicio

  * @apiParamExample {json} Request-Example:
{
	 "Usuarios": {
            "IdRol": 1,
            "IdUbicacion": 1,
			"IdTipoDocumento": 1,
			"Documento": "41144069",
			"Nombres": "Loik",
			"Apellidos" : "Choua",
            "EstadoCivil": "S",
            "Telefono": "+54(381)5483777",
            "Email": "loikchoua4@gmail.com",
            "CantidadHijos": 0,
			"Usuario": "lchoua",
			"Password":"LoikCapo",
            "FechaNacimiento": "1998-05-27",
            "FechaInicio": "2020-01-01"
        }
 }
 * @apiSuccessExample {json} Success-Response:
{
    "error": null,
    "respuesta": {
        "Usuarios": {
            "IdUsuario": 6,
            "IdRol": 1,
            "IdUbicacion": 1,
            "IdTipoDocumento": 1,
            "Documento": "41144069",
            "Nombres": "Loik",
            "Apellidos": "Choua",
            "EstadoCivil": "S",
            "Telefono": "+54(381)5483777",
            "Email": "loikchoua4@gmail.com",
            "CantidadHijos": 0,
            "Usuario": "lchoua",
            "Password": "LoikCapo",
            "FechaNacimiento": "1998-05-27",
            "FechaInicio": "2020-01-01",
            "FechaAlta": "2020-05-12 18:11:39.000000",
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

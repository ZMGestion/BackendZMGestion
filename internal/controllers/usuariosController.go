package controllers

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/gestores"
	"BackendZMGestion/internal/helpers"
	"BackendZMGestion/internal/interfaces"
	"BackendZMGestion/internal/models"
	"BackendZMGestion/internal/structs"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/mitchellh/mapstructure"
)

type UsuariosController struct {
	DbHandler *db.DbHandler
}

/**
 * @api {POST} /usuarios/dame Dame Usuario
 * @apiDescription Devuelve un usuario a partir de un Id
 * @apiGroup Usuarios
 * @apiHeader {String} Authorization
 * @apiParam {Object} Usuarios
 * @apiParam {int} Usuarios.IdUsuario
 * @apiParamExample {json} Request-Example:
 {
	 "Usuarios": {
		 "IdUsuario":1
	 }
 }
 * @apiSuccessExample {json} Success-Response:
 {
    "Error": null,
    "Respuesta": {
        "Usuarios": {
            "IdUsuario": 1,
            "IdRol": 1,
            "IdUbicacion": 1,
            "IdTipoDocumento": 1,
            "Nombres": "Adam el super admin",
            "EstadoCivil": "C",
            "Telefono": "+54(381)4321719",
            "Email": "zimmermanmueblesgestion@gmail.com",
            "CantidadHijos": 2,
            "Usuario": "adam",
            "FechaNacimiento": "2020-04-11",
            "FechaInicio": "2020-04-11",
            "FechaAlta": "2020-04-11 19:50:01.000000",
            "Estado": "A"
        }
    }
}
* @apiErrorExample {json} Error-Response:
{
    "Error": {
        "Codigo": "ERROR_DEFAULT",
        "Mensaje": "Ha ocurrido un error mientras se procesaba su petición."
    },
    "Respuesta": null
}
*/
//Dame Devuelve un usuario a partir de un Id

func (uc *UsuariosController) Dame(c echo.Context) error {

	usuario := structs.Usuarios{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	mapstructure.Decode(jsonMap["Usuarios"], &usuario)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	usuariosService := models.UsuariosService{
		DbHandler: uc.DbHandler,
		Usuario:   &usuario,
	}
	result, err := usuariosService.Dame(*token)

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
 * @api {GET} /usuarios/damePorToken Dame Usuario por Token
 * @apiDescription Devuelve un usuario a partir de un Token
 * @apiGroup Usuarios
 * @apiHeader {String} Authorization
 * @apiSuccessExample {json} Success-Response:
{
    "Error": null,
    "Respuesta": {
        "Usuarios": {
            "IdUsuario": 1,
            "IdRol": 1,
            "IdUbicacion": 1,
            "IdTipoDocumento": 1,
            "Nombres": "Adam el super admin",
            "EstadoCivil": "C",
            "Telefono": "+54(381)4321719",
            "Email": "zimmermanmueblesgestion@gmail.com",
            "CantidadHijos": 2,
            "Usuario": "adam",
            "FechaNacimiento": "2020-04-11",
            "FechaInicio": "2020-04-11",
            "FechaAlta": "2020-04-11 19:50:01.000000",
            "Estado": "A"
        }
    }
}
* @apiErrorExample {json} Error-Response:
{
    "Error": {
        "Codigo": "ERROR_DEFAULT",
        "Mensaje": "Ha ocurrido un error mientras se procesaba su petición."
    },
    "Respuesta": null
}
*/
//Dame Devuelve un usuario a partir de un Token

func (uc *UsuariosController) DamePorToken(c echo.Context) error {

	usuario := structs.Usuarios{}

	token := c.Request().Header.Get("Authorization")
	usuariosService := models.UsuariosService{
		DbHandler: uc.DbHandler,
		Usuario:   &usuario,
	}
	result, err := usuariosService.DamePorToken(token)

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
    "Error": null,
    "Respuesta": {
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
    "Error": {
        "Codigo": "ERROR_DEFAULT",
        "Mensaje": "Ha ocurrido un error mientras se procesaba su petición."
    },
    "Respuesta": null
}
*/
//Dame Devuelve un usuario a partir de un Token

func (uc *UsuariosController) Crear(c echo.Context) error {

	usuario := structs.Usuarios{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Usuarios"], &usuario)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	hash := helpers.Hash(usuario.Password)

	usuario.Password = *hash

	gestorUsuarios := gestores.GestorUsuarios{
		DbHandler: uc.DbHandler,
	}
	result, err := gestorUsuarios.Crear(usuario, *token)

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
 * @api {POST} /usuarios/modificar Modificar Usuario
 * @apiDescription Permite modificar un usuario
 * @apiGroup Usuarios
 * @apiHeader {String} Authorization
 * @apiParam {Object} Usuarios
 * @apiParam {int} Usuarios.IdUsuario
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
            "IdUsuario":6,
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
    "Error": null,
    "Respuesta": {
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
    "Error": {
        "Codigo": "ERROR_DEFAULT",
        "Mensaje": "Ha ocurrido un error mientras se procesaba su petición."
    },
    "Respuesta": null
}
*/
func (uc *UsuariosController) Modificar(c echo.Context) error {

	usuario := structs.Usuarios{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Usuarios"], &usuario)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorUsuarios := gestores.GestorUsuarios{
		DbHandler: uc.DbHandler,
	}
	result, err := gestorUsuarios.Modificar(usuario, *token)

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
 * @api {POST} /usuarios/borar Borrar Usuario
 * @apiDescription Permite borrar un usuario
 * @apiGroup Usuarios
 * @apiHeader {String} Authorization
 * @apiParam {Object} Usuarios


  * @apiParamExample {json} Request-Example:
{
	 "Usuarios": {
            "IdUsuario":6,
        }
 }
 * @apiSuccessExample {json} Success-Response:
{
    "Error": null,
    "Respuesta": null
}
* @apiErrorExample {json} Error-Response:
{
    "Error": {
        "Codigo": "ERROR_DEFAULT",
        "Mensaje": "Ha ocurrido un error mientras se procesaba su petición."
    },
    "Respuesta": null
}
*/
func (uc *UsuariosController) Borrar(c echo.Context) error {

	usuario := structs.Usuarios{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Usuarios"], &usuario)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorUsuarios := gestores.GestorUsuarios{
		DbHandler: uc.DbHandler,
	}
	err = gestorUsuarios.Borrar(usuario, *token)

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
 * @api {POST} /usuarios/darAlta Dar alta Usuario
 * @apiDescription Permite dar de alta un usuario
 * @apiGroup Usuarios
 * @apiHeader {String} Authorization
 * @apiParam {Object} Usuarios


  * @apiParamExample {json} Request-Example:
{
	 "Usuarios": {
            "IdUsuario":6,
        }
 }
 * @apiSuccessExample {json} Success-Response:
{
    "Error": null,
    "Respuesta": {
        "Usuarios": {
            "IdUsuario": 8,
            "IdRol": 1,
            "IdUbicacion": 1,
            "IdTipoDocumento": 1,
            "Documento": "41144069",
            "Nombres": "Loik",
            "Apellidos": "Choua",
            "EstadoCivil": "S",
            "Telefono": "+54(381)5483777",
            "Email": "loikchoua@gmail.com",
            "CantidadHijos": 0,
            "Usuario": "lchoua",
            "FechaNacimiento": "1998-05-27",
            "FechaInicio": "2020-01-01",
            "FechaAlta": "2020-05-12 20:24:34.000000",
            "Estado": "A"
        }
    }
}
* @apiErrorExample {json} Error-Response:
{
    "Error": {
        "Codigo": "ERROR_DEFAULT",
        "Mensaje": "Ha ocurrido un error mientras se procesaba su petición."
    },
    "Respuesta": null
}
*/
func (uc *UsuariosController) DarAlta(c echo.Context) error {

	usuario := structs.Usuarios{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Usuarios"], &usuario)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	usuarioService := models.UsuariosService{
		DbHandler: uc.DbHandler,
		Usuario:   &usuario,
	}

	result, err := usuarioService.DarAlta(*token)

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
 * @api {POST} /usuarios/darBaja Dar baja Usuario
 * @apiDescription Permite dar de baja un usuario
 * @apiGroup Usuarios
 * @apiHeader {String} Authorization
 * @apiParam {Object} Usuarios


  * @apiParamExample {json} Request-Example:
{
	 "Usuarios": {
            "IdUsuario":6,
        }
 }
 * @apiSuccessExample {json} Success-Response:
{
    "Error": null,
    "Respuesta": {
        "Usuarios": {
            "IdUsuario": 8,
            "IdRol": 1,
            "IdUbicacion": 1,
            "IdTipoDocumento": 1,
            "Documento": "41144069",
            "Nombres": "Loik",
            "Apellidos": "Choua",
            "EstadoCivil": "S",
            "Telefono": "+54(381)5483777",
            "Email": "loikchoua@gmail.com",
            "CantidadHijos": 0,
            "Usuario": "lchoua",
            "FechaNacimiento": "1998-05-27",
            "FechaInicio": "2020-01-01",
            "FechaAlta": "2020-05-12 20:24:34.000000",
            "Estado": "A"
        }
    }
}
* @apiErrorExample {json} Error-Response:
{
    "Error": {
        "Codigo": "ERROR_DEFAULT",
        "Mensaje": "Ha ocurrido un error mientras se procesaba su petición."
    },
    "Respuesta": null
}
*/
func (uc *UsuariosController) DarBaja(c echo.Context) error {

	usuario := structs.Usuarios{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Usuarios"], &usuario)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	usuarioService := models.UsuariosService{
		DbHandler: uc.DbHandler,
		Usuario:   &usuario,
	}

	result, err := usuarioService.DarBaja(*token)

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
 * @api {POST} /usuarios/buscar Buscar Usuarios
 * @apiDescription Permite buscar un usuario
 * @apiGroup Usuarios
 * @apiHeader {String} Authorization
 * @apiParam {Object} Usuarios
 * @apiParam {int} [Usuarios.IdRol]
 * @apiParam {int} [Usuarios.IdUbicacion]
 * @apiParam {string} [Usuarios.Documento]
 * @apiParam {string} [Usuarios.Nombres]
 * @apiParam {string} [Usuarios.Apellidos]
 * @apiParam {string} [Usuarios.EstadoCivil]
 * @apiParam {string} [Usuarios.Telefono]
 * @apiParam {string} [Usuarios.Email]
 * @apiParam {string} [Usuarios.Usuario]

  * @apiParamExample {json} Request-Example:
{
	"Usuarios":{
		"Usuario":"nbachs"
	}
}
 * @apiSuccessExample {json} Success-Response:
{
    "Error": null,
    "Respuesta": [
        {
            "Roles": {
                "IdRol": 1,
                "Rol": "Administradores"
            },
            "Ubicaciones": {
                "IdUbicacion": 1,
                "Ubicacion": "Casa Central Tucumán"
            },
            "Usuarios": {
                "IdUsuario": 5,
                "IdRol": 1,
                "IdUbicacion": 1,
                "IdTipoDocumento": 1,
                "Documento": "39477073",
                "Nombres": "Nicolas",
                "Apellidos": "Bachs",
                "EstadoCivil": "S",
                "Telefono": "+543814491954",
                "Email": "nicolas.bachs@gmail.com",
                "CantidadHijos": 0,
                "Usuario": "nbachs",
                "FechaUltIntento": "2020-05-08 00:52:23.000000",
                "FechaNacimiento": "1995-12-27",
                "FechaInicio": "2019-11-22",
                "FechaAlta": "2020-05-08 00:49:18.000000",
                "Estado": "A"
            }
        }
    ]
}
* @apiErrorExample {json} Error-Response:
{
    "Error": {
        "Codigo": "ERROR_DEFAULT",
        "Mensaje": "Ha ocurrido un error mientras se procesaba su petición."
    },
    "Respuesta": null
}
*/
//Dame Devuelve un usuario a partir de un Token

func (uc *UsuariosController) Buscar(c echo.Context) error {

	usuario := structs.Usuarios{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Usuarios"], &usuario)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorUsuarios := gestores.GestorUsuarios{
		DbHandler: uc.DbHandler,
	}
	result, err := gestorUsuarios.Buscar(usuario, *token)

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
 * @api {POST} /usuarios/restablecerPassword Restablecer contraseña
 * @apiDescription Permite restablecer la contraseña de un usuario
 * @apiGroup Usuarios
 * @apiHeader {String} Authorization
 * @apiParam {Object} Usuarios
 * @apiParam {int} Usuarios.IdUsuario
 * @apiParam {string} Usuarios.Password

  * @apiParamExample {json} Request-Example:
{
	"Usuarios":{
        "IdUsuario": 1,
        "Password":"Nueva pass"
	}
}
 * @apiSuccessExample {json} Success-Response:
{
    "Error": null
    "Respuesta": null
}
* @apiErrorExample {json} Error-Response:
{
    "Error": {
        "Codigo": "ERROR_USUARIO_ESTA_BAJA",
        "Mensaje": "El usuario no existe o ya está dado de baja."
    },
    "Respuesta": null
}
*/
//RestablecerPassword Reestablece la password de un usuario
func (uc *UsuariosController) RestablecerPassword(c echo.Context) error {

	usuario := structs.Usuarios{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Usuarios"], &usuario)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	hash := helpers.Hash(usuario.Password)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	usuario.Password = *hash
	fmt.Println(usuario.Password)

	usuarioService := models.UsuariosService{
		DbHandler: uc.DbHandler,
		Usuario:   &usuario,
	}
	err = usuarioService.RestablecerPassword(*token)

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
 * @api {POST} /usuarios/modificarPassword Modificar contraseña
 * @apiDescription Permite modificar la contraseña de un usuario
 * @apiGroup Usuarios
 * @apiHeader {String} Authorization
 * @apiParam {Object} UsuariosActual
 * @apiParam {string} UsuariosActual.Password
 * @apiParam {Object} UsuariosNuevo
 * @apiParam {string} UsuariosNuevo.Password
  * @apiParamExample {json} Request-Example:
{
	"UsuariosActual":{
		"Password":"Nueva Pass"
	},
	"UsuariosNuevo":{
		"Password":"Hola123"
	}
}
 * @apiSuccessExample {json} Success-Response:
{
    "Error": null,
    "Respuesta": {
        "Usuarios": {
            "IdUsuario": 9,
            "IdRol": 1,
            "IdUbicacion": 1,
            "IdTipoDocumento": 1,
            "Documento": "42664256",
            "Nombres": "Chloe",
            "Apellidos": "Choua",
            "EstadoCivil": "S",
            "Telefono": "+54(381)4451337",
            "Email": "chloechoua@gmail.com",
            "CantidadHijos": 0,
            "Usuario": "cchoua",
            "FechaNacimiento": "2000-05-17",
            "FechaInicio": "2020-01-01",
            "FechaAlta": "2020-05-13 20:25:19.000000",
            "Estado": "A"
        }
    }
}
* @apiErrorExample {json} Error-Response:
{
    "Error": {
        "Codigo": "ERROR_PASSWORD_INCORRECTA",
        "Mensaje": "La contraseña ingresada no es correcta."
    },
    "Respuesta": null
}
*/
//ModificarPassword Modifica la password de un usuario
func (uc *UsuariosController) ModificarPassword(c echo.Context) error {

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	var usuarioActual, usuarioNuevo structs.Usuarios

	err = mapstructure.Decode(jsonMap["UsuariosActual"], &usuarioActual)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	err = mapstructure.Decode(jsonMap["UsuariosNuevo"], &usuarioNuevo)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	hashActual := helpers.Hash(usuarioActual.Password)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	hashNuevo := helpers.Hash(usuarioNuevo.Password)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	gestorUsuarios := gestores.GestorUsuarios{
		DbHandler: uc.DbHandler,
	}

	result, err := gestorUsuarios.ModificarPassword(*hashActual, *hashNuevo, *token)

	if err != nil || result == nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error: nil,
	}
	result.Password = ""

	response.AddModels(result)

	return c.JSON(http.StatusOK, response)
}

/**
 * @api {POST} /usuarios/iniciarSesion Iniciar Sesion
 * @apiDescription Permite a un usuario iniciar sesion
 * @apiGroup Usuarios
 * @apiParam {Object} Usuarios
 * @apiParam {string} [Usuarios.Email]
 * @apiParam {string} [Usuarios.Usuario]
 * @apiParam {string} Usuarios.Password

  * @apiParamExample {json} Request-Example:
{
	 "Usuarios": {
			"Usuario": "lchoua",
			"Password":"LoikCapo",
        }
 }
 * @apiSuccessExample {json} Success-Response:
{
    "Error": null,
    "Respuesta": {
        "Usuarios": {
            "IdUsuario": 5,
            "IdRol": 1,
            "IdUbicacion": 1,
            "IdTipoDocumento": 1,
            "Documento": "39477073",
            "Nombres": "Nicolas",
            "Apellidos": "Bachs",
            "EstadoCivil": "S",
            "Telefono": "+543814491954",
            "Email": "nicolas.bachs@gmail.com",
            "CantidadHijos": 0,
            "Usuario": "nbachs",
            "Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODk0MTI5MzIsImlhdCI6MTU4OTQwOTMzMiwidXNlciI6Im5iYWNocyJ9.IX4fpwOpjp8-TjPMVKqT1NmZ0gQG0shyrLfLq6ETi-M",
            "FechaNacimiento": "1995-12-27",
            "FechaInicio": "2019-11-22",
            "FechaAlta": "2020-05-08 00:49:18.000000",
            "Estado": "A"
        }
    }
}
* @apiErrorExample {json} Error-Response:
{
    "Error": {
        "Codigo": "ERROR_LOGIN_INCORRECTO",
        "Mensaje": "El nombre de usuario o contrasena ingresados no son correctos."
    },
    "Respuesta": null
}
*/
//Iniciar Sesion Permite a un usuario iniciar sesion
func (uc *UsuariosController) IniciarSesion(c echo.Context) error {

	usuario := structs.Usuarios{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Usuarios"], &usuario)

	hash := helpers.Hash(usuario.Password)

	usuario.Password = *hash

	token, err := helpers.CreateToken(usuario.Usuario + usuario.Email)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	if token == nil {
		return interfaces.GenerarRespuestaError(errors.New("ERROR_DEFAULT"), http.StatusUnprocessableEntity)
	}

	usuario.Token = *token

	usuariosService := models.UsuariosService{
		DbHandler: uc.DbHandler,
		Usuario:   &usuario,
	}

	result, err := usuariosService.IniciarSesion()

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
 * @api {POST} /usuarios/cerrarSesion Cerrar Sesion
 * @apiDescription Permite a un usuario cerrar la sesion de otro usuario
 * @apiGroup Usuarios
 * @apiHeader {String} Authorization
 * @apiParam {Object} Usuarios
 * @apiParam {int} Usuarios.IdUsuario

  * @apiParamExample {json} Request-Example:
{
	"Usuarios":{
        "IdUsuario": 1
	}
}
 * @apiSuccessExample {json} Success-Response:
{
    "Error": null
    "Respuesta": null
}
* @apiErrorExample {json} Error-Response:
{
    "Error": {
        "Codigo": "ERROR_NOEXISTE_USUARIO",
        "Mensaje": "El usuario indicado no existe."
    },
    "Respuesta": null
}
*/
//RestablecerPassword Reestablece la password de un usuario
func (uc *UsuariosController) CerrarSesion(c echo.Context) error {

	usuario := structs.Usuarios{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Usuarios"], &usuario)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	usuarioService := models.UsuariosService{
		DbHandler: uc.DbHandler,
		Usuario:   &usuario,
	}
	err = usuarioService.CerrarSesion(*token)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error:     nil,
		Respuesta: nil,
	}

	return c.JSON(http.StatusOK, response)
}

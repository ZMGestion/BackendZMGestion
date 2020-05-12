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

//RolesController contiene los metodos para los endpoints de la API referidos a Roles.
type RolesController struct {
	DbHandler *db.DbHandler
}

/**
 * @api {POST} /roles/dame Dame Rol
 * @apiPermission Administradores
 * @apiDescription Devuelve un rol a partir de un Id
 * @apiGroup Roles
 * @apiParam {Object} Roles
 * @apiParam {int} Roles.IdRol
 * @apiParamExample {json} Request-Example:
 {
	 "Roles": {
		 "IdRol":2
	 }
 }
 * @apiSuccessExample {json} Success-Response:
 {
    "Error": null,
    "Respuesta": {
		"Roles": {
            "IdRol": 2,
            "Rol": "Vendedores",
            "FechaAlta": "2020-04-09 15:01:35.000000",
            "Descripcion": "Este rol es para los vendedores"
        }
	}
}
* @apiErrorExample {json} Error-Response:
 {
    "Error": {
        "Codigo": "ERROR_NOEXISTE_ROL",
        "Mensaje": "No existe el rol."
    },
    "Respuesta": null
}
*/
//Dame Devuelve un rol a partir de un Id
func (rc *RolesController) Dame(c echo.Context) error {

	rol := structs.Roles{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	mapstructure.Decode(jsonMap["Roles"], &rol)

	//_ = c.Request().Header.Get("Authorization")
	rolesService := models.RolesService{
		DbHandler: rc.DbHandler,
		Rol:       &rol,
	}
	result, err := rolesService.Dame()

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
 * @api {GET} /roles/crear Crear Rol
 * @apiPermission Administradores
 * @apiDescription Permite crear un rol
 * @apiGroup Roles
 * @apiHeader {String} Authorization
 * @apiParam {Object} Roles
 * @apiParam {string} Roles.Rol
 * @apiParam {string} [Roles.Descripcion]
 * @apiParamExample {json} Request-Example:
 {
	 "Roles": {
		 "Rol": "Encargados"
	 }
 }
 * @apiSuccessExample {json} Success-Response:
 {
    "Error": null,
    "Respuesta": {
		"Roles" : {
			"IdRol": 7,
			"Rol": "Encargados",
			"FechaAlta": "2020-04-09 15:01:35.000000",
			"Descripcion": ""
		}
	}
}
* @apiErrorExample {json} Error-Response:
 {
    "Error": {
        "Codigo": "ERROR_EXISTE_NOMBREROL",
        "Mensaje": "El nombre de rol ya existe."
    },
    "Respuesta": null
}
*/
func (rc *RolesController) Crear(c echo.Context) error {
	rol := structs.Roles{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	mapstructure.Decode(jsonMap["Roles"], &rol)

	token := c.Request().Header.Get("Authorization")

	gestorRoles := gestores.GestorRoles{
		DbHandler: rc.DbHandler,
	}

	result, err := gestorRoles.Crear(rol, token)

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
 * @api {POST} /roles/listar Listar Roles
 * @apiPermission Administradores
 * @apiDescription Devuelve una lista de roles
 * @apiGroup Roles
 * @apiSuccessExample {json} Success-Response:
 {
    "Error": null,
    "Respuesta": [
		{
			"Roles":{
				"IdRol": 1,
				"Rol": "Administradores",
				"FechaAlta": "2020-04-09 15:01:35.000000",
				"Observaciones": ""
			}
		},
		{
			"Roles":{
				"IdRol": 2,
				"Rol": "Vendedores",
				"FechaAlta": "2020-04-09 15:01:35.000000",
				"Observaciones": ""
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
//Listar Lista los roles
func (rc *RolesController) Listar(c echo.Context) error {

	gestorRoles := gestores.GestorRoles{
		DbHandler: rc.DbHandler,
	}

	result, err := gestorRoles.Listar()

	if err != nil || result == nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error: nil,
	}

	var respuesta []map[string]interface{}
	for _, el := range result {
		objeto := make(map[string]interface{})
		objeto["Roles"] = el
		respuesta = append(respuesta, objeto)
	}
	response.Respuesta = respuesta

	return c.JSON(http.StatusOK, response)
}

/**
 * @api {POST} /roles/borrar Borrar Rol
 * @apiPermission Administradores
 * @apiDescription Borra un rol a partir de su Id
 * @apiGroup Roles
 * @apiHeader {String} Authorization
 * @apiParam {Object} Roles
 * @apiParam {int} Roles.IdRol
 * @apiParamExample {json} Request-Example:
 {
	 "Roles": {
		 "IdRol": "2"
	 }
 }
 * @apiSuccessExample {json} Success-Response:
 {
    "Error": null,
    "Respuesta": null
}
*/
//Borrar Devuelve un rol a partir de un Id
func (rc *RolesController) Borrar(c echo.Context) error {

	rol := structs.Roles{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	mapstructure.Decode(jsonMap["Roles"], &rol)

	//_ = c.Request().Header.Get("Authorization")
	gestorRoles := gestores.GestorRoles{
		DbHandler: rc.DbHandler,
	}

	token := c.Request().Header.Get("Authorization")

	err = gestorRoles.Borrar(rol, token)

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
 * @api {GET} /roles/modificar Modificar Rol
 * @apiPermission Administradores
 * @apiDescription Permite modificar un rol
 * @apiGroup Roles
 * @apiHeader {String} Authorization
 * @apiParam {Object} Roles
 * @apiParam {int} Roles.IdRol
 * @apiParam {string} Roles.Rol
 * @apiParam {string} [Roles.Descripcion]
 * @apiParamExample {json} Request-Example:
 {
	 "Roles": {
		 "IdRol": 7,
		 "Rol": "Los encargados",
		 "Descripcion": "Nueva descripcion"
	 }
 }
 * @apiSuccessExample {json} Success-Response:
 {
    "Error": null,
    "Respuesta": {
		"IdRol": 7,
		"Rol": "Los encargados",
		"FechaAlta": "2020-04-09 15:01:35.000000",
		"Descripcion": "Nueva descripcion"
	}
}
* @apiErrorExample {json} Error-Response:
 {
    "Error": {
        "Codigo": "ERROR_EXISTE_NOMBREROL",
        "Mensaje": "El nombre de rol ya existe."
    },
    "Respuesta": null
}
*/
func (rc *RolesController) Modificar(c echo.Context) error {
	rol := structs.Roles{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	mapstructure.Decode(jsonMap["Roles"], &rol)

	token := c.Request().Header.Get("Authorization")

	gestorRoles := gestores.GestorRoles{
		DbHandler: rc.DbHandler,
	}

	result, err := gestorRoles.Modificar(rol, token)

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
 * @api {POST} /roles/listarPermisos Listar Permisos
 * @apiPermission Administradores
 * @apiName Listar Permisos
 * @apiDescription Devuelve una lista de permisos para un determinado rol
 * @apiGroup Roles
 * @apiSuccessExample {json} Success-Response:
 {
    "Error": null,
    "Respuesta": [
		{
            "Permisos": {
                "IdPermiso": 1,
                "Permiso": "Crear rol"
            }
        },
        {
            "Permisos": {
                "IdPermiso": 2,
                "Permiso": "Borrar rol"
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
//Listar Lista los permisos para un rol
func (rc *RolesController) ListarPermisos(c echo.Context) error {

	rol := structs.Roles{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	mapstructure.Decode(jsonMap["Roles"], &rol)

	rolesService := models.RolesService{
		DbHandler: rc.DbHandler,
		Rol:       &rol,
	}

	result, err := rolesService.ListarPermisos()

	if err != nil || result == nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error: nil,
	}

	var respuesta []map[string]interface{}
	for _, el := range result {
		objeto := make(map[string]interface{})
		objeto["Permisos"] = el
		respuesta = append(respuesta, objeto)
	}
	response.Respuesta = respuesta

	return c.JSON(http.StatusOK, response)
}

/**
 * @api {POST} /roles/asignarPermisos Asignar Permisos
 * @apiPermission Administradores
 * @apiDescription Permite asignar los permisos a un determinado rol
 * @apiGroup Roles
 * @apiParam {Object} Roles
 * @apiParam {int} Roles.IdRol
 * @apiParam {Object[]} Permisos
 * @apiParamExample {json} Request-Example:
 {
	"Roles": {
		"IdRol": 9
	},
	"Permisos": [
		{
			"IdPermiso": 3
		},
		{
			"IdPermiso": 4
		}
	]
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
//Listar Lista los permisos para un rol
func (rc *RolesController) AsignarPermisos(c echo.Context) error {

	rol := structs.Roles{}
	var permisos []structs.Permisos

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	mapstructure.Decode(jsonMap["Roles"], &rol)
	mapstructure.Decode(jsonMap["Permisos"], &permisos)

	rolesService := models.RolesService{
		DbHandler: rc.DbHandler,
		Rol:       &rol,
	}

	token := c.Request().Header.Get("Authorization")

	err = rolesService.AsignarPermisos(permisos, token)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error:     nil,
		Respuesta: nil,
	}

	return c.JSON(http.StatusOK, response)
}

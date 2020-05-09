package controllers

import (
	"BackendZMGestion/internal/gestores"
	"BackendZMGestion/internal/helpers"
	"BackendZMGestion/internal/interfaces"
	"BackendZMGestion/internal/models"
	"BackendZMGestion/internal/structs"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/mitchellh/mapstructure"
)

//RolesController contiene los metodos para los endpoints de la API referidos a Roles.
type RolesController struct {
	Service *models.RolesService
	Gestor  *gestores.GestorRoles
}

/**
 * @api {GET} /roles/crear
 * @apiPermission Administradores
 * @apiDescription Permite crear un rol
 * @apiGroup Roles
 * @apiSuccessExample {json} Success-Response:
 {
    "Error": null,
    "Respuesta": {
		"IdRol": 7,
		"Rol": "Encargados",
		"FechaAlta": "2020-04-09 15:01:35.000000",
		"Descripcion": ""
	}

}
*/
func (rc *RolesController) Crear(c echo.Context) error {
	rol := structs.Roles{}

	jsonMap, err := helpers.GenerateMapFromContext(c)
	if err != nil {
		objError := models.Errores{
			Codigo:  "ERROR_DEFAULT",
			Mensaje: helpers.GetError(err),
		}
		response := interfaces.Response{
			Error:     &objError,
			Respuesta: nil,
		}
		return echo.NewHTTPError(http.StatusUnprocessableEntity, response)
	}
	mapstructure.Decode(jsonMap["Roles"], &rol)

	token := c.Request().Header.Get("Authorization")

	result, err := rc.Gestor.Crear(rol, token)

	if err != nil || result == nil {
		objError := models.Errores{
			Codigo:  "ERROR_DEFAULT",
			Mensaje: helpers.GetError(err),
		}
		response := interfaces.Response{
			Error:     &objError,
			Respuesta: nil,
		}
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}

	response := interfaces.Response{
		Error: nil,
	}

	response.AddModels(result)

	return c.JSON(http.StatusOK, response)

}

/**
 * @api {GET} /roles/listar
 * @apiPermission Administradores
 * @apiDescription Listar todos los roles
 * @apiGroup Roles
 * @apiSuccessExample {json} Success-Response:
 {
    "Error": null,
    "Respuesta": [
        {
            "IdRol": 1,
            "Rol": "Administradores",
            "FechaAlta": "2020-04-09 15:01:35.000000",
            "Descripcion": ""
        },
        {
            "IdRol": 2,
            "Rol": "Vendedores",
            "FechaAlta": "2020-04-09 15:01:35.000000",
            "Descripcion": ""
        },
        {
            "IdRol": 3,
            "Rol": "Fabricantes",
            "FechaAlta": "2020-04-09 15:01:35.000000",
            "Descripcion": ""
        }
    ]
}
*/

/**
 * @api {POST} /roles/listar
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
*/
//Listar Lista los roles
func (rc *RolesController) Listar(c echo.Context) error {

	result, err := rc.Gestor.Listar()

	if err != nil || result == nil {

		objError := models.Errores{
			Codigo:  "ERROR_DEFAULT",
			Mensaje: helpers.GetError(err),
		}
		response := interfaces.Response{
			Error:     &objError,
			Respuesta: nil,
		}
		return echo.NewHTTPError(http.StatusBadRequest, response)
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
 * @api {POST} /roles/dame
 * @apiPermission Administradores
 * @apiDescription Devuelve un rol a partir de un Id
 * @apiGroup Roles
 * @apiParam {Object} Roles
 * @apiParam {int} Roles.IdRol
 * @apiParamExample {json} Request-Example:
 {
	 "IdRol":2
 }
 * @apiSuccessExample {json} Success-Response:
 {
    "Error": null,
    "Respuesta": [
        {
            "IdRol": 2,
            "Rol": "Vendedores",
            "FechaAlta": "2020-04-09 15:01:35.000000",
            "Observaciones": ""
        },
    ]
}
*/
//Dame Devuelve un rol a partir de un Id
func (rc *RolesController) Dame(c echo.Context) error {

	rol := structs.Roles{}

	jsonMap, err := helpers.GenerateMapFromContext(c)
	mapstructure.Decode(jsonMap["Roles"], &rol)

	//_ = c.Request().Header.Get("Authorization")
	rc.Service.Rol = &rol
	result, err := rc.Service.Dame()

	if err != nil || result == nil {
		if err != nil {
			fmt.Println(err.Error())
		}
		objError := models.Errores{
			Codigo:  "ERROR_DEFAULT",
			Mensaje: helpers.GetError(err),
		}
		response := interfaces.Response{
			Error:     &objError,
			Respuesta: nil,
		}
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}

	response := interfaces.Response{
		Error: nil,
	}
	response.AddModels(result)

	return c.JSON(http.StatusOK, response)
}

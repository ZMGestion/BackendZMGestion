package controllers

import (
	"BackendZMGestion/src/gestores"
	"BackendZMGestion/src/interfaces"
	"BackendZMGestion/src/models"
	"BackendZMGestion/src/structs"
	"net/http"

	"github.com/labstack/echo"
)

type RolesController struct {
	Service *models.RolesService
	Gestor  *gestores.GestorRoles
}

/**
 * @api {GET} /roles/listar
 * @apiPermission Administradores
 * @apiDescription Listar todos los roles
 * @apiGroup Roles
 * @apiSuccessExample {json} Success-Response:
 {
    "Estado": "",
    "Mensaje": "Ok",
    "Respuesta": [
        {
            "IdRol": 1,
            "Rol": "Administradores",
            "FechaAlta": "2020-04-09 15:01:35.000000",
            "Observaciones": ""
        },
        {
            "IdRol": 2,
            "Rol": "Vendedores",
            "FechaAlta": "2020-04-09 15:01:35.000000",
            "Observaciones": ""
        },
        {
            "IdRol": 3,
            "Rol": "Fabricantes",
            "FechaAlta": "2020-04-09 15:01:35.000000",
            "Observaciones": ""
        }
    ]
}
*/

//Listar Lista los roles
func (rc *RolesController) Listar(c echo.Context) error {
	roles, err := rc.Gestor.Listar()

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	response := interfaces.Response{
		Estado:    "OK",
		Mensaje:   "Ok",
		Respuesta: roles,
	}

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
    "Estado": "",
    "Mensaje": "Ok",
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

	if err := c.Bind(&rol); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	rc.Service.Rol = &rol
	err := rc.Service.Dame()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	response := interfaces.Response{
		Estado:    "OK",
		Mensaje:   "Ok",
		Respuesta: rc.Service.Rol,
	}

	return c.JSON(http.StatusOK, response)
}

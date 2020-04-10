package controllers

import (
	"BackendZMGestion/src/gestores"
	"BackendZMGestion/src/interfaces"
	"BackendZMGestion/src/models"
	"net/http"

	"github.com/labstack/echo"
)

type RolesController struct {
	Modelo *models.Roles
	Gestor *gestores.GestorRoles
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
    "Objetos": [
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
		Estado:  "OK",
		Mensaje: "Ok",
		Objetos: roles,
	}

	return c.JSON(http.StatusOK, response)
}

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

//Listar commented
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

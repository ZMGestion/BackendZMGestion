package main

import (
	"BackendZMGestion/src/controllers"
	"BackendZMGestion/src/db"
	"BackendZMGestion/src/gestores"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func InitRouter(h *db.DbHandler) *echo.Echo {
	r := echo.New()
	initRoutes(r, h)

	r.Use(middleware.Logger())
	//r.Use(middleware.Gzip())

	return r
}

func initRoutes(r *echo.Echo, h *db.DbHandler) {
	gestorRoles := &gestores.GestorRoles{DbHandler: h}
	controllerRoles := &controllers.RolesController{Gestor: gestorRoles}

	r.GET("/roles/listar", controllerRoles.Listar)

}

package main

import (
	"BackendZMGestion/internal/controllers"
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/gestores"
	"BackendZMGestion/internal/models"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func InitRouter(h *db.DbHandler) *echo.Echo {
	r := echo.New()
	initRoutes(r, h)

	r.Use(middleware.Logger())
	//r.Use(middleware.Gzip())

	r.Static("/apidoc", "apidoc")

	return r
}

func initRoutes(r *echo.Echo, h *db.DbHandler) {
	gestorRoles := &gestores.GestorRoles{DbHandler: h}
	serviceRoles := &models.RolesService{DbHandler: h}
	controllerRoles := &controllers.RolesController{
		Gestor:  gestorRoles,
		Service: serviceRoles,
	}

	r.GET("/roles/listar", controllerRoles.Listar)
	r.POST("/roles/dame", controllerRoles.Dame)

}

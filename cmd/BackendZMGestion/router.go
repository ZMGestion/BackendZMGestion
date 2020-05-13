package main

import (
	"BackendZMGestion/internal/controllers"
	"BackendZMGestion/internal/db"

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
	controllerRoles := &controllers.RolesController{
		DbHandler: h,
	}

	r.GET("/roles/listar", controllerRoles.Listar)
	r.POST("/roles/dame", controllerRoles.Dame)
	r.POST("/roles/crear", controllerRoles.Crear)
	r.POST("/roles/borrar", controllerRoles.Borrar)
	r.POST("/roles/modificar", controllerRoles.Modificar)
	r.POST("/roles/asignarPermisos", controllerRoles.AsignarPermisos)
	r.POST("/roles/listarPermisos", controllerRoles.ListarPermisos)

	controllerUsuarios := &controllers.UsuariosController{
		DbHandler: h,
	}

	r.POST("/usuarios/dame", controllerUsuarios.Dame)
	r.GET("/usuarios/damePorToken", controllerUsuarios.DamePorToken)
	r.POST("/usuarios/crear", controllerUsuarios.Crear)
	r.POST("/usuarios/modificar", controllerUsuarios.Modificar)
	r.POST("/usuarios/borrar", controllerUsuarios.Borrar)
	r.POST("/usuarios/darAlta", controllerUsuarios.DarAlta)
	r.POST("/usuarios/darBaja", controllerUsuarios.DarBaja)
	r.POST("/usuarios/buscar", controllerUsuarios.Buscar)
	r.POST("/usuarios/restablecerPassword", controllerUsuarios.RestablecerPassword)
	r.POST("/usuarios/modificarPassword", controllerUsuarios.ModificarPassword)
}

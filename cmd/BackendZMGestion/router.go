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

	r.Use(middleware.CORS())
	r.Use(middleware.Logger())
	//r.Use(middleware.Gzip())

	r.Static("/apidoc", "apidoc")

	return r
}

func initRoutes(r *echo.Echo, h *db.DbHandler) {
	controllerRoles := &controllers.RolesController{
		DbHandler: h,
	}
	r.GET("/roles", controllerRoles.Listar)
	r.POST("/roles/dame", controllerRoles.Dame)
	r.POST("/roles/crear", controllerRoles.Crear)
	r.POST("/roles/borrar", controllerRoles.Borrar)
	r.POST("/roles/modificar", controllerRoles.Modificar)
	r.POST("/roles/asignarPermisos", controllerRoles.AsignarPermisos)
	r.POST("/roles/permisos", controllerRoles.ListarPermisos)

	controllerPermisos := &controllers.PermisosController{
		DbHanlder: h,
	}
	r.POST("/permisos", controllerPermisos.Listar)
	controllerUsuarios := &controllers.UsuariosController{
		DbHandler: h,
	}
	r.POST("/usuarios", controllerUsuarios.Buscar)
	r.POST("/usuarios/dame", controllerUsuarios.Dame)
	r.GET("/usuarios/damePorToken", controllerUsuarios.DamePorToken)
	r.POST("/usuarios/crear", controllerUsuarios.Crear)
	r.POST("/usuarios/modificar", controllerUsuarios.Modificar)
	r.POST("/usuarios/borrar", controllerUsuarios.Borrar)
	r.POST("/usuarios/darAlta", controllerUsuarios.DarAlta)
	r.POST("/usuarios/darBaja", controllerUsuarios.DarBaja)
	r.POST("/usuarios/restablecerPassword", controllerUsuarios.RestablecerPassword)
	r.POST("/usuarios/modificarPassword", controllerUsuarios.ModificarPassword)
	r.POST("/usuarios/iniciarSesion", controllerUsuarios.IniciarSesion)
	r.POST("/usuarios/cerrarSesion", controllerUsuarios.CerrarSesion)
	r.GET("/usuarios/tiposDocumento", controllerUsuarios.ListarTiposDocumento)

	controllerUbicaciones := &controllers.UbicacionesController{
		DbHandler: h,
	}
	r.GET("/ubicaciones", controllerUbicaciones.Listar)
	r.POST("/ubicaciones/dame", controllerUbicaciones.Dame)
	r.POST("/ubicaciones/crear", controllerUbicaciones.Crear)
	r.POST("/ubicaciones/borrar", controllerUbicaciones.Borrar)
	r.POST("/ubicaciones/modificar", controllerUbicaciones.Modificar)
	r.POST("/ubicaciones/darAlta", controllerUbicaciones.DarAlta)
	r.POST("/ubicaciones/darBaja", controllerUbicaciones.DarBaja)

	controllerPaises := &controllers.PaisesController{
		DbHandler: h,
	}
	r.GET("/paises", controllerPaises.Listar)

	controllerProvincias := &controllers.ProvinciasController{
		DbHanlder: h,
	}
	r.POST("/provincias", controllerProvincias.Listar)

	controllerCiudades := &controllers.CiudadesController{
		DbHanlder: h,
	}
	r.POST("/ciudades", controllerCiudades.Listar)

	controllerClientes := &controllers.ClientesController{
		DbHanlder: h,
	}
	r.POST("/clientes", controllerClientes.Buscar)
	r.POST("/clientes/crear", controllerClientes.Crear)
	r.POST("/clientes/dame", controllerClientes.Dame)
	r.POST("/clientes/modificar", controllerClientes.Modificar)
	r.POST("/clientes/darAlta", controllerClientes.DarAlta)
	r.POST("/clientes/darBaja", controllerClientes.DarBaja)
	r.POST("/clientes/borrar", controllerClientes.Borrar)
	r.POST("/clientes/domicilios", controllerClientes.ListarDomicilios)
	r.POST("/clientes/domicilios/agregar", controllerClientes.AgregarDomicilio)
	r.POST("/clientes/domicilios/quitar", controllerClientes.QuitarDomicilio)

	controllerTelas := &controllers.TelasController{
		DbHandler: h,
	}
	r.POST("/telas", controllerTelas.Buscar)
	r.POST("/telas/crear", controllerTelas.Crear)
	r.POST("/telas/modificar", controllerTelas.Modificar)
	r.POST("/telas/darAlta", controllerTelas.DarAlta)
	r.POST("/telas/darBaja", controllerTelas.DarBaja)
	r.POST("/telas/borrar", controllerTelas.Borrar)
	r.POST("/telas/dame", controllerTelas.Dame)
	r.POST("/telas/precios", controllerTelas.ListarPrecios)
	r.POST("/telas/precios/modificar", controllerTelas.ModificarPrecio)

	controllerGruposProducto := &controllers.GruposProductoController{
		DbHandler: h,
	}
	r.POST("/gruposProducto", controllerGruposProducto.Buscar)
	r.POST("/gruposProducto/dame", controllerGruposProducto.Dame)
	r.POST("/gruposProducto/crear", controllerGruposProducto.Crear)
	r.POST("/gruposProducto/modificar", controllerGruposProducto.Modificar)
	r.POST("/gruposProducto/borrar", controllerGruposProducto.Borrar)
	r.POST("/gruposProducto/darAlta", controllerGruposProducto.DarAlta)
	r.POST("/gruposProducto/darBaja", controllerGruposProducto.DarBaja)

	controllerProductos := &controllers.ProductosController{
		DbHandler: h,
	}
	r.POST("/productos", controllerProductos.Buscar)
	r.POST("/productos/dame", controllerProductos.Dame)
	r.POST("/productos/crear", controllerProductos.Crear)
	r.POST("/productos/modificar", controllerProductos.Modificar)
	r.POST("/productos/borrar", controllerProductos.Borrar)
	r.POST("/productos/darBaja", controllerProductos.DarBaja)
	r.POST("/productos/darAlta", controllerProductos.DarAlta)
	r.POST("/productos/precios", controllerProductos.ListarPrecios)
	r.GET("/productos/tiposProducto", controllerProductos.ListarTiposProducto)
	r.GET("/productos/categoriasProducto", controllerProductos.ListarCategoriasProducto)
	r.POST("/productos/precios/modificar", controllerProductos.ModificarPrecio)

	controllerProductosFinales := &controllers.ProductosFinalesController{
		DbHandler: h,
	}
	r.POST("/productosFinales", controllerProductosFinales.Buscar)
	r.POST("/productosFinales/dame", controllerProductosFinales.Dame)
	r.POST("/productosFinales/crear", controllerProductosFinales.Crear)
	r.POST("/productosFinales/modificar", controllerProductosFinales.Modificar)
	r.POST("/productosFinales/borrar", controllerProductosFinales.Borrar)
	r.POST("/productosFinales/darBaja", controllerProductosFinales.DarBaja)
	r.POST("/productosFinales/darAlta", controllerProductosFinales.DarAlta)
	r.GET("/productosFinales/lustres", controllerProductosFinales.ListarLustres)

	controllerPresupuestos := &controllers.PresupuestosController{
		DbHanlder: h,
	}
	r.POST("/presupuestos", controllerPresupuestos.Buscar)
	r.POST("/presupuestos/crear", controllerPresupuestos.Crear)
	r.POST("/presupuestos/dame", controllerPresupuestos.Dame)
	r.POST("/presupuestos/dameMultiple", controllerPresupuestos.DamePresupuestos)
	r.POST("/presupuestos/modificar", controllerPresupuestos.Modificar)
	r.POST("/presupuestos/borrar", controllerPresupuestos.Borrar)
	r.POST("/presupuestos/pasarACreado", controllerPresupuestos.PasarACreado)
	r.POST("/presupuestos/transformarEnVenta", controllerPresupuestos.TransformarEnVenta)
	//Lineas presupuesto
	r.POST("/presupuestos/lineasPresupuesto/crear", controllerPresupuestos.CrearLineaPresupuesto)
	r.POST("/presupuestos/lineasPresupuesto/dame", controllerPresupuestos.DameLineaPresupuesto)
	r.POST("/presupuestos/lineasPresupuesto/modificar", controllerPresupuestos.ModificarLineaPresupuesto)
	r.POST("/presupuestos/lineasPresupuesto/borrar", controllerPresupuestos.BorrarLineaPresupuestos)

	controllerVentas := &controllers.VentasController{
		DbHandler: h,
	}
	r.POST("/ventas", controllerVentas.Buscar)
	r.POST("/ventas/crear", controllerVentas.Crear)
	r.POST("/ventas/dame", controllerVentas.Dame)
	r.POST("/ventas/modificar", controllerVentas.Modificar)
	r.POST("/ventas/borrar", controllerVentas.Borrar)
	r.POST("/ventas/chequearPrecios", controllerVentas.ChequearPrecios)
	r.POST("/ventas/revisar", controllerVentas.Revisar)
	r.POST("/ventas/cancelar", controllerVentas.Cancelar)
	//LineasVenta
	r.POST("/ventas/lineasVenta/crear", controllerVentas.CrearLineaVenta)
	r.POST("/ventas/lineasVenta/dame", controllerVentas.DameLineaVenta)
	r.POST("/ventas/lineasVenta/modificar", controllerVentas.ModificarLineaVenta)
	r.POST("/ventas/lineasVenta/borrar", controllerVentas.BorrarLineaVenta)
	r.POST("/ventas/lineasVenta/cancelar", controllerVentas.CancelarLineaVenta)

}

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

type ProductosController struct {
	DbHandler *db.DbHandler
}

/**
 * @api {POST} /productos/crear Crear Producto
 * @apiDescription Permite crear un prodcuto
 * @apiGroup Productos
 * @apiHeader {String} Authorization
 * @apiParam {Object} Productos
 * @apiParam {int} Productos.IdCategoriaProducto
 * @apiParam {int} Productos.IdGrupoProducto
 * @apiParam {string} Productos.IdTipoProducto
 * @apiParam {string} Productos.Producto
 * @apiParam {double} Productos.LongitudTela
 * @apiParam {string} Productos.Observaciones
 * @apiParam {Object} Precios
 * @apiParam {double} Precios.Precio

  * @apiParamExample {json} Request-Example:
{
    "Productos":{
		"IdCategoriaProducto":1,
		"IdGrupoProducto":2,
		"IdTipoProducto":"P",
		"Producto": "Silla 2",
		"LongitudTela":15.0,
		"Observaciones":""
    },
    "Precios":{
        "Precio":1.20
    }
}
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":{
		"Productos":{
			"IdProducto": 6,
			"IdCategoriaProducto": 1,
			"IdGrupoProducto": 2,
			"IdTipoProducto": "P",
			"Producto": "Silla 2",
			"LongitudTela": 15,
			"FechaAlta": "2020-07-06 20:58:48.000000",
			"FechaBaja": "",
			"Observaciones": "",
			"Estado": "A"
		}
	}
}
* @apiErrorExample {json} Error-Response:
{
    "error": {
        "codigo": "ERROR_DEFAULT",
        "mensaje": "Ha ocurrido un error mientras se procesaba su petición."
    },
    "respuesta": null
}
*/
//Crear
func (pc *ProductosController) Crear(c echo.Context) error {

	producto := structs.Productos{}
	precios := structs.Precios{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Productos"], &producto)
	mapstructure.Decode(jsonMap["Precios"], &precios)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorProductos := gestores.GestorProductos{
		DbHandler: pc.DbHandler,
	}
	result, err := gestorProductos.Crear(producto, precios, *token)

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
 * @api {POST} /productos/modificar Modificar Producto
 * @apiDescription Permite modificar un prodcuto
 * @apiGroup Productos
 * @apiHeader {String} Authorization
 * @apiParam {Object} Productos
 * @apiParam {int} Productos.IdProducto
 * @apiParam {int} Productos.IdCategoriaProducto
 * @apiParam {int} Productos.IdGrupoProducto
 * @apiParam {string} Productos.IdTipoProducto
 * @apiParam {string} Productos.Producto
 * @apiParam {double} Productos.LongitudTela
 * @apiParam {string} Productos.Observaciones
 * @apiParamExample {json} Request-Example:
{
    "Productos":{
		"IdProducto":6,
		"IdCategoriaProducto":1,
		"IdGrupoProducto":2,
		"IdTipoProducto":"P",
		"Producto": "Silla 2",
		"LongitudTela":15.0,
		"Observaciones":""
	},
	"Precios": {
		"Precio": 1200
	}
}
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta": null
}
* @apiErrorExample {json} Error-Response:
{
    "error": {
        "codigo": "ERROR_DEFAULT",
        "mensaje": "Ha ocurrido un error mientras se procesaba su petición."
    },
    "respuesta": null
}
*/
//Modificar
func (pc *ProductosController) Modificar(c echo.Context) error {

	producto := structs.Productos{}
	precios := structs.Precios{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Productos"], &producto)
	mapstructure.Decode(jsonMap["Precios"], &precios)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorProductos := gestores.GestorProductos{
		DbHandler: pc.DbHandler,
	}
	result, err := gestorProductos.Modificar(producto, precios, *token)

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
 * @api {POST} /productos/borrar Borrar Producto
 * @apiDescription Permite borrar un prodcuto
 * @apiGroup Productos
 * @apiHeader {String} Authorization
 * @apiParam {Object} Productos
 * @apiParam {int} Productos.IdProducto

{
    "Productos":{
		"IdProducto":6
    }
}
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta": null
}
* @apiErrorExample {json} Error-Response:
{
    "error": {
        "codigo": "ERROR_DEFAULT",
        "mensaje": "Ha ocurrido un error mientras se procesaba su petición."
    },
    "respuesta": null
}
*/
//Modificar
func (pc *ProductosController) Borrar(c echo.Context) error {

	producto := structs.Productos{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Productos"], &producto)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorProductos := gestores.GestorProductos{
		DbHandler: pc.DbHandler,
	}
	err = gestorProductos.Borrar(producto, *token)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error: nil,
	}

	return c.JSON(http.StatusOK, response)
}

/**
 * @api {POST} /productos/darAlta Dar alta Producto
 * @apiDescription Permite dar de alta un producto
 * @apiGroup Productos
 * @apiHeader {String} Authorization
 * @apiParam {Object} Productos
 * @apiParam {int} Productos.IdProducto


  * @apiParamExample {json} Request-Example:
{
	 "Productos": {
            "IdProducto":6
        }
 }
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":{
		"Productos":{
			"IdProducto": 5,
			"IdCategoriaProducto": 1,
			"IdGrupoProducto": 5,
			"IdTipoProducto": "P",
			"Producto": "Silla de prueba 4",
			"LongitudTela": 12,
			"FechaAlta": "2020-07-05 22:51:36.000000",
			"FechaBaja": "2020-07-06 21:22:39.000000",
			"Observaciones": "",
			"Estado": "A"
		}
	}
}
* @apiErrorExample {json} Error-Response:
{
    "error": {
        "codigo": "ERROR_DEFAULT",
        "mensaje": "Ha ocurrido un error mientras se procesaba su petición."
    },
    "respuesta": null
}
*/
func (pc *ProductosController) DarAlta(c echo.Context) error {

	producto := structs.Productos{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Productos"], &producto)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	productosService := models.ProductosService{
		Producto:  &producto,
		DbHandler: pc.DbHandler,
	}

	result, err := productosService.DarAlta(*token)

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
 * @api {POST} /productos/darBaja Dar baja Producto
 * @apiDescription Permite dar de baja un producto
 * @apiGroup Productos
 * @apiHeader {String} Authorization
 * @apiParam {Object} Productos
 * @apiParam {int} Productos.IdProducto


  * @apiParamExample {json} Request-Example:
{
	 "Productos": {
            "IdProducto": 16
        }
 }
 * @apiSuccessExample {json} Success-Response:
{
    "error": null,
    "respuesta": {
        "Productos": {
            "IdProducto": 16,
            "IdCategoriaProducto": 1,
            "IdGrupoProducto": 5,
            "IdTipoProducto": "P",
            "Producto": "Braun",
            "LongitudTela": 0,
            "FechaAlta": "2020-08-07 05:04:46.000000",
            "FechaBaja": "2020-08-16 21:25:54.000000",
            "Observaciones": "",
            "Estado": "B"
        }
    }
}
* @apiErrorExample {json} Error-Response:
{
    "error": {
        "codigo": "ERROR_DEFAULT",
        "mensaje": "Ha ocurrido un error mientras se procesaba su petición."
    },
    "respuesta": null
}
*/
func (pc *ProductosController) DarBaja(c echo.Context) error {

	producto := structs.Productos{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Productos"], &producto)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	productosService := models.ProductosService{
		Producto:  &producto,
		DbHandler: pc.DbHandler,
	}

	result, err := productosService.DarBaja(*token)

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
 * @api {POST} /productos Buscar Producto
 * @apiDescription Permite buscar un producto
 * @apiGroup Productos
 * @apiHeader {String} Authorization
 * @apiParam {Object} Productos
 * @apiParam {int} [Productos.IdCategoriaProducto]
 * @apiParam {int} [Productos.IdGrupoProducto]
 * @apiParam {string} [Productos.IdTipoProducto]
 * @apiParam {string} [Productos.Producto]
 * @apiParam {string} [Productos.Estado]
 * @apiParam {Object} Paginaciones
 * @apiParam {int} [Paginaciones.Pagina]
 * @apiParam {int} [Paginaciones.LongitudPagina]
 * @apiParamExample {json} Request-Example:
{
	 "Productos": {},
	 "Paginaciones":{
		 "Pagina":1,
		 "LongitudPagina":1
	 }
}
 * @apiSuccessExample {json} Success-Response:
{
    "error": null,
    "respuesta": {
        "resultado": [
            {
                "CategoriasProducto": {
                    "Categoria": "Sillas",
                    "IdCategoriaProducto": 1
                },
                "GruposProducto": {
                    "Estado": "A",
                    "Grupo": "Grupo1",
                    "IdGrupoProducto": 5
                },
                "Precios": {
                    "FechaAlta": "2020-08-08 01:22:27.000000",
                    "IdPrecio": 102,
                    "Precio": 2780
                },
                "Productos": {
                    "Estado": "A",
                    "FechaAlta": "2020-08-07 04:43:12.000000",
                    "FechaBaja": null,
                    "IdCategoriaProducto": 1,
                    "IdGrupoProducto": 5,
                    "IdProducto": 15,
                    "IdTipoProducto": "P",
                    "LongitudTela": 0,
                    "Observaciones": null,
                    "Producto": "Becker"
                },
                "TiposProducto": {
                    "IdCategoriaProducto": "P",
                    "TipoProducto": "Productos fabricables"
                }
            },
            {
                "CategoriasProducto": {
                    "Categoria": "Sillas",
                    "IdCategoriaProducto": 1
                },
                "GruposProducto": {
                    "Estado": "A",
                    "Grupo": "Grupo1",
                    "IdGrupoProducto": 5
                },
                "Precios": {
                    "FechaAlta": "2020-08-07 05:04:46.000000",
                    "IdPrecio": 96,
                    "Precio": 4500
                },
                "Productos": {
                    "Estado": "A",
                    "FechaAlta": "2020-08-07 05:04:46.000000",
                    "FechaBaja": "2020-08-16 18:12:12.000000",
                    "IdCategoriaProducto": 1,
                    "IdGrupoProducto": 5,
                    "IdProducto": 16,
                    "IdTipoProducto": "P",
                    "LongitudTela": 0,
                    "Observaciones": null,
                    "Producto": "Braun"
                },
                "TiposProducto": {
                    "IdCategoriaProducto": "P",
                    "TipoProducto": "Productos fabricables"
                }
            }
        ],
        "Paginaciones": {
            "Pagina": 1,
            "LongitudPagina": 2,
            "CantidadTotal": 12
        }
    }
}
* @apiErrorExample {json} Error-Response:
{
    "error": {
        "codigo": "ERROR_DEFAULT",
        "mensaje": "Ha ocurrido un error mientras se procesaba su petición."
    },
    "respuesta": null
}
*/
func (pc *ProductosController) Buscar(c echo.Context) error {

	producto := structs.Productos{}
	paginacion := structs.Paginaciones{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Productos"], &producto)
	mapstructure.Decode(jsonMap["Paginaciones"], &paginacion)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	gestorProductos := gestores.GestorProductos{
		DbHandler: pc.DbHandler,
	}

	result, err := gestorProductos.Buscar(producto, paginacion, *token)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error:     nil,
		Respuesta: result,
	}

	return c.JSON(http.StatusOK, response)
}

/**
 * @api {POST} /productos/precios/modificar Modificar Precio Producto
 * @apiDescription Permite modificar el precio de un prodcuto
 * @apiGroup Productos
 * @apiHeader {String} Authorization
 * @apiParam {Object} Productos
 * @apiParam {int} Productos.IdProducto
 * @apiParam {Object} Precios
 * @apiParam {double} Precios.Precio

  * @apiParamExample {json} Request-Example:
{
    "Productos":{
		"IdProducto":5
    },
    "Precios":{
        "Precio":13.20
    }
}
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":{
		"Precios":{
			"FechaAlta": "2020-07-06 22:47:21.000000",
			"IdPrecio": 30,
			"Precio": 13
		},
		"Productos":{
			"Estado": "A",
			"FechaAlta": "2020-07-05 22:51:36.000000",
			"FechaBaja": "2020-07-06 21:22:39.000000",
			"IdCategoriaProducto": 1,
			"IdGrupoProducto": 5,
			"IdProducto": 5,
			"IdTipoProducto": "P",
			"LongitudTela": 12,
			"Observaciones": null,
			"Producto": "Silla de prueba 4"
		}
	}
}
* @apiErrorExample {json} Error-Response:
{
    "error": {
        "codigo": "ERROR_DEFAULT",
        "mensaje": "Ha ocurrido un error mientras se procesaba su petición."
    },
    "respuesta": null
}
*/
//ModificarPrecio
func (pc *ProductosController) ModificarPrecio(c echo.Context) error {

	producto := structs.Productos{}
	precio := structs.Precios{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Productos"], &producto)
	mapstructure.Decode(jsonMap["Precios"], &precio)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	productoService := models.ProductosService{
		DbHandler: pc.DbHandler,
		Producto:  &producto,
	}
	result, err := productoService.ModificarPrecio(precio, *token)

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
 * @api {POST} /productos/precios Listar Precios Producto
 * @apiDescription Permite listar el historico de precios de un producto
 * @apiGroup Productos
 * @apiHeader {String} Authorization
 * @apiParam {Object} Productos
 * @apiParam {int} Productos.IdProducto
 * @apiParamExample {json} Request-Example:
{
    "Productos":{
		"IdProducto":5
    }
}
 * @apiSuccessExample {json} Success-Response:
{
    "error": null,
    "respuesta": [
        {
            "Precios": {
                "FechaAlta": "2020-08-07 04:43:12.000000",
                "IdPrecio": 94,
                "Precio": 2350
            }
        },
        {
            "Precios": {
                "FechaAlta": "2020-08-07 04:48:46.000000",
                "IdPrecio": 95,
                "Precio": 2490
            }
        },
        {
            "Precios": {
                "FechaAlta": "2020-08-08 01:22:27.000000",
                "IdPrecio": 102,
                "Precio": 2780
            }
        }
    ]
}
* @apiErrorExample {json} Error-Response:
{
    "error": {
        "codigo": "ERROR_DEFAULT",
        "mensaje": "Ha ocurrido un error mientras se procesaba su petición."
    },
    "respuesta": null
}
*/
//ModificarPrecio
func (pc *ProductosController) ListarPrecios(c echo.Context) error {

	producto := structs.Productos{}
	precio := structs.Precios{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Productos"], &producto)
	mapstructure.Decode(jsonMap["Precios"], &precio)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	productoService := models.ProductosService{
		DbHandler: pc.DbHandler,
		Producto:  &producto,
	}
	result, err := productoService.ListarPrecios(*token)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error:     nil,
		Respuesta: result,
	}

	return c.JSON(http.StatusOK, response)
}

/**
 * @api {POST} /productos/dame Dame Producto
 * @apiDescription Permite instanciar un producto a partir de su Id
 * @apiGroup Productos
 * @apiHeader {String} Authorization
 * @apiParam {Object} Productos
 * @apiParam {int} Productos.IdProducto


  * @apiParamExample {json} Request-Example:
{
	 "Productos": {
            "IdProducto":6
        }
 }
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":{
		"Precios":{
			"IdPrecio": 30,
			"Precio": 13,
			"Tipo": "",
			"IdReferencia": 0,
			"FechaAlta": "2020-07-06 22:47:21.000000"
		},
		"Productos":{
			"IdProducto": 5,
			"IdCategoriaProducto": 1,
			"IdGrupoProducto": 2,
			"IdTipoProducto": "P",
			"Producto": "Silla de prueba 4",
			"LongitudTela": 12,
			"FechaAlta": "2020-07-05 22:51:36.000000",
			"FechaBaja": "2020-07-06 21:22:39.000000",
			"Observaciones": "",
			"Estado": "A"
		},
		"TiposProducto": {
            "TipoProducto": "Productos fabricables",
            "IdCategoriaProducto": "P"
        },
        "GruposProducto": {
            "Grupo": "Grupo de prueba 2",
            "Estado": "A",
            "IdGrupoProducto": 2
        },
        "CategoriasProducto": {
            "Categoria": "Sillas",
            "IdCategoriaProducto": 1
        }
	}
}
* @apiErrorExample {json} Error-Response:
{
    "error": {
        "codigo": "ERROR_DEFAULT",
        "mensaje": "Ha ocurrido un error mientras se procesaba su petición."
    },
    "respuesta": null
}
*/
func (pc *ProductosController) Dame(c echo.Context) error {

	producto := structs.Productos{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Productos"], &producto)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	productosService := models.ProductosService{
		Producto:  &producto,
		DbHandler: pc.DbHandler,
	}

	result, err := productosService.Dame(*token)

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
 * @api {GET} /productos/tiposProducto Listar Tipos de Producto
 * @apiDescription Permite listar los tipos de productos
 * @apiGroup Productos
 * @apiHeader {String} Authorization
 * @apiSuccessExample {json} Success-Response:
{
    "error": null,
    "respuesta": [
        {
            "TiposProducto": {
                "Descripcion": null,
                "IdTipoProducto": "P",
                "TipoProducto": "Productos fabricables"
            }
		},
		{
            "TiposProducto": {
                "Descripcion": null,
                "IdTipoProducto": "N",
                "TipoProducto": "Productos no fabricables"
            }
        }
    ]
}
* @apiErrorExample {json} Error-Response:
{
    "error": {
        "codigo": "ERROR_DEFAULT",
        "mensaje": "Ha ocurrido un error mientras se procesaba su petición."
    },
    "respuesta": null
}
*/
func (pc *ProductosController) ListarTiposProducto(c echo.Context) error {

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	productosService := models.ProductosService{
		DbHandler: pc.DbHandler,
	}

	result, err := productosService.ListarTiposProducto(*token)

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
 * @api {GET} /productos/categoriasProducto Listar Categorias de Producto
 * @apiDescription Permite listar las categorias de producto
 * @apiGroup Productos
 * @apiHeader {String} Authorization
 * @apiSuccessExample {json} Success-Response:
{
    "error": null,
    "respuesta": [
        {
            "CategoriasProducto": {
                "Categoria": "Sillas",
                "Descripcion": "Categoria de Sillas",
                "IdCategoriaProducto": 1
            }
		},
		{
            "CategoriasProducto": {
                "Categoria": "Sillones",
                "Descripcion": "Categoria de Sillones",
                "IdCategoriaProducto": 2
            }
        }
    ]
}
* @apiErrorExample {json} Error-Response:
{
    "error": {
        "codigo": "ERROR_DEFAULT",
        "mensaje": "Ha ocurrido un error mientras se procesaba su petición."
    },
    "respuesta": null
}
*/
func (pc *ProductosController) ListarCategoriasProducto(c echo.Context) error {

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	productosService := models.ProductosService{
		DbHandler: pc.DbHandler,
	}

	result, err := productosService.ListarCategorias(*token)

	if err != nil || result == nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error:     nil,
		Respuesta: result,
	}

	return c.JSON(http.StatusOK, response)
}

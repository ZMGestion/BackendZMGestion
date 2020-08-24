package controllers

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/helpers"
	"BackendZMGestion/internal/interfaces"
	"BackendZMGestion/internal/models"
	"BackendZMGestion/internal/structs"
	"net/http"

	"github.com/labstack/echo"
	"github.com/mitchellh/mapstructure"
)

type LineasPresupuestoController struct {
	DbHanlder *db.DbHandler
}

/**
 * @api {POST} /lineasPresupuestos/crear Crear linea de presupuesto
 * @apiDescription Permite crear una linea de presupuesto
 * @apiGroup Presupuestos
 * @apiHeader {String} Authorization
 * @apiParam {Object} LineasProducto
 * @apiParam {int} LineasProducto.IdPresupuesto
 * @apiParam {float} LineasProducto.PrecioUnitario
 * @apiParam {int} LineasProducto.Cantidad
 * @apiParam {Object} ProductosFinales
 * @apiParam {int} ProductosFinales.IdProducto
 * @apiParam {int} ProductosFinales.IdTela
 * @apiParam {int} ProductosFinales.IdLuste

  * @apiParamExample {json} Request-Example:
{
    "Presupuestos": {
		"IdCliente": 12,
		"IdUbicacion": 13
	}
}
 * @apiSuccessExample {json} Success-Response:
{
	"error": null,
	"respuesta":{
		"Presupuestos":{
			"Estado": "E",
			"FechaAlta": "2020-08-22 20:02:10.000000",
			"IdCliente": 3,
			"IdPresupuesto": 2,
			"IdUbicacion": 1,
			"IdUsuario": 1,
			"IdVenta": null,
			"Observaciones": null,
			"PeriodoValidez": 15
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
func (lpc *LineasPresupuestoController) Crear(c echo.Context) error {

	lineaPresupuesto := structs.LineasProducto{}
	productoFinal := structs.ProductosFinales{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["LineasProducto"], &lineaPresupuesto)
	mapstructure.Decode(jsonMap["ProductosFinales"], &productoFinal)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	presupuestosService := models.PresupuestosService{
		DbHanlder: lpc.DbHanlder,
	}
	result, err := presupuestosService.CrearLineaPresupuesto(lineaPresupuesto, productoFinal, *token)

	if err != nil || result == nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error:     nil,
		Respuesta: result,
	}

	return c.JSON(http.StatusOK, response)
}

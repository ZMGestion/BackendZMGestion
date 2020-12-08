package controllers

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/helpers"
	"BackendZMGestion/internal/interfaces"
	"BackendZMGestion/internal/models"
	"net/http"

	"github.com/labstack/echo"
)

//ReportesController Estructura de controller
type ReportesController struct {
	DbHandler *db.DbHandler
}

/**
 * @api {GET} /reportes/stock Reporte Stock
 * @apiDescription Devuelve el listado de los productos finales con su stock
 * @apiGroup Paises
 * @apiSuccessExample {json} Success-Response:
 {
    "error": null,
    "respuesta": [
		{
    "error": null,
    "respuesta": [
        {
            "Telas": {
                "Tela": "Bella Lino",
                "IdTela": 17
            },
            "Lustres": {
                "Lustre": "CA1",
                "IdLustre": 1
            },
            "Productos": {
                "Estado": "A",
                "Producto": "Silla Cadiz Baja",
                "FechaAlta": "2020-12-07 00:07:20.000000",
                "FechaBaja": null,
                "IdProducto": 50,
                "LongitudTela": 1.25,
                "Observaciones": null,
                "IdTipoProducto": "P",
                "IdGrupoProducto": 9,
                "IdCategoriaProducto": 18
            },
            "ProductosFinales": {
                "Estado": "A",
                "IdTela": 17,
                "IdLustre": 1,
                "FechaAlta": "2020-12-07 15:59:35.000000",
                "FechaBaja": null,
                "_Cantidad": 0,
                "IdProducto": 50,
                "IdProductoFinal": 1
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
func (rc *ReportesController) Stock(c echo.Context) error {

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	reportesService := models.ReportesService{
		DbHandler: rc.DbHandler,
	}

	result, err := reportesService.Stock(*token)

	if err != nil || result == nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error:     nil,
		Respuesta: result,
	}

	return c.JSON(http.StatusOK, response)
}

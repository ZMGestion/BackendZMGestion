package controllers

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/helpers"
	"BackendZMGestion/internal/interfaces"
	"BackendZMGestion/internal/models"
	"net/http"

	"github.com/labstack/echo"
)

type ProductosFinalesController struct {
	DbHandler *db.DbHandler
}

/**
 * @api {GET} /productosFinales/lustres Listar Lustres
 * @apiDescription Devuelve una lista de lustres
 * @apiGroup ProductosFinales
 * @apiSuccessExample {json} Success-Response:
 {
    "error": null,
    "respuesta": [
		{
			"Roles":{
				"IdRol": 1,
				"Rol": "Administradores",
				"FechaAlta": "2020-04-09 15:01:35.000000",
				"Observaciones": ""
			}
		},
		{
			"Roles":{
				"IdRol": 2,
				"Rol": "Vendedores",
				"FechaAlta": "2020-04-09 15:01:35.000000",
				"Observaciones": ""
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
func (lc *ProductosFinalesController) ListarLustres(c echo.Context) error {

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	productosFinalesService := models.ProductosFinalesService{
		DbHandler: lc.DbHandler,
	}

	result, err := productosFinalesService.ListarLustres(*token)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error:     nil,
		Respuesta: result,
	}

	return c.JSON(http.StatusOK, response)
}

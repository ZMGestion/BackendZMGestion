package controllers

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/gestores"
	"BackendZMGestion/internal/interfaces"
	"net/http"

	"github.com/labstack/echo"
)

type PaisesController struct {
	DbHandler *db.DbHandler
}

/**
 * @api {GET} /paises Listar Paises
 * @apiDescription Devuelve una lista de paises
 * @apiGroup Paises
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
//Listar Lista los roles
func (pc *PaisesController) Listar(c echo.Context) error {

	gestorPaises := gestores.GestorPaises{
		DbHandler: pc.DbHandler,
	}

	result, err := gestorPaises.Listar()

	if err != nil || result == nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error:     nil,
		Respuesta: result,
	}

	return c.JSON(http.StatusOK, response)
}

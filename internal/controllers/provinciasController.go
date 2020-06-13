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

type ProvinciasController struct {
	DbHanlder *db.DbHandler
}

/**
 * @api {POST} /provincias Listar Provincias
 * @apiDescription Devuelve la lista de provincias de un país
 * @apiGroup Provincias
 * @apiParam {Object} Paises
 * @apiParam {int} Paises.IdPais
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
func (pc *ProvinciasController) Listar(c echo.Context) error {
	pais := structs.Paises{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	mapstructure.Decode(jsonMap["Paises"], &pais)

	paisesService := models.PaisesService{
		DbHandler: pc.DbHanlder,
		Paises:    &pais,
	}

	result, err := paisesService.ListarProvincias()

	if err != nil || result == nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error:     nil,
		Respuesta: result,
	}

	return c.JSON(http.StatusOK, response)
}

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
 * @apiParamExample {json} Request-Example:
{
    "Paises":{
        "IdPais":"AR"
    }
}
 * @apiSuccessExample {json} Success-Response:
 {
    "error": null,
    "respuesta": [
        {
            "IdPais": "",
            "IdProvincia": 2,
            "Provincia": "Salta"
        },
        {
            "IdPais": "",
            "IdProvincia": 1,
            "Provincia": "Tucumán"
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

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error:     nil,
		Respuesta: result,
	}

	return c.JSON(http.StatusOK, response)
}

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

type CiudadesController struct {
	DbHanlder *db.DbHandler
}

/**
 * @api {POST} /ciudades Listar Ciudades
 * @apiDescription Devuelve la lista de ciudades de una provincia
 * @apiGroup Ciudades
 * @apiParam {Object} Provincias
 * @apiParam {int} Provincias.IdProvincia
 * @apiParam {int} Provincias.IdCiudad
 * @apiParamExample {json} Request-Example:
{
    "Provincias":{
        "IdPais":"AR",
        "IdProvincia":1
    }
}
 * @apiSuccessExample {json} Success-Response:
{
    "error": null,
    "respuesta": [
        {
            "IdCiudad": 1,
            "IdProvincia": 1,
            "IdPais": "AR",
            "Ciudad": "San Miguel de Tucumán"
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
func (cc *CiudadesController) Listar(c echo.Context) error {
	provincia := structs.Provincias{}

	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	mapstructure.Decode(jsonMap["Provincias"], &provincia)

	provinciasService := models.ProvinciasService{
		DbHandler:  cc.DbHanlder,
		Provincias: &provincia,
	}

	result, err := provinciasService.ListarCiudades()

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error: nil,
	}
	var respuesta []map[string]interface{}
	for _, el := range result {
		objeto := make(map[string]interface{})
		objeto["Ciudades"] = el
		respuesta = append(respuesta, objeto)
	}
	response.Respuesta = respuesta

	return c.JSON(http.StatusOK, response)
}

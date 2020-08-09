package controllers

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/gestores"
	"BackendZMGestion/internal/helpers"
	"BackendZMGestion/internal/interfaces"
	"BackendZMGestion/internal/structs"
	"net/http"

	"github.com/labstack/echo"
	"github.com/mitchellh/mapstructure"
)

type PermisosController struct {
	DbHanlder *db.DbHandler
}

/**
 * @api {POST} /permisos Listar Permisos
 * @apiDescription Devuelve la lista de permisos
 * @apiGroup Permisos
 * @apiSuccessExample {json} Success-Response:
 {
    "error": null,
    "respuesta": [
		{
			"Permisos":{
				"IdPermiso": 1,
				"Permiso": "Crear rol",
				"Descripcion": "Permite a un usuario crear un nuevo rol."
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
func (pc *PermisosController) Listar(c echo.Context) error {

	paginacion := structs.Paginaciones{}
	jsonMap, err := helpers.GenerateMapFromContext(c)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	mapstructure.Decode(jsonMap["Paginaciones"], &paginacion)

	headerToken := c.Request().Header.Get("Authorization")
	token, err := helpers.GetToken(headerToken)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}
	gestorPermisos := gestores.GestorPermisos{
		DbHandler: pc.DbHanlder,
	}

	result, err := gestorPermisos.Listar(paginacion, *token)

	if err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error:     nil,
		Respuesta: result,
	}

	return c.JSON(http.StatusOK, response)
}

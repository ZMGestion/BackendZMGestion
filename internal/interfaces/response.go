package interfaces

import (
	"BackendZMGestion/internal/helpers"
	"BackendZMGestion/internal/models"

	"github.com/labstack/echo"
)

type Response struct {
	Error     *models.Errores
	Respuesta interface{}
}

func (r *Response) AddModels(elements ...interface{}) *Response {
	res := helpers.GenerateJSONFromModels(elements...)
	r.Respuesta = res
	return r
}

func GenerarRespuestaError(err error, code int) *echo.HTTPError {
	codigo, mensaje := helpers.GetError(err)
	objError := models.Errores{
		Codigo:  codigo,
		Mensaje: mensaje,
	}
	response := Response{
		Error:     &objError,
		Respuesta: nil,
	}
	return echo.NewHTTPError(code, response)
}

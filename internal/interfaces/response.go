package interfaces

import (
	"BackendZMGestion/internal/helpers"
	"BackendZMGestion/internal/models"
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

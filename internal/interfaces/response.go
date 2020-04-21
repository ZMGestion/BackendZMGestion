package interfaces

import (
	"BackendZMGestion/internal/enums"
	"BackendZMGestion/internal/helpers"
)

type Response struct {
	Estado    enums.Status
	Mensaje   string
	Respuesta interface{}
}

func (r *Response) AddModels(elements ...interface{}) *Response {
	res := helpers.GenerateJSONFromModels(elements...)
	r.Respuesta = res
	return r
}

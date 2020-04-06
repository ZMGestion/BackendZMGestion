package interfaces

import (
	"BackendZMGestion/models"
)

type IGestor interface {
	crear(o IObjeto) models.Response
	buscarAvanzado(o IObjeto) []IObjeto
	modificar(o IObjeto) models.Response
	borrar(o IObjeto) models.Response
}

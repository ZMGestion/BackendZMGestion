package models

import (
	"BackendZMGestion/interfaces"
)

type Response struct {
	Estado  interfaces.Status
	Mensaje string
	Objetos []IObjeto
}

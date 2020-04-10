package interfaces

import "BackendZMGestion/src/enums"

type Response struct {
	Estado  enums.Status
	Mensaje string
	Objetos interface{}
}

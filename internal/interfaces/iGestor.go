package interfaces

type IGestor interface {
	Crear(o IObjeto) (*IObjeto, error)
	BuscarAvanzado(o IObjeto) ([]*IObjeto, error)
	Modificar(o IObjeto) (*IObjeto, error)
	Borrar(o IObjeto) (*IObjeto, error)
}

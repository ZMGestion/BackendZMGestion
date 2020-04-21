package interfaces

type IObjeto interface {
	Dame() (*IObjeto, error)
	DarAlta() (string, error)
	DarBaja() (string, error)
}

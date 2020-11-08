package structs

type ParametrosBusqueda struct {
	FechaInicio       string `json:"FechaInicio"`
	FechaFin          string `json:"FechaFin"`
	FechaEntregaDesde string `json:FechaEntregaDesde`
	FechaEntregaHasta string `json:FechaEntregaHasta`
}

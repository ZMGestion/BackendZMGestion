package structs

type OrdenesProduccion struct {
	IdOrdenProduccion int    `json:"IdOrdenProduccion"`
	IdUsuario         int    `json:"IdUsuario"`
	IdVenta           int    `json:"IdVenta"`
	FechaAlta         string `json:"FechaAlta"`
	Observaciones     string `json:"Observaciones"`
	Estado            string `json:"Estado"`
}

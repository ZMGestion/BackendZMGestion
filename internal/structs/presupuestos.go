package structs

type Presupuestos struct {
	IdPresupuesto  int    `json:"IdPresupuesto"`
	IdCliente      int    `json:"IdCliente"`
	IdVenta        int    `json:"IdVenta"`
	IdUbicacion    int    `json:"IdUbicacion"`
	IdUsuario      int    `json:"IdUsuario"`
	PeriodoValidez int    `json:"PeriodoValidez"`
	FechaAlta      string `json:"FechaAlta"`
	Observaciones  string `json:"Observaciones"`
	Estado         string `json:"Estado"`
}

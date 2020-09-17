package structs

type Comprobantes struct {
	IdComprobante     int     `json:"IdComprobante"`
	IdVenta           int     `json:"IdVenta"`
	IdUsuario         int     `json:"IdUsuario"`
	Tipo              string  `json:"Tipo"`
	NumeroComprobante int     `json:"NumeroComprobante"`
	Monto             float32 `json:"Monto"`
	FechaAlta         string  `json:"FechaAlta"`
	FechaBaja         string  `json:"FechaBaja"`
	Observaciones     string  `json:"Observaciones"`
	Estado            string  `json:"Estado"`
}

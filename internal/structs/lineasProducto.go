package structs

type LineasProducto struct {
	IdLineaProducto      int     `json:"IdLineaProducto"`
	IdLineaProductoPadre int     `json:"IdLineaProductoPadre"`
	IdUbicacion          int     `json:"IdUbicacion"`
	IdReferencia         int     `json:"IdReferencia"`
	Tipo                 string  `json:"Tipo"`
	PrecioUnitario       float32 `json:"PrecioUnitario"`
	Cantidad             int     `json:"Cantidad"`
	FechaAlta            string  `json:"FechaAlta"`
	FechaCancelacion     string  `json:"FechaCancelacion"`
	Estado               string  `json:"Estado"`
}

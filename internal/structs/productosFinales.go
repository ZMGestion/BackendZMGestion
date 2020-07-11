package structs

type ProductosFinales struct {
	IdProductoFinal int    `json:"IdProductoFinal"`
	IdProducto      int    `json:"IdProducto"`
	IdLustre        int    `json:"IdLustre"`
	IdTela          int    `json:"IdTela"`
	FechaAlta       string `json:"FechaAlta"`
	FechaBaja       string `json:"FechaBaja"`
	Estado          string `json:"Estado"`
}

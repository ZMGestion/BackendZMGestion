package structs

type Productos struct {
	IdProducto          int     `json:"IdProducto"`
	IdCategoriaProducto int     `json:"IdCategoriaProducto"`
	IdGrupoProducto     int     `json:"IdGrupoProducto"`
	IdTipoProducto      string  `json:"IdTipoProducto"`
	Producto            string  `json:"Producto"`
	LongitudTela        float32 `json:"LongitudTela"`
	FechaAlta           string  `json:"FechaAlta"`
	FechaBaja           string  `json:"FechaBaja"`
	Observaciones       string  `json:"Observaciones"`
	Estado              string  `json:"Estado"`
}

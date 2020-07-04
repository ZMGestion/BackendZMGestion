package structs

type GruposProducto struct {
	IdGrupoProducto int    `json:"IdGrupoProducto"`
	Grupo           string `json:"Grupo"`
	FechaAlta       string `json:"FechaAlta"`
	FechaBaja       string `json:"FechaBaja"`
	Descripcion     string `json:"Descripcion"`
	Estado          string `json:"Estado"`
}

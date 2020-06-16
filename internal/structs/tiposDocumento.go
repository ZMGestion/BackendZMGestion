package structs

type TiposDocumento struct {
	IdTipoDocumento int    `json:"IdTipoDocumento"`
	TipoDocumento   string `json:"TipoDocumento"`
	Descripcion     string `json:"Descripcion"`
}

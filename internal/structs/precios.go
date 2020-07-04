package structs

type Precios struct {
	IdPrecio     int     `json:"IdPrecio"`
	Precio       float32 `json:"Precio"`
	Tipo         string  `json:"Tipo"`
	IdReferencia int     `json:"IdReferencia"`
	FechaAlta    string  `json:"FechaAlta"`
}

package structs

type Remitos struct {
	IdRemito      int    `json:"IdRemito,omitempty"`
	IdUbicacion   int    `json:"IdUbicacion,omitempty"`
	IdUsuario     int    `json:"IdUsuario,omitempty"`
	Tipo          string `json:"Tipo,omitempty"`
	FechaEntrega  string `json:"FechaEntrega,omitempty"`
	FechaAlta     string `json:"FechaAlta,omitempty"`
	Observaciones string `json:"Observaciones,omitempty"`
	Estado        string `json:"Estado,omitempty"`
}

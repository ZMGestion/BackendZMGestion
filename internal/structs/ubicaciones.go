package structs

type Ubicaciones struct {
	IdUbicacion   int    `json:"IdUbicacion,omitempty"`
	IdDomicilio   int    `json:"IdDomicilio,omitempty"`
	Ubicacion     string `json:"Ubicacion,omitempty"`
	FechaAlta     string `json:"FechaAlta,omitempty"`
	FechaBaja     string `json:"FechaBaja,omitempty"`
	Observaciones string `json:"Observaciones,omitempty"`
	Estado        string `json:"Estado,omitempty"`
}

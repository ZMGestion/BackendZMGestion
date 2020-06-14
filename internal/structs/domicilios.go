package structs

type Domicilios struct {
	IdDomicilio   int    `json:"IdDomicilio"`
	IdCiudad      int    `json:"IdCiudad"`
	IdProvincia   int    `json:"IdProvincia"`
	IdPais        string `json:"IdPais"`
	Domicilio     string `json:"Domicilio"`
	CodigoPostal  string `json:"CodigoPostal"`
	FechaAlta     string `json:"FechaAlta"`
	Observaciones string `json:"Observaciones"`
}

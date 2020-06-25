package structs

type Ciudades struct {
	IdCiudad    int    `json:"IdCiudad"`
	IdProvincia int    `json:"IdProvincia"`
	IdPais      string `json:"IdPais"`
	Ciudad      string `json:"Ciudad"`
}

package structs

type Ciudades struct {
	IdCiudad    int    `json:"IdCiudad"`
	IdProvincia int    `json:"IdProvincia"`
	IdPais      int    `json:"IdPais"`
	Ciudad      string `json:"Ciudad"`
}

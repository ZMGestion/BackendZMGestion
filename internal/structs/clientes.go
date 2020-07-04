package structs

type Clientes struct {
	IdCliente       int    `json:"IdCliente"`
	IdPais          string `json:"IdPais"`
	IdTipoDocumento int    `json:"IdTipoDocumento"`
	Documento       string `json:"Documento"`
	Tipo            string `json:"Tipo"`
	FechaNacimiento string `json:"FechaNacimiento"`
	Nombres         string `json:"Nombres"`
	Apellidos       string `json:"Apellidos"`
	RazonSocial     string `json:"RazonSocial"`
	Email           string `json:"Email"`
	Telefono        string `json:"Telefono"`
	FechaAlta       string `json:"FechaAlta"`
	FechaBaja       string `json:"FechaBaja"`
	Estado          string `json:"Estado"`
}

package structs

type Clientes struct {
	IdCliente       int    `json:"IdCliente"`
	IdPais          string `json:"IdPais"`
	IdTipoDocumento int    `json:"IdTipoDocumento"`
	Documento       string `json:"Documento"`
	Tipo            string `json:"Tipo"`
	FechaNacimiento string `json:"FechaNacimiento"`
	Nombres         string `json:"Nombres,omitempty"`
	Apellidos       string `json:"Apellidos,omitempty"`
	RazonSocial     string `json:"RazonSocial,omitempty"`
	Email           string `json:"Email"`
	Telefono        string `json:"Telefono"`
	FechaAlta       string `json:"FechaAlta"`
	FechaBaja       string `json:"FechaBaja,omitempty"`
	Estado          string `json:"Estado"`
}

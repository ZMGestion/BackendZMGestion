package structs

type Usuarios struct {
	IdUsuario       int    `json:"IdUsuario,omitempty"`
	IdRol           int    `json:"IdRol,omitempty"`
	IdUbicacion     int    `json:"IdUbicacion,omitempty"`
	IdTipoDocumento int    `json:"IdTipoDocumento,omitempty"`
	Documento       string `json:"Documento"`
	Nombres         string `json:"Nombres"`
	Apellidos       string `json:"Apellidos"`
	EstadoCivil     string `json:"EstadoCivil,omitempty"`
	Telefono        string `json:"Telefono,omitempty"`
	Email           string `json:"Email,omitempty"`
	CantidadHijos   int    `json:"CantidadHijos"`
	Usuario         string `json:"Usuario,omitempty"`
	Password        string `json:"Password,omitempty"`
	Token           string `json:"Token,omitempty"`
	FechaUltIntento string `json:"FechaUltIntento,omitempty"`
	Intentos        int    `json:"Intentos,omitempty"`
	FechaNacimiento string `json:"FechaNacimiento,omitempty"`
	FechaInicio     string `json:"FechaInicio,omitempty"`
	FechaAlta       string `json:"FechaAlta,omitempty"`
	FechaBaja       string `json:"FechaBaja,omitempty"`
	Estado          string `json:"Estado,omitempty"`
}

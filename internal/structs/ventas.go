package structs

//Ventas Ventas
type Ventas struct {
	IdVenta       int    `json:"IdVenta,omitempty"`
	IdCliente     int    `json:"IdCliente,omitempty"`
	IdDomicilio   int    `json:"IdDomicilio,omitempty"`
	IdUbicacion   int    `json:"IdUbicacion,omitempty"`
	IdUsuario     int    `json:"IdUsuario,omitempty"`
	FechaAlta     string `json:"FechaAlta,omitempty"`
	Observaciones string `json:"Observaciones,omitempty"`
	Estado        string `json:"Estado,omitempty"`
}

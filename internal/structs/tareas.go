package structs

type Tareas struct {
	IdTarea             int    `json:"IdTarea,omitempty"`
	IdLineaProducto     int    `json:"IdLineaProducto,omitempty"`
	IdTareaSiguiente    int    `json:"IdTareaSiguiente,omitempty"`
	IdUsuarioFabricante int    `json:"IdUsuarioFabricante,omitempty"`
	IdUsuarioRevisor    int    `json:"IdUsuarioRevisor,omitempty"`
	Tarea               string `json:"Tarea,omitempty"`
	FechaInicio         string `json:"FechaInicio,omitempty"`
	FechaPausa          string `json:"FechaPausa,omitempty"`
	FechaFinalizacion   string `json:"FechaFinalizacion,omitempty"`
	FechaRevision       string `json:"FechaRevision,omitempty"`
	FechaAlta           string `json:"FechaAlta,omitempty"`
	FechaCancelacion    string `json:"FechaCancelacion,omitempty"`
	Observaciones       string `json:"Observaciones,omitempty"`
	Estado              string `json:"Estado,omitempty"`
}

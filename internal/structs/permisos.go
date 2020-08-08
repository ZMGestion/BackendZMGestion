package structs

type Permisos struct {
	IdPermiso   int    `json:"IdPermiso,omitempty"`
	Permiso     string `json:"Permiso,omitempty"`
	Descripcion string `json:"Descripcion,omitempty"`
}

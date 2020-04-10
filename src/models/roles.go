package models

type Roles struct {
	IdRol         uint32 `json:"IdRol"`
	Rol           string `json:"Rol"`
	FechaAlta     string `json:"FechaAlta"`
	Observaciones string `json:"Observaciones"`
}

func (r *Roles) Dame() (*Roles, error) {
	return r, nil
}

func (r *Roles) DarAlta() (string, error) {
	return "", nil
}

func (r *Roles) DarBaja() (string, error) {
	return "", nil
}

package models

import "time"

type Roles struct {
	IdRol         int       `json:"IdRol"`
	Rol           string    `json:"Rol"`
	FechaAlta     time.Time `json:"FechaAlta"`
	Observaciones string    `json:"Observaciones"`
}

func (r *Roles) toMap() map[string]interface{} {
	return map[string]interface{}{
		"IdRol": r.IdRol,
		"Rol": r.Rol,
		"FechaAlta": r.FechaAlta,
		"Observaciones": r.Observaciones
	}
}

func (r *Roles) fromMap(m Map[string]interface{}){
	r.IdRol = m["IdRol"]
	r.Rol = m["Rol"]
	r.FechaAlta = m["FechaAlta"]
	r.Observaciones = m["Observaciones"]
}

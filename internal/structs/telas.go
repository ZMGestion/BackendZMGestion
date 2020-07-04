package structs

type Telas struct {
	IdTela        int    `json:"IdTela"`
	Tela          string `json:"Tela"`
	FechaAlta     string `json:"FechaAlta"`
	FechaBaja     string `json:"FechaBaja"`
	Observaciones string `json:"Observaciones"`
	Estado        string `json:"Estado"`
}

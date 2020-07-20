package structs

type Telas struct {
	IdTela        int    `json:"IdTela"`
	Tela          string `json:"Tela"`
	FechaAlta     string `json:"FechaAlta,omitempty"`
	FechaBaja     string `json:"FechaBaja,omitempty"`
	Observaciones string `json:"Observaciones"`
	Estado        string `json:"Estado"`
}

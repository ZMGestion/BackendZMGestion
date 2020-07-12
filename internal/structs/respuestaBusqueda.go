package structs

type RespuestaBusqueda struct {
	Resultado    []map[string]interface{} `json:"resultado"`
	Paginaciones *Paginaciones            `json:"Paginaciones"`
}

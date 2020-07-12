package structs

type RespuestaBusqueda struct {
	Resultado    []map[string]interface{} `json:"Resultado"`
	Paginaciones *Paginaciones            `json:"Paginaciones"`
}

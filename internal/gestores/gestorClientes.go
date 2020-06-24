package gestores

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/structs"
	"encoding/json"

	"github.com/mitchellh/mapstructure"
)

//GestorUsuarios commented
type GestorClientes struct {
	DbHandler *db.DbHandler
}

//Crear Usuario
func (gc *GestorClientes) Crear(cliente structs.Clientes, domicilio structs.Domicilios, token string) (*structs.Clientes, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Clientes":        cliente,
		"UsuariosEjecuta": usuarioEjecuta,
		"Domicilios":      domicilio,
	}

	out, err := gc.DbHandler.CallSP("zsp_cliente_crear", params)

	if err != nil || out == nil {
		return nil, err
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	if err == nil {
		if response["Clientes"] != nil {
			err = mapstructure.Decode(response["Clientes"], &cliente)
		} else {
			return nil, nil
		}
	}

	return &cliente, err
}

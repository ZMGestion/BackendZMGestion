package gestores

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/structs"
	"encoding/json"

	"github.com/mitchellh/mapstructure"
)

//GestorDomicilios
type GestorDomicilios struct {
	DbHanlder *db.DbHandler
}

func (gd *GestorDomicilios) Crear(domicilio structs.Domicilios, cliente structs.Clientes, token string) (*structs.Domicilios, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Domicilios":      domicilio,
		"Clientes":        cliente,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	out, err := gd.DbHanlder.CallSP("zsp_domicilio_crear", params)

	if err != nil || out == nil {
		return nil, err
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	if err == nil {
		if response["Domicilios"] != nil {
			err = mapstructure.Decode(response["Domicilios"], &domicilio)
		} else {
			return nil, nil
		}
	}

	return &domicilio, err
}

func (gd *GestorDomicilios) Borrar(domicilio structs.Domicilios, cliente structs.Clientes, token string) error {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Domicilios":      domicilio,
		"UsuariosEjecuta": usuarioEjecuta,
		"Clientes":        cliente,
	}

	_, err := gd.DbHanlder.CallSP("zsp_domicilio_borrar", params)

	if err != nil {
		return err
	}
	return nil
}

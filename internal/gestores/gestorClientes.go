package gestores

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/structs"
	"encoding/json"
	"fmt"

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

func (gc *GestorClientes) Modificar(cliente structs.Clientes, token string) (*structs.Clientes, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Clientes":        cliente,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	out, err := gc.DbHandler.CallSP("zsp_cliente_modificar", params)

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

//Crear Usuario
func (gc *GestorClientes) Borrar(cliente structs.Clientes, token string) error {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Clientes":        cliente,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	_, err := gc.DbHandler.CallSP("zsp_cliente_borrar", params)

	if err != nil {
		return err
	}
	return nil
}

// Buscar
func (gc *GestorClientes) Buscar(cliente structs.Clientes, token string) ([]*structs.Clientes, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Clientes":        cliente,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	out, err := gc.DbHandler.CallSP("zsp_clientes_buscar", params)

	if err != nil {
		return nil, err
	}

	if out == nil {
		return nil, nil
	}

	var response []map[string]interface{}

	err = json.Unmarshal(*out, &response)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	var clientes []*structs.Clientes
	for _, el := range response {
		var cliente structs.Clientes
		if el["Clientes"] != nil {
			err = mapstructure.Decode(el["Clientes"], &cliente)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, nil
		}
		clientes = append(clientes, &cliente)
	}

	return clientes, nil
}

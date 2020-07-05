package gestores

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/structs"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

//GestorUsuarios commented
type GestorUsuarios struct {
	DbHandler *db.DbHandler
}

//Crear Usuario
func (gu *GestorUsuarios) Crear(usuario structs.Usuarios, token string) (*structs.Usuarios, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Usuarios":        usuario,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	out, err := gu.DbHandler.CallSP("zsp_usuario_crear", params)

	if err != nil || out == nil {
		return nil, err
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	if err == nil {
		if response["Usuarios"] != nil {
			err = mapstructure.Decode(response["Usuarios"], &usuario)
		} else {
			return nil, nil
		}
	}

	return &usuario, err
}

//Modificar Usuario
func (gu *GestorUsuarios) Modificar(usuario structs.Usuarios, token string) (*structs.Usuarios, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Usuarios":        usuario,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	out, err := gu.DbHandler.CallSP("zsp_usuario_modificar", params)

	if err != nil || out == nil {
		return nil, err
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	if err == nil {
		if response["Usuarios"] != nil {
			err = mapstructure.Decode(response["Usuarios"], &usuario)
		} else {
			return nil, nil
		}
	}

	return &usuario, err
}

//ModificarPassword Usuario
func (gu *GestorUsuarios) ModificarPassword(passActual string, passNueva string, token string) (*structs.Usuarios, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	usuarioActual := structs.Usuarios{
		Password: passActual,
	}

	usuarioNuevo := structs.Usuarios{
		Password: passNueva,
	}

	params := map[string]interface{}{
		"UsuariosActual":  usuarioActual,
		"UsuariosNuevo":   usuarioNuevo,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	out, err := gu.DbHandler.CallSP("zsp_usuario_modificar_pass", params)

	if err != nil || out == nil {
		return nil, err
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	if err == nil {
		if response["Usuarios"] != nil {
			err = mapstructure.Decode(response["Usuarios"], &usuarioNuevo)
		} else {
			return nil, nil
		}
	}

	return &usuarioNuevo, err
}

//Borrar Usuario
func (gu *GestorUsuarios) Borrar(usuario structs.Usuarios, token string) error {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Usuarios":        usuario,
		"UsuariosEjecuta": usuarioEjecuta,
	}

	_, err := gu.DbHandler.CallSP("zsp_usuario_borrar", params)

	if err != nil {
		return err
	}
	return nil

}

func (gu *GestorUsuarios) Buscar(usuario structs.Usuarios, token string) ([]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}
	params := map[string]interface{}{
		"Usuarios":        usuario,
		"UsuariosEjecuta": usuarioEjecuta,
	}
	out, err := gu.DbHandler.CallSP("zsp_usuarios_buscar", params)

	if err != nil {
		return nil, err
	}

	if out == nil {
		return nil, nil
	}

	var response []map[string]interface{}

	err = json.Unmarshal(*out, &response)

	if err != nil {
		return nil, err
	}

	var respuesta []interface{}

	for _, el := range response {

		if el["Usuarios"] != nil && el["Roles"] != nil && el["Ubicaciones"] != nil {
			err = mapstructure.Decode(el["Usuarios"], &usuario)
			if err != nil {
				return nil, err
			}
			var roles structs.Roles
			err = mapstructure.Decode(el["Roles"], &roles)
			if err != nil {
				return nil, err
			}
			var ubicaciones structs.Ubicaciones
			err = mapstructure.Decode(el["Ubicaciones"], &ubicaciones)
			if err != nil {
				return nil, err
			}
			objetos := map[string]interface{}{
				"Usuarios":    usuario,
				"Roles":       roles,
				"Ubicaciones": ubicaciones,
			}
			respuesta = append(respuesta, objetos)
		} else {
			return nil, nil
		}
	}

	return respuesta, nil
}

func (gu *GestorUsuarios) ListarTiposDocumento() ([]*structs.TiposDocumento, error) {
	out, err := gu.DbHandler.CallSP("zsp_tiposDocumento_listar", nil)

	if err != nil {
		return nil, err
	}

	if out == nil {
		return nil, errors.New("ERROR_DEFAULT")
	}

	var response []map[string]interface{}

	err = json.Unmarshal(*out, &response)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	var tiposDocumento []*structs.TiposDocumento
	for _, el := range response {
		var tipoDocumento structs.TiposDocumento
		if el["TiposDocumento"] != nil {
			err = mapstructure.Decode(el["TiposDocumento"], &tipoDocumento)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, nil
		}
		tiposDocumento = append(tiposDocumento, &tipoDocumento)
	}

	return tiposDocumento, nil
}

package models

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/helpers"
	"BackendZMGestion/internal/structs"
	"encoding/json"
	"errors"

	"github.com/mitchellh/mapstructure"
)

type UsuariosService struct {
	Usuario   *structs.Usuarios //Es lo que se recibe de request
	DbHandler *db.DbHandler
}

func (s *UsuariosService) Dame(token string) (interface{}, error) {

	usuariosEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"Usuarios":        s.Usuario,
		"UsuariosEjecuta": usuariosEjecuta,
	}

	out, err := s.DbHandler.CallSP("zsp_usuario_dame", params)

	if err != nil {
		return nil, err
	}

	if out == nil {
		return nil, errors.New("ERROR_DEFAULT")
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	if err != nil {
		return nil, nil
	}

	var objetos interface{}
	if response["Usuarios"] != nil && response["Roles"] != nil && response["Ubicaciones"] != nil {
		var usuario structs.Usuarios
		err = mapstructure.Decode(response["Usuarios"], &usuario)
		if err != nil {
			return nil, err
		}
		var roles structs.Roles
		err = mapstructure.Decode(response["Roles"], &roles)
		if err != nil {
			return nil, err
		}
		var ubicaciones structs.Ubicaciones
		err = mapstructure.Decode(response["Ubicaciones"], &ubicaciones)
		if err != nil {
			return nil, err
		}
		objetos = map[string]interface{}{
			"Usuarios":    usuario,
			"Roles":       roles,
			"Ubicaciones": ubicaciones,
		}
	} else {
		return nil, nil
	}
	return objetos, nil
}

//Dame por token
func (s *UsuariosService) DamePorToken(token string) (*structs.Usuarios, error) {

	usuariosEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuariosEjecuta,
	}

	out, err := s.DbHandler.CallSP("zsp_usuario_dame_por_token", params)

	if err != nil {
		return nil, err
	}

	if out == nil {
		return nil, errors.New("ERROR_DEFAULT")
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	if err != nil {
		return nil, nil
	}
	var usuario structs.Usuarios
	if response["Usuarios"] != nil {
		err = mapstructure.Decode(response["Usuarios"], &usuario)
	} else {
		return nil, nil
	}
	if err != nil {
		return nil, nil
	}
	return &usuario, nil
}

//Dar alta
func (s *UsuariosService) DarAlta(token string) (*structs.Usuarios, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"Usuarios":        s.Usuario,
	}

	out, err := s.DbHandler.CallSP("zsp_usuario_dar_alta", params)

	if err != nil {
		return nil, err
	}

	if out == nil {
		return nil, errors.New("ERROR_DEFAULT")
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	if err != nil {
		return nil, nil
	}
	var usuario structs.Usuarios
	if response["Usuarios"] != nil {
		err = mapstructure.Decode(response["Usuarios"], &usuario)
	} else {
		return nil, nil
	}
	if err != nil {
		return nil, nil
	}
	return &usuario, nil
}

//Dar alta
func (s *UsuariosService) DarBaja(token string) (*structs.Usuarios, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"Usuarios":        s.Usuario,
	}

	out, err := s.DbHandler.CallSP("zsp_usuario_dar_baja", params)

	if err != nil {
		return nil, err
	}

	if out == nil {
		return nil, errors.New("ERROR_DEFAULT")
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	if err != nil {
		return nil, nil
	}
	var usuario structs.Usuarios
	if response["Usuarios"] != nil {
		err = mapstructure.Decode(response["Usuarios"], &usuario)
	} else {
		return nil, nil
	}
	if err != nil {
		return nil, nil
	}
	return &usuario, nil
}

func (s *UsuariosService) RestablecerPassword(token string) error {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"Usuarios":        s.Usuario,
	}

	_, err := s.DbHandler.CallSP("zsp_usuario_restablecer_pass", params)

	if err != nil {
		return err
	}

	return nil

}

func (s *UsuariosService) IniciarSesion() (*structs.Usuarios, error) {

	out, err := s.DbHandler.CallSP("zsp_sesion_iniciar", helpers.GenerateJSONFromModels(s.Usuario))

	if err != nil {
		return nil, err
	}

	if out == nil {
		return nil, errors.New("ERROR_DEFAULT")
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	if err != nil {
		return nil, nil
	}
	var usuario structs.Usuarios
	if response["Usuarios"] != nil {
		err = mapstructure.Decode(response["Usuarios"], &usuario)
	} else {
		return nil, nil
	}
	if err != nil {
		return nil, nil
	}
	return &usuario, nil
}

func (s *UsuariosService) CerrarSesion(token string) error {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"Usuarios":        s.Usuario,
	}

	_, err := s.DbHandler.CallSP("zsp_sesion_cerrar", params)

	if err != nil {
		return err
	}

	return nil

}

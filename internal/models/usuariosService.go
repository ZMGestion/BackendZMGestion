package models

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/structs"
	"encoding/json"
	"errors"

	"github.com/mitchellh/mapstructure"
)

type UsuariosService struct {
	Usuario   *structs.Usuarios //Es lo que se recibe de request
	DbHandler *db.DbHandler
}

func (s *UsuariosService) Dame(token string) (*structs.Usuarios, error) {

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

	var usuario structs.Usuarios
	if err == nil {
		if response["Usuarios"] != nil {
			err = mapstructure.Decode(response["Usuarios"], &usuario)
		} else {
			return nil, nil
		}
	}
	return &usuario, err
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

	var usuario structs.Usuarios
	if err == nil {
		if response["Usuarios"] != nil {
			err = mapstructure.Decode(response["Usuarios"], &usuario)
		} else {
			return nil, nil
		}
	}
	return &usuario, err
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

	var usuario structs.Usuarios
	if err == nil {
		if response["Usuarios"] != nil {
			err = mapstructure.Decode(response["Usuarios"], &usuario)
		} else {
			return nil, nil
		}
	}
	return &usuario, err
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

	var usuario structs.Usuarios
	if err == nil {
		if response["Usuarios"] != nil {
			err = mapstructure.Decode(response["Usuarios"], &usuario)
		} else {
			return nil, nil
		}
	}
	return &usuario, err
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

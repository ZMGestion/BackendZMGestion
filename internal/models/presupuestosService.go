package models

import (
	"BackendZMGestion/internal/db"
	"BackendZMGestion/internal/structs"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/smtp"
	"os"

	"github.com/jordan-wright/email"
)

//PresupuestosService PresupuestosService
type PresupuestosService struct {
	DbHanlder    *db.DbHandler
	Presupuestos structs.Presupuestos
}

//Dame Dame
func (ps *PresupuestosService) Dame(token string) (map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"Presupuestos":    ps.Presupuestos,
	}

	out, err := ps.DbHanlder.CallSP("zsp_presupuesto_dame", params)

	if err != nil {
		return nil, err
	}

	if out == nil {
		return nil, errors.New("ERROR_DEFAULT")
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	return response, err
}

//PasarACreado PasarACreado
func (ps *PresupuestosService) PasarACreado(token string) (map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"Presupuestos":    ps.Presupuestos,
	}

	out, err := ps.DbHanlder.CallSP("zsp_presupuesto_pasar_a_creado", params)

	if err != nil {
		return nil, err
	}

	if out == nil {
		return nil, errors.New("ERROR_DEFAULT")
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	return response, err
}

//CrearLineaPresupuesto CrearLineaPresupuesto
func (ps *PresupuestosService) CrearLineaPresupuesto(lineaProducto structs.LineasProducto, productoFinal structs.ProductosFinales, token string) (map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta":  usuarioEjecuta,
		"LineasProducto":   lineaProducto,
		"ProductosFinales": productoFinal,
	}

	out, err := ps.DbHanlder.CallSP("zsp_lineaPresupuesto_crear", params)

	if err != nil {
		return nil, err
	}

	if out == nil {
		return nil, errors.New("ERROR_DEFAULT")
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	return response, err
}

//ModificarLineaPresupuesto ModificarLineaPresupuesto
func (ps *PresupuestosService) ModificarLineaPresupuesto(lineaProducto structs.LineasProducto, productoFinal structs.ProductosFinales, token string) (map[string]interface{}, error) {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta":  usuarioEjecuta,
		"LineasProducto":   lineaProducto,
		"ProductosFinales": productoFinal,
	}

	out, err := ps.DbHanlder.CallSP("zsp_lineaPresupuesto_modificar", params)

	if err != nil {
		return nil, err
	}

	if out == nil {
		return nil, errors.New("ERROR_DEFAULT")
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)

	return response, err
}

//BorrarLineaPresupuesto BorrarLineaPresupuesto
func (ps *PresupuestosService) BorrarLineaPresupuesto(lineaProducto structs.LineasProducto, token string) error {
	usuarioEjecuta := structs.Usuarios{
		Token: token,
	}

	params := map[string]interface{}{
		"UsuariosEjecuta": usuarioEjecuta,
		"LineasProducto":  lineaProducto,
	}

	_, err := ps.DbHanlder.CallSP("zsp_lineaPresupuesto_borrar", params)

	return err
}

// EnviarEmail Funcion para enviar por mail un presupuesto
func (ps *PresupuestosService) EnviarEmail(cliente structs.Clientes, presupuesto multipart.FileHeader, token string) error {

	pdf, err := presupuesto.Open()
	if err != nil {
		return err
	}
	defer pdf.Close()

	archivo, err := os.Create(presupuesto.Filename)
	if err != nil {
		return err
	}
	defer archivo.Close()

	if _, err = io.Copy(archivo, pdf); err != nil {
		return err
	}

	email := &email.Email{
		From:    "Zimmerman Muebles" + "<" + os.Getenv("email") + ">",
		To:      []string{cliente.Email},
		Subject: "Presupuesto Zimmerman Muebles",
	}
	email.AttachFile(archivo.Name())
	return email.Send("smtp.gmail.com:587", smtp.PlainAuth("", os.Getenv("email"), os.Getenv("pass"), "smtp.gmail.com"))

}

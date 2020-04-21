package enums

import (
	"bytes"
	"encoding/json"
)

// Status representa el estado de la request
type Status string

const (
	// Cuanto la request fue exitosa
	OK Status = "Ok"
	//Cuando ocurrio un error
	ERROR Status = "Error"
)

var toString = map[Status]string{
	OK:    "Ok",
	ERROR: "Error",
}

var toID = map[string]Status{
	"Ok":    OK,
	"Error": ERROR,
}

// MarshalJSON marshals the enum as a quoted json string
func (s Status) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (s *Status) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*s = toID[j]
	return nil
}

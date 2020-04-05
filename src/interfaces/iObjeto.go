package interfaces

type IObjeto interface {
	toMap() map[string]interface{}
	fromMap(m map[string]interface{})
}

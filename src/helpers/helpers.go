package helpers

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/labstack/echo"
)

// URLParamInt retorna un parámetro de la url como int
func URLParamInt(c echo.Context, key string) (int, error) {
	return strconv.Atoi(c.Param(key))
}

func GetModelName(clase interface{}) string {
	tipo := reflect.TypeOf(clase)
	var className string
	if tipo != nil {
		className = fmt.Sprintf("%s", tipo)
		className = strings.Split(className, "structs.")[1]
	}

	return className
}

func GenerateJSONFromModels(elements ...interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	for _, el := range elements {
		if name := GetModelName(el); name != "" {
			res[name] = el
		}
	}
	return res
}

func GenerateMapFromContext(c echo.Context) (map[string]interface{}, error) {
	var jsonMap map[string]interface{}
	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	if err != nil {
		return nil, err
	}
	return jsonMap, nil
}

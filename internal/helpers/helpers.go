package helpers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
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

//GetError genera un error HTTP que puede ser mostrado.
func GetError(err error) (string, string) {
	//Se deberia recibir tambien el codigo de lenguaje
	errorsFile, _ := ioutil.ReadFile("errores_ES.json")
	var errorsMap map[string]string
	_ = json.Unmarshal([]byte(errorsFile), &errorsMap)
	errorCode := "ERROR_DEFAULT"
	errorMsg := errorsMap[errorCode]
	if err != nil {
		txt := err.Error()
		if txt[0:6] == "ERROR_" {
			if _, ok := errorsMap[err.Error()]; ok {
				errorCode = err.Error()
				errorMsg = errorsMap[err.Error()]
			}
		}
	}
	return errorCode, errorMsg
}

func Hash(password string) *string {

	hash := sha256.New()
	hash.Write([]byte(password))

	passwordHash := hex.EncodeToString(hash.Sum(nil))

	fmt.Println(passwordHash)

	return &passwordHash
}

func CreateToken(user string) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user,
		"exp":  time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		"iat":  time.Now().Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("app-key")))

	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}

func GetToken(token string) (*string, error) {

	split := strings.SplitN(token, " ", 2)

	if split[0] != "JWT" {
		return nil, errors.New("ERROR_DEFAULT")
	}
	return &split[1], nil
}

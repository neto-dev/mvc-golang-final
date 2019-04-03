package utils

/*
Importamos las librerias

We import libraries
*/
import (
	"net/http"

	"github.com/mvc-golang-final/models"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

/*
Declaramos la estructura que utilizaremos para retornar las respuestas a los clientes

Declare the structure we will use to return responses to customers
*/
type ResponseJSON struct {
	Result interface{} `json:"results"`
}

/*
Declaramos la estructura que utilizaremos para retornar las repuestas de error a los clientes

Declare the structure we will use to return error responses to customers
*/
type ErrorJSON struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	MysqlError string `json:"mysql_error"`
}

/*
Funcion para retornar la respuesta en JSON

Function to return the answer in JSON
*/
func ReturnRespondJSON(ctx echo.Context, code int, payload interface{}) error {
	r := ResponseJSON{payload}

	return ctx.JSON(code, r)
}

/*
Funcion para retornar la respuesta de error en formato JSON

Function to return the error response in JSON format
*/
func ReturnErrorJSON(ctx echo.Context, errorMsj string, err error) error {
	r := ErrorJSON{getStatusCode(err), errorMsj, err.Error()}

	return ctx.JSON(getStatusCode(err), r)
}

/*
Funcion para retornar el error http resultante

Function to return the resulting http error
*/
func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
		case models.INTERNAL_SERVER_ERROR:

			return http.StatusInternalServerError
		case models.NOT_FOUND_ERROR:
			return http.StatusNotFound
		case models.CONFLIT_ERROR:
			return http.StatusConflict
		default:
			return http.StatusInternalServerError
	}
}
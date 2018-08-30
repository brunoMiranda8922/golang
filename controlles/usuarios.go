package controlles

import (
	"net/http"

	"github.com/labstack/echo"
)

//Callback é página inicial da aplicação
func Callback(content echo.Context) error {
	menssage := "Hello Golang"
	returr := content.String(http.StatusOK, menssage)
	return returr

}

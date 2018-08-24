package routers

import (
	"net/http"

	"github.com/labstack/echo"
)

var App *echo.Echo

func init() {
	App = echo.New()

	App.GET("/", callback)
}

func callback(content echo.Context) error {
	menssage := "Hello Golang"
	returr := content.String(http.StatusOK, menssage)
	return returr

}

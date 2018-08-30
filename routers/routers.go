package routers

import (
	"github.com/Projeto/usuarios/controlles"
	"github.com/labstack/echo"
)

//App é uma instância de echo
var App *echo.Echo

func init() {
	App = echo.New()

	App.GET("/", controlles.Callback)

	api := App.Group("/v1/")

	api.POST("/inserir", controlles.Inserir)

}

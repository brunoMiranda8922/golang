package main

import (
	"github.com/Projeto/usuarios/routers"
	"github.com/labstack/echo/middleware"
)

func main() {
	serverUp := routers.App
	serverUp.Use(middleware.Logger())
	serverUp.Logger.Fatal(serverUp.Start(":3000"))
}

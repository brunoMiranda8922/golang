package main

import (
	"github.com/Projeto/usuarios/routers"
)

func main() {
	//aula 16
	routers.App.Start(":3000")
}

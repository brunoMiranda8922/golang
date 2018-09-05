package controlles

import (
	"net/http"

	"github.com/Projeto/usuarios/models"
	"github.com/labstack/echo"
)

//Callback é página inicial da aplicação
func Callback(content echo.Context) error {
	menssage := "hello"
	returr := content.String(http.StatusOK, menssage)
	return returr

}

//Inserir new user
func Inserir(i echo.Context) error {
	nome := i.FormValue("nome")
	email := i.FormValue("email")

	var usuario models.Usuarios
	usuario.Nome = nome
	usuario.Email = email

	if nome != "" && email != "" {
		if _, err := models.UsuariosModel.Insert(usuario); err != nil {
			return i.JSON(http.StatusBadRequest, map[string]string{
				"mensagem": "Não é possível adicionar o registro no banco!",
			})
		}

		return i.JSON(http.StatusCreated, map[string]string{
			"mensagem": "Gravado com sucesso",
		})
	}

	return i.JSON(http.StatusBadRequest, map[string]string{
		"menssagem": "OS campos precisam ser completados!",
	})
}

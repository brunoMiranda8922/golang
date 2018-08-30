package lib

import (
	"log"

	"upper.io/db.v3"
	"upper.io/db.v3/mysql"
)

var conexao = mysql.ConnectionURL{
	Host:     "localhost",
	User:     "bruno",
	Password: "",
	Database: "membros",
}

//Sess é uma variavel que faz a conexão com banco de dados
var Sess db.Database

func init() {
	var err error

	Sess, err = mysql.Open(conexao)
	if err != nil {
		log.Fatal(err.Error())
	}
}

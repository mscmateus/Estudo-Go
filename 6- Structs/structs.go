package main

import "fmt"

type Usuario struct {
	nome     string
	idade    uint8
	endereco Endereco
}
type Endereco struct {
	rua    string
	numero uint16
}

func main() {
	endereco := Endereco{"severino", 213}
	var usuario1 Usuario = Usuario{nome: "nome", idade: 21}
	usuario2 := Usuario{"nome", 21, endereco}
	var usuario3 Usuario
	usuario3.nome = "jão"
	usuario3.idade = 21
	fmt.Println(usuario1, usuario2, usuario3)

}

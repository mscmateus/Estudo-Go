package main

import "fmt"

func main() {
	var variavel1 int = 10
	var variavel2 int = variavel1
	fmt.Println(variavel2)

	var variavel3 int
	var ponteiro *int
	variavel3 = 100
	ponteiro = &variavel3 //criando a referecia

	fmt.Println(variavel3, *ponteiro) //desfazendo a referencia
}

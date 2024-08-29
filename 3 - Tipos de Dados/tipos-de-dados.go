package main

import (
	"errors"
	"fmt"
)

func main() {
	// int8 int32 int16 int64
	var numero int64 = -1000000000000000000
	fmt.Println(numero)
	var numero2 uint = 10000
	fmt.Println(numero2)

	//alias
	//int32 = rune
	var n3 rune = 12456
	fmt.Println(n3)

	//uint8 = byte
	var n4 byte = 123
	fmt.Println(n4)

	//float32 float64
	var nreal1 float32 = 123.4
	fmt.Println(nreal1)

	var nreal2 float64 = 12300011111100.4
	fmt.Println(nreal2)

	nreal3 := 12300011111100.4
	fmt.Println(nreal3)

	//Strings

	var str string = "texto"
	fmt.Println(str)

	str2 := "texto"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)

	//valor "zero"
	var texto string
	fmt.Println(texto)

	var booleanol bool
	fmt.Println(booleanol)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}

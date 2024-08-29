package main

import "fmt"

func main() {
	var array1 [5]string

	array1[0] = "possicao 1"
	fmt.Println(array1)
	array2 := [5]string{"1", "2", "3", "4", "5"}
	fmt.Println(array2)
	array3 := [...]string{"1", "2", "3", "4", "5"}
	fmt.Println(array3)

	slice := []int{10, 11, 22, 33, 66, 55, 44}
	fmt.Println(slice)

	slice = append(slice, 155)

	slice2 := slice[1:3]
	fmt.Println(slice2)

	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)

}

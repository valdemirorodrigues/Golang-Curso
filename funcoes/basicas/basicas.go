package main

import "fmt"

func main() {
	imprimir()
}
func imprimir() {
	fmt.Println(soma(10, 3))

}

func soma(nota1 int, nota2 int) int {
	resultado := nota1 + nota2
	return resultado
}

//funcao retornando string

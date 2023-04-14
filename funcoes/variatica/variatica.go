package main

import "fmt"

var total float64

func main() {
	media(6.7, 8.9, 7.1, 9.5)
	fmt.Println(media(6.7, 8.9, 7.1, 9.5))

}

/*
Uma função recebe essa nomenclatura quando pode receber um número indefinido de parâmetros,
para se declarar uma função desse tipo usamos …
seguido do tipo que esperamos.
Exemplo:
*/

func media(valor ...float64) float64 {
	for _, valores := range valor {
		total += valores

	}
	return total / float64(len(valor))
}

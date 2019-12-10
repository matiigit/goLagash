/*
*	Definir un vector de 5 elementos de tipo entero en la función main.
*	Implementar una función que retorne el mayor y el menor elemento del vector
*	por medio de dos parámetros de tipo puntero:
 */

package main

import "fmt"

func mayorMenor(vec [5]int, pmayor *int, pmenor *int) {
	*pmayor = vec[0]
	*pmenor = vec[0]

	for _, value := range vec {
		if value < *pmenor {
			*pmenor = value
		} else {
			if value > *pmayor {
				*pmayor = value
			}
		}
	}
}

func main() {
	vec := [5]int{10, 22, 3, 44, 12}
	var pmayor, pmenor int

	fmt.Println("Valor menor main=", pmenor)
	fmt.Println("Valor mayor main=", pmayor)

	mayorMenor(vec, &pmayor, &pmenor)

	fmt.Println("Menor elemento=", pmenor)
	fmt.Println("Mayor elemento=", pmayor)

}

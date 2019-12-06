package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	var array [5]int

	for i := 0; i < len(array); i++ {

		for {
			var aux bool = true
			array[i] = rand.Intn(len(array))

			for j := 0; j < len(array); j++ {
				if i != j && array[i] == array[j] {
					aux = false
					break
				}
			}

			if aux {
				break
			}
		}
	}

	position := 0
	var jumpTotal float64 = 0

	for i := 0; i < len(array); i++ {
		//Cantidad de saltos a realizar
		jump := math.Abs(float64(array[position]) - float64(position))
		//Suma de la cantidad de saltos
		jumpTotal += jump
		//
		fmt.Println("De:", position, "=>", array[position], "\tc/saltos:", jump)
		position = array[position]
	}
	fmt.Println("Saltos totales:", jumpTotal)
	fmt.Println(array)
}

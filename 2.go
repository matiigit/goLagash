package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func mod(x int, y int) int {
	if x >= 0 {
		for x >= 0 {
			if x >= y {
				x -= y
			} else {
				break
			}
		}
		return x
	} else {
		for x < 0 {
			if x < y {
				x += y
			}
		}
		return mod(x, y)
	}
}

func main() {
	var dimension int
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Scan(&dimension)
	array := make([]int, dimension)
	stop := make([]bool, dimension)

	for i := 0; i < len(array); i++ {
		array[i] = r.Intn(101) - 50
		stop[i] = false
	}

	var (
		nextPosition   int = 0
		actualPosition int
		step           int = 1
	)

	fmt.Println("Vector =", array)
	for {
		actualPosition = nextPosition // Guardo la posicion en la que estoy parado

		nextPosition += array[nextPosition]
		nextPosition = mod(nextPosition, len(array))

		fmt.Println(
			"De", actualPosition,
			"==>", nextPosition,
			"\tSe caminaron:", math.Abs(float64(array[actualPosition])), "pasos")

		if stop[nextPosition] || nextPosition == 0 {
			break
		}

		stop[nextPosition] = true
		step++ // Guardo la cantidad de pasos totales
	}
	fmt.Println("Se dieron", step, "iteraciones.")
}

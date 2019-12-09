package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func Mod(x int, y int) int {
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
		return Mod(x, y)
	}
}

func Ejercicio2(array []int) {
	var (
		nextPosition   int = 0
		actualPosition int
		step           int = 1
	)
	stop := make([]bool, len(array))

	fmt.Println("Vector =", array)
	for {
		actualPosition = nextPosition // Guardo la posicion en la que estoy parado

		nextPosition += array[nextPosition]
		nextPosition = Mod(nextPosition, len(array))

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

func main() {
	path := "input.txt"
	byteArray, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println("Archivo inválido en la ruta: '", path, "' ", err.Error())
		os.Exit(1)
	}

	content := string(byteArray)

	stringArray := strings.Split(content, "\n")

	road := make([][]int, len(stringArray))

	for i, v := range stringArray {
		aux := strings.Split(v, ",")

		for _, value := range aux {
			auxNumber, err := strconv.Atoi(value)

			if err != nil {
				auxNumber, err = strconv.Atoi((value[:len(value)-1]))

				if err != nil {
					fmt.Println("Caracter no valido:", err.Error())
					os.Exit(1)
				}
			}

			road[i] = append(road[i], auxNumber)
		}
	}

	for _, v := range road {
		Ejercicio2(v)
	}

	os.Exit(0)
}

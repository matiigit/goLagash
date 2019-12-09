package main

import "fmt"

func main() {
	fmt.Println("Hola Mundo 🌍!")

	myMap := make(map[string]float32)

	myMap["Mauro"] = 10
	myMap["Matías"] = 3.34
	myMap["Ivan"] = 9.999999999999999999999999999
	myMap["🕺"] = 9.999999999999999999999999999

	for key, value := range myMap {
		fmt.Printf("%v: %0.2f\n", key, value)
	}

	// value, exists := myMap["keyquenoexiste"]
	// if exists {
	// 	fmt.Printf("%v: %0.2f\n", "keyquenoexiste", value)
	// }

	// key := "keyquenoexiste"
	key := "🕺"

	if value, exists := myMap[key]; exists {
		fmt.Printf("%v: %0.2f\n", key, value)
	}
}

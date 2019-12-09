package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
)

// Reverse da vuelta un string
func Reverse(s string) string {
	var r string

	for i := len(s); i > 0; i-- {
		r += string(s[i-1 : i])
	}

	return r
}

// Delete Borra caracteres especiales
func Delete(s string) string {

	for _, v := range s {
		// if !(v >= 48 && v <= 57 || v >= 65 && v <= 90 || v >= 97 && v <= 122) {
		// 	s = strings.Replace(s, string(v), "", 1)
		// }
		if !(unicode.IsNumber(v) || unicode.IsLetter(v)) {
			s = strings.Replace(s, string(v), "", 1)
		}
	}
	return s
}

func main() {
	path := "palabrotas.txt"
	byteArray, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println("Archivo no valido:", path, err.Error())
		os.Exit(1)
	}

	content := strings.Split(string(byteArray), ",")

	palindromos := make(map[string]int)

	for _, value := range content {

		value = Delete(value)

		if _, exist := palindromos[value]; exist {
			palindromos[value]++
		} else {
			palindromos[value] = 1
		}
	}

	for key, value := range palindromos {

		if _, exist := palindromos[Reverse(key)]; exist {
			fmt.Println(
				key, "Es palindromo.\n",
				"Se repite", value, "veces.")
		} else {
			fmt.Println(
				key, "no es palindromo.\n",
				"Se repite", value, "veces.")
		}
		fmt.Println()
	}

	os.Exit(0)
}

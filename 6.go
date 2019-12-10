package main

import (
	"fmt"
)

type Pais struct {
	nombre      string
	cantidadhab int
}

func load(pais *Pais) {
	fmt.Print("Nombre pais:")
	fmt.Scan(&pais.nombre)
	fmt.Print("Cantidad de habitantes:")
	fmt.Scan(&pais.cantidadhab)
}

func show(pais Pais) {
	fmt.Print(pais.nombre, ":", pais.cantidadhab)
}

func main() {
	var miPais Pais
	load(&miPais)
	show(miPais)
}

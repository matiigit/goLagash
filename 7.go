package main

import (
	"fmt"
	"os"
	"os/exec"
)

type Usuario struct {
	nombre string
	clave  string
}

type Administrador struct {
	Usuario
	privilegio int
}

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (usu *Usuario) cargar(nombre, clave string) {
	usu.nombre = nombre
	usu.clave = clave
}

func (usu Usuario) imprimir() {
	fmt.Println(
		"Nombre\t", usu.nombre, "\n",
		"Clave\t", usu.clave)
}

func (admin *Administrador) cargar(nombre, clave string, privilegio int) {
	admin.Usuario.cargar(nombre, clave)
	admin.privilegio = privilegio
	clear()
}

func (admin Administrador) imprimirPrivilegio() {
	admin.imprimir()
	fmt.Println("Privilegio\t", admin.privilegio)
}

func main() {
	var administrador1 Administrador
	administrador1.cargar("juan", "abc123", 1)
	administrador1.imprimirPrivilegio()
}

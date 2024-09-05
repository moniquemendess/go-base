package main

import "fmt"

const prefixoOlaPortugues = "Hola "

// creei una función Ola string
func Ola(nome string) string {
	return prefixoOlaPortugues + nome
}

// Llame la funcion Ola mundo
func main() {
	fmt.Println(Ola("mundo"))
}

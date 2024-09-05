package main

import "fmt"

// creei una función Ola string
func Ola(nome string) string {
	return "Hola " + nome
}

// Llame la funcion Ola
func main() {
	fmt.Println(Ola("mundo"))
}

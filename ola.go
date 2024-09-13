package main

import "fmt"

const espanhol = "espanhol"
const frances = "frances"
const prefixoOlaFrances = "Bonjour, "
const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "

// creei una función Ola string con dos argumentos
func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "mundo"

	}
	if idioma == espanhol {
		return prefixoOlaEspanhol + nome
	}
	if idioma == frances {
		return prefixoOlaFrances + nome
	}
	return prefixoOlaPortugues + nome
}

// Llame la funcion Ola mundo
func main() {
	fmt.Println(Ola("mundo", "espanhol"))
}

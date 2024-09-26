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
		nome = "Mundo"
	}

	return prefixodeSaudacao(idioma) + nome
}

func prefixodeSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case frances:
		prefixo = prefixoOlaFrances
	case espanhol:
		prefixo = prefixoOlaEspanhol
	default:
		prefixo = prefixoOlaPortugues
	}
	return
}

// Llame la funcion Ola mundo
func main() {
	fmt.Println(Ola("Monique", "frances"))
}

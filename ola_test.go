package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola("mundo")
	esperado := "Hola mundo"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}

func TestOlaChris(t *testing.T) {
	resultado := Ola("Chris")
	esperado := "Hola Chris"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}

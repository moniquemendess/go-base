package main

import (

	"reflect"
	"testing"
)

func TestSoma(t *testing.T) {


	t.Run("coleção de qualquer numero", func(t *testing.T) {
		numeros := []int{1, 2, 3}

		resultado := Soma(numeros)
		esperado := 6

		if resultado != esperado {
			t.Errorf("resultado %d, esperado %d, dado, %v", resultado, esperado, numeros)
		}
	})

}

func TestSomaTodoOResto(t *testing.T) {

	resultado := SomaTodoOResto([]int{1, 2}, []int{0, 9})
	esperado := []int{3, 9}

	if !reflect.DeepEqual(resultado, esperado) {
		t.Errorf("recebido %v esperado %v", resultado, esperado)

	}
}

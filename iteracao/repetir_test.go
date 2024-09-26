package iteracao

import "testing"

func TestRepetir(t *testing.T) {
	repeticoes := Repetir("a")
	esperado := "aaaaa"

	if repeticoes != esperado {
		t.Errorf("repeticoes '%s', mas obteve '%s'", repeticoes, esperado)

	}
}

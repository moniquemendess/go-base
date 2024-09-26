package inteiros

import "testing"

func TestAdicionador(t *testing.T) {
	soma := Adiciona(6, 6)
	esperado := 12

	if soma != esperado {
		t.Errorf("esperado '%d', resultado '%d'", esperado, soma)
	}
}

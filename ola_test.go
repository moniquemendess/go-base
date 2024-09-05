package main

import "testing"

func TestOla(t *testing.T) {
	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}

	}
	t.Run("diz hola para las personas ", func(t *testing.T) {

		resultado := Ola("Chris")
		esperado := "Hola Chris"
		verificaMensagemCorreta(t, resultado, esperado)
	})
	t.Run("'Mundo' como padrão para 'string' vazia", func(t *testing.T) {

		resultado := Ola("")
		esperado := "Hola mundo"
		verificaMensagemCorreta(t, resultado, esperado)
	})

}

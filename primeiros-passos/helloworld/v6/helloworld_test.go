package main

import "testing"

func TestOla(t *testing.T) {

	verificarSaidaTeste := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}

	// Apresentado os subtestes
	t.Run("diz olá para as pessoas", func(t *testing.T) {
		resultado := Ola("Moraski", "")
		esperado := "Olá, Moraski"
		verificarSaidaTeste(t, resultado, esperado)
	})

	t.Run("diz 'Olá mundo' quando uma string vazia for passada", func(t *testing.T) {
		resultado := Ola("", "")
		esperado := "Olá, mundo"
		verificarSaidaTeste(t, resultado, esperado)
	})

	// Agora em outra lingua
	t.Run("diz olá para as pessoas em espanhol", func(t *testing.T) {
		resultado := Ola("Moraski", "espanhol")
		esperado := "Hola, Moraski"
		verificarSaidaTeste(t, resultado, esperado)
	})

	t.Run("diz olá para as pessoas em frances", func(t *testing.T) {
		resultado := Ola("Moraski", "frances")
		esperado := "Bonjour, Moraski"
		verificarSaidaTeste(t, resultado, esperado)
	})

	t.Run("diz olá para as pessoas em japones", func(t *testing.T) {
		resultado := Ola("Moraski", "japones")
		esperado := "Ohayo, Moraski"
		verificarSaidaTeste(t, resultado, esperado)
	})

	t.Run("diz olá para as pessoas em ingles", func(t *testing.T) {
		resultado := Ola("Moraski", "ingles")
		esperado := "Hello, Moraski"
		verificarSaidaTeste(t, resultado, esperado)
	})
}

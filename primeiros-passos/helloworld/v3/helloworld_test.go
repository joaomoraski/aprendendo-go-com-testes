package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola("Moraski")
	esperado := "Olá, Moraski"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}

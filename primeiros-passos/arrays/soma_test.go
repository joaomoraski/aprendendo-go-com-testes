package arrays

import "testing"

func TestSoma(t *testing.T) {
	numeros := [5]int{1, 2, 3, 4, 5}

	resultado := Soma(numeros)
	esperado := 15

	if resultado != esperado {
		t.Errorf("Esperado '%s', resultado '%s', dado %v", esperado, resultado, numeros)
	}
}

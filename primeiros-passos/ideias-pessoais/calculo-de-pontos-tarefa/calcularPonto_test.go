package calculodepontostarefa

import "testing"

func TestCalcularPonto(t *testing.T) {
	numeros := []int{5, 5, 13, 21, 3}

	resultado := Soma(numeros)
	esperado := 47

	if resultado != esperado {
		t.Errorf("Esperado '%d', resultado '%d', dado %v", esperado, resultado, numeros)
	}
}

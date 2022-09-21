package arrays

import "testing"

func TestSoma(t *testing.T) {

	t.Run("Array com 5 elementos", func(t *testing.T) {
		numeros := []int{1, 2, 3, 4, 5}

		resultado := Soma(numeros)
		esperado := 15

		if resultado != esperado {
			t.Errorf("Esperado '%d', resultado '%d', dado %v", esperado, resultado, numeros)
		}
	})

	// t.Run("Slice de qualquer tamanho", func(t *testing.T) {
	// 	numeros := []int{1, 2, 3, 4, 5, 6}

	// 	resultado := Soma(numeros)
	// 	esperado := 21

	// 	if resultado != esperado {
	// 		t.Errorf("Esperado '%d', resultado '%d', dado %v", esperado, resultado, numeros)
	// 	}
	// })

}

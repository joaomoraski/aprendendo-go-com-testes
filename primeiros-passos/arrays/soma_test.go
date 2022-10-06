package arrays

import (
	"reflect"
	"testing"
)

func TestSoma(t *testing.T) {

	t.Run("Array com 5 elementos", func(t *testing.T) {
		numeros := []int{1, 2, 3, 4, 5}

		resultado := Soma(numeros)
		esperado := 15

		if resultado != esperado {
			t.Errorf("Esperado '%d', resultado '%d', dado %v", esperado, resultado, numeros)
		}
	})
	// go test -cover para ver o nivel de cobertura dos testes
	// nesse caso era redundante ter 2 testes, ja que o codigo foi refatorado para usar apenas slices

	// t.Run("Slice de qualquer tamanho", func(t *testing.T) {
	// 	numeros := []int{1, 2, 3, 4, 5, 6}

	// 	resultado := Soma(numeros)
	// 	esperado := 21

	// 	if resultado != esperado {
	// 		t.Errorf("Esperado '%d', resultado '%d', dado %v", esperado, resultado, numeros)
	// 	}
	// })
}

func TestSomaTudo(t *testing.T) {

	resultado := SomaTudo([]int{1, 2}, []int{0, 9})
	esperado := []int{3, 9}
	// deep equal nao diferencia tipo
	if !reflect.DeepEqual(resultado, esperado) {
		t.Errorf("Esperado '%v', resultado '%v'", esperado, resultado)
	}
}

func TestSomaTodoOResto(t *testing.T) {

	t.Run("faz a soma de alguns slices", func(t *testing.T) {
		resultado := SomaTodoOResto([]int{1, 2}, []int{0, 9})
		esperado := []int{2, 9}
		// deep equal nao diferencia tipo
		// verifica se dois objetos sao iguais
		if !reflect.DeepEqual(resultado, esperado) {
			t.Errorf("Esperado '%v', resultado '%v'", esperado, resultado)
		}
	})

	t.Run("soma slices de maneira segura", func(t *testing.T) {
		resultado := SomaTodoOResto([]int{}, []int{3, 4, 5})
		esperado := []int{0, 9}
		// deep equal nao diferencia tipo
		// verifica se dois objetos sao iguais
		if !reflect.DeepEqual(resultado, esperado) {
			t.Errorf("Esperado '%v', resultado '%v'", esperado, resultado)
		}
	})
}

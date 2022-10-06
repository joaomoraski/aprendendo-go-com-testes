package main

import "testing"

func TestPerimetro(t *testing.T) {
	t.Run("Perimetro: retangulo", func(t *testing.T) {
		retangulo := Retangulo{10.0, 10.0}
		resultado := Perimetro(retangulo)
		esperado := 40.0

		if resultado != esperado {
			t.Errorf("Resultado: %.2f\n esperado: %.2f", resultado, esperado)
		}
	})
}

func TestArea(t *testing.T) {
	t.Run("Area: retangulo", func(t *testing.T) {
		retangulo := Retangulo{12.0, 6.0}
		resultado := Area(retangulo)
		esperado := 72.0

		if resultado != esperado {
			t.Errorf("Resultado: %.2f\n esperado: %.2f", resultado, esperado)
		}
	})
}

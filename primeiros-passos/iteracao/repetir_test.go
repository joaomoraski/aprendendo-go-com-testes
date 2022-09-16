package iteracao

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepetir(t *testing.T) {
	repeticoes := Repetir("a", 10)
	esperado := strings.Repeat("a", 10)

	if repeticoes != esperado {
		t.Errorf("Esperado '%s', resultado '%s'", esperado, repeticoes)
	}
}

func ExampleRepetir() {
	soma := Repetir("a", 5)
	fmt.Println(soma)
	// Output: aaaaa
}

// testing.B dara acesso ao b.N
// Quando ele roda ele executa o b N vezes e mede quando tempo leva
// go test -bench=.
// Por padrao é executa sequencialmente
// ns/op é quantos nanossegundos para ser executado.
/*
	goos: linux
	goarch: amd64
	pkg: github.com/joaomoraski/aprendendo-go-com-testes/primeiros-passos/iteracao
	cpu: Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz
	BenchmarkRepetir-8   	10263254	       107.8 ns/op
	PASS
	ok  	github.com/joaomoraski/aprendendo-go-com-testes/primeiros-passos/iteracao	1.227s
*/
func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a", 10)
	}
}

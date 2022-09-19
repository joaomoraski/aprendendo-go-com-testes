package calculodepontostarefa

import "fmt"

func Soma(numeros []int) (resultado int) {
	// usar o range para ajudar a limpar o codigo
	// _ blank identifier para ignorar o valor do range
	for _, numero := range numeros {
		resultado += numero
	}
	return
}

func CalcularPonto() (resultado int) {
	var (
		numeroTarefas int = 0
		pontos        int
		numeros       []int
	)
	fmt.Println("Informe o número de tarefas: ")
	fmt.Scanf("%d", numeroTarefas)
	i := 0
	for i < numeroTarefas {
		fmt.Scanf("%d", pontos)
		numeros = append(numeros, pontos)
	}
	resultado = Soma(numeros)
	return
}

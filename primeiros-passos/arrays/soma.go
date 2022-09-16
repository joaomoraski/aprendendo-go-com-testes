package arrays

func Soma(numeros [5]int) (resultado int) {
	// usar o range para ajudar a limpar o codigo
	// _ blank identifier para ignorar o valor do range
	for _, numero := range numeros {
		resultado += numero
	}
	return
}

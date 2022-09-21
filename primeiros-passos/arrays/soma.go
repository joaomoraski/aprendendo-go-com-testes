package arrays

func Soma(numeros []int) (resultado int) {
	// usar o range para ajudar a limpar o codigo
	// _ blank identifier para ignorar o valor do indice
	for _, numero := range numeros {
		resultado += numero
	}
	return
}

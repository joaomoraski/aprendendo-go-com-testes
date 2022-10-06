package arrays

func Soma(numeros []int) (resultado int) {
	// usar o range para ajudar a limpar o codigo
	// _ blank identifier para ignorar o valor do indice
	for _, numero := range numeros {
		resultado += numero
	}
	return
}

// func SomaTudo(numerosParaSomar ...[]int) (somas []int) {
// 	somas = make([]int, len(numerosParaSomar))

// 	for i, numeros := range numerosParaSomar {
// 		somas[i] = Soma(numeros)
// 	}

// 	return
// }

// Refactoring
func SomaTudo(numerosParaSomar ...[]int) []int {
	// cria um slice vazio
	var somas []int
	// ignora o indice do array e percorre cada slice passado
	for _, numeros := range numerosParaSomar {
		// append passa o slice e adiciona o valor passado na frente,
		// nesse caso a soma de cada slice passado
		somas = append(somas, Soma(numeros))
	}
	return somas
}

// Soma todo o resto
// Slices podem ser fatiados, usando a sintaxe
// slice[inicio:final], se vc omitir o valor e algum dos lados
// ele vai capturar todo o lado omitido
// No caso de 1: ele fala para pegar da posiçao 1 ate o final
func SomaTodoOResto(numerosParaSomar ...[]int) []int {
	// cria um slice vazio
	var somas []int
	// ignora o indice do array e percorre cada slice passado
	for _, numeros := range numerosParaSomar {
		final := numeros[1:]
		// append passa o slice e adiciona o valor passado na frente,
		// nesse caso a soma de cada slice passado
		somas = append(somas, Soma(final))
	}
	return somas
}

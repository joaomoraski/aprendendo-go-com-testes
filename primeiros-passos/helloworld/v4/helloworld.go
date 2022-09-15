package main

import "fmt"

// constantes em alguns casos podem melhorar a performance do codigo
// Este nao é um deles kkkkkk
const prefixoOlaPortugues = "Olá, "

// separação da regra de negocio
func Ola(nome string) string {
	return prefixoOlaPortugues + nome
}

func main() {
	fmt.Println(Ola("mundo"))
}

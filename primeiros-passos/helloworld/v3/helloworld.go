package main

import "fmt"

// separação da regra de negocio
func Ola(nome string) string {
	return "Olá, " + nome
}

func main() {
	fmt.Println(Ola("mundo"))
}

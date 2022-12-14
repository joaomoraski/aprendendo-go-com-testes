package main

import "fmt"

// constantes em alguns casos podem melhorar a performance do codigo
// Este nao é um deles kkkkkk
const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaFrances = "Bonjour, "
const prefixoOlaJapones = "Ohayo, "
const prefixoOlaIngles = "Hello, "

// separação da regra de negocio
func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "mundo"
	}
	prefixo := prefixoOlaPortugues
	switch idioma {
	case "frances":
		prefixo = prefixoOlaFrances
	case "espanhol":
		prefixo = prefixoOlaEspanhol
	case "japones":
		prefixo = prefixoOlaJapones
	case "ingles":
		prefixo = prefixoOlaIngles

	}
	return prefixo + nome
}

func main() {
	fmt.Println(Ola("mundo", ""))
}

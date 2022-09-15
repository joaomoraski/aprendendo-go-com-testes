package main

import "fmt"

// Este nao é um deles kkkkkk
const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaFrances = "Bonjour, "
const prefixoOlaJapones = "Ohayo, "
const prefixoOlaIngles = "Hello, "

func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "mundo"
	}
	return prefixoSaudacao(idioma) + nome
}

func prefixoSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case "frances":
		prefixo = prefixoOlaFrances
	case "espanhol":
		prefixo = prefixoOlaEspanhol
	case "japones":
		prefixo = prefixoOlaJapones
	case "ingles":
		prefixo = prefixoOlaIngles
	default:
		prefixo = prefixoOlaPortugues
	}
	return
}

func main() {
	fmt.Println(Ola("mundo", ""))
}

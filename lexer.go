package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

// import (
// 	lex "github.com/timtadh/lexmachine"
// 	"github.com/timtadh/lexmachine/machines"
// )

// var Literals []string // The tokens representing literal strings
var Keywords []string

var Tokens []string // All of the tokens (including literals and keywords)
var Tipos []string

// var TokenIds map[string]int // A map from the token names to their int ids
// var Lexer *lex.Lexer // The lexer object. Use this to construct a Scanner

func main() {
	lexTokens := []string{}

	arquivo, err := ioutil.ReadFile("teste.gon")
	str := string(arquivo)

	if err != nil {
		fmt.Println(err)
	}

	stringue := ""
	for i := 0; i < len(str); i++ {
		caracter := string([]rune(str)[i])
		if caracter != " " {
			stringue = stringue + caracter
		} else {
			lexTokens = append(lexTokens, stringue)
			stringue = ""
		}
	}
	if stringue != "" {
		lexTokens = append(lexTokens, stringue)
	}

	initTokens()
	verificarTokens(str, criarRegras())
}

func initTokens() {
	Keywords = []string{
		"seGonzo",
		"entaoGonzo",
		"senaoGonzo",
		"enquantoGonzo",
		"paraGonzo",
		"gonzoIn",
		"gonzoOut",
		"-!GONZOSTART!-",
		"-!GONZOEND!-",
		"fazGonzo",
		"acaboGonzo",
	}
	Tipos = []string{
		"enio",
		"beto",
		"piggy",
		"caco",
	}
	Tokens = append(Tokens, Keywords...)
	Tokens = append(Tokens, Tipos...)
}

func verificarTokens(texto string, regras []string) {

	for i := 0; i < len(regras); i++ {
		r, _ := regexp.Compile(regras[i])
		fmt.Println(r.FindAllString(texto, -1))
	}

}

func criarRegras() []string {
	regras := []string{}

	// Primeira regra: Declarações de variaveis

	declaracao := "\\$("
	for j := 0; j < len(Tipos); j++ {
		if j > 0 {
			declaracao += "|"
		}
		declaracao += "(" + Tipos[j] + ")"
	}
	declaracao += ") mimimi[a-zA-Z0-9]+"

	regras = append(regras, declaracao)
	return regras
}

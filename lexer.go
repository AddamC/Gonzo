package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

// import (
// 	lex "github.com/timtadh/lexmachine"
// 	"github.com/timtadh/lexmachine/machines"
// )

// var Literals []string // The tokens representing literal strings
var Keywords []string
var TiposTokens []string
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

	words := strings.Fields(str)

	initTokens()
	verificarTokens(words, criarRegras())
}

func initTokens() {
	TiposTokens = []string{
		"Tipo de Dados",
		"Variável",
		"Atribuição",
		"Instruçaõ de Entrada ou Saida de Dados",
	}
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

func verificarTokens(texto []string, regras []string) {
	fContador := 0
	for i := 0; i < len(texto); i++ {
		falha := true
		for j := 0; j < len(regras); j++ {
			r, _ := regexp.Compile(regras[j])
			if r.MatchString(texto[i]) {
				fmt.Println(strconv.Itoa(i+1) + " - " + texto[i] + " -> " + TiposTokens[j])
				falha = false
			}
		}
		if falha {
			fContador++
			fmt.Println(strconv.Itoa(i+1) + " - " + texto[i] + " -> Falha")
		}
	}
	fmt.Println("Quantidade de Falhas: " + strconv.Itoa(fContador))

}

func criarRegras() []string {
	regras := []string{}

	// Primeira regra: Declarações de variaveis

	declaracao := "^\\$("
	for j := 0; j < len(Tipos); j++ {
		if j > 0 {
			declaracao += "|"
		}
		declaracao += "(" + Tipos[j] + ")"
	}
	declaracao += ")$"
	// declaracao += ") mimimi[a-zA-Z0-9]+"

	regras = append(regras, declaracao)

	declaracao = "^mimimi[a-zA-Z0-9]+$"

	regras = append(regras, declaracao)

	declaracao = "^@$"

	regras = append(regras, declaracao)
	//^(gonzoIn\((mimimi[a-zA-Z0-9]+)\)|gonzoOut\((mimimi[a-zA-Z0-9]+|[0-9]+|"[a-zA-Z0-9]*")\))$
	declaracao = "^(gonzoIn\\((mimimi[a-zA-Z0-9]+)\\)|gonzoOut\\((mimimi[a-zA-Z0-9]+|[0-9]+|\"[a-zA-Z0-9]*\")\\))$"

	regras = append(regras, declaracao)

	return regras
}

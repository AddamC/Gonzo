package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// import (
// 	lex "github.com/timtadh/lexmachine"
// 	"github.com/timtadh/lexmachine/machines"
// )

var Keywords []string
var TiposTokens []string
var Tokens []string // All of the tokens (including literals and keywords)
var Tipos []string
var Symbols []string
var Textos []string

// var TokenIds map[string]int // A map from the token names to their int ids
// var Lexer *lex.Lexer // The lexer object. Use this to construct a Scanner

func main() {
	//lexTokens := []string{}
	arquivo, err := ioutil.ReadFile("teste.gon")
	str := string(arquivo)

	if err != nil {
		fmt.Println(err)
	}

	words := strings.Fields(verificarTexto(str))

	initTokens()
	//initFile()
	verificarTokens(str, words, criarRegras())
	//endingFile()
}

func verificarTokens(text string, words []string, regras []string) {
	f, _ := os.Create("result.html")
	f.WriteString("<html> <head> </head> <body> ")
	f.WriteString("<center>")
	f.WriteString("<h1 style=\"color: magenta;\"> Analisador Léxigo - GoNzooooooooo</h1>")
	f.WriteString("<img src=\"gonzo.jpg\"> ")
	f.WriteString("</center>")
	f.WriteString("<ul>")
	fContador := 0
	sContador := 0
	tokensAceito := []string{}
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(regras); j++ {
			r, _ := regexp.Compile(regras[j])
			str := r.FindString(words[i])
			if str != "" {
				words[i] = strings.Replace(words[i], str, "", -1)
				if j == 1 {
					f.WriteString("<li>" + Textos[sContador] + " -> " + TiposTokens[j] + "</li>")
					sContador++
				} else {
					f.WriteString("<li>" + str + " -> " + TiposTokens[j] + "</li>")
				}

				tokensAceito = append(tokensAceito, str)
				j = -1
			}
		}
		if words[i] != "" {
			fContador++
			f.WriteString("<li style=\"color: red;\">" + words[i] + " -> Falha </li>")
		}
	}
	f.WriteString("</ul>")
	f.WriteString("<h1> Quantidade de Falhas: " + strconv.Itoa(fContador) + "</h1>")
	f.WriteString("<img src=\"pedrosolaandandodemoto.gif\">")
	f.WriteString("</body> </html>")
	os.Open("result.html")
}

func verificarTexto(text string) string {
	r, _ := regexp.Compile("\"[\\w\\ \\s.?!]*\"")
	Textos = r.FindAllString(text, -1)
	for i := 0; i < len(Textos); i++ {
		text = strings.Replace(text, Textos[i], "string", -1)
	}
	return text
}

func initTokens() {
	TiposTokens = []string{
		"Simbolos",
		"Texto",
		"Numeral",
		"Tipo de Dados",
		"Variável",
		"Palavras Reservadas",
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
		"gonzoIn",
		"gonzoOut",
	}
	Symbols = []string{
		"GG_GONZO",
		"<",
		">",
		"<gonzo",
		"\\|",
		"\\(",
		"\\)",
		"==",
		"@",
		"GONZADD",
		"GONZSUB",
		"GONZDIV",
		"GONZMULT",
		"^eGonzo$",
		"ouGonzo",
	}
	Tipos = []string{
		"enio",
		"beto",
		"piggy",
		"caco",
	}
	Tokens = append(Tokens, Keywords...)
	Tokens = append(Tokens, Tipos...)
	Tokens = append(Tokens, Symbols...)
}

func criarRegras() []string {
	regras := []string{}
	declaracao := "("
	for i := 0; i < len(Symbols); i++ {
		if i > 0 {
			declaracao += "|"
		}
		declaracao += Symbols[i]
	}
	declaracao += ")"
	regras = append(regras, declaracao)
	declaracao = "string"
	regras = append(regras, declaracao)
	declaracao = "[0-9]+[.]{0,1}[0-9]*"
	regras = append(regras, declaracao)
	declaracao = "\\$("
	for j := 0; j < len(Tipos); j++ {
		if j > 0 {
			declaracao += "|"
		}
		declaracao += "(" + Tipos[j] + ")"
	}
	declaracao += ")"
	regras = append(regras, declaracao)
	declaracao = "mimimi[a-zA-Z0-9]+"
	regras = append(regras, declaracao)
	declaracao = "("
	for i := 0; i < len(Keywords); i++ {
		if i > 0 {
			declaracao += "|"
		}
		declaracao += Keywords[i]
	}
	declaracao += ")"
	regras = append(regras, declaracao)
	return regras
}

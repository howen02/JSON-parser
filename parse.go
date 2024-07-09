package main

import (
	"log"
	"os"
)

func ValidateJSON(file *os.File) bool {
	l := NewLexer(file)
	tokens := []Token{}
	braces := []string{}

	// tokenise the input file
	for {
		_, tok, _ := l.Lex()
		tokens = append(tokens, tok)

		if tok == EOF {
			break
		}

		if tok == LCURL || tok == RCURL {
			braces = append(braces, tok.toString())
		}
	}

	// validate braces
	if len(braces)%2 != 0 || len(braces) == 0 {
		log.Println("Invalid braces")
		return false
	}

	for i, tok := range tokens {
		if tok == RCURL {
			if tokens[i-1] == COMMA {
				log.Println("Comma before closing curly brace")
				return false
			}
		}

		if tok == COMMA && tokens[i+1] != DQUOTE && !tokens[i+1].isValue() {
			log.Println("Comma not followed by a key")
			return false
		}

		if tok == COLON {
			if tokens[i-1] != DQUOTE && !tokens[i-1].isValue() {
				log.Println("Invalid key")
				return false
			}

			if tokens[i+1] != DQUOTE && !tokens[i+1].isValue() {
				log.Println("Invalid value")
				return false
			}

		}

		if tok == IDENT && (tokens[i-1] != DQUOTE || tokens[i+1] != DQUOTE) {
			log.Println("Identifier not wrapped in quotes")
			return false
		}

		if tok == SQUOTE && tokens[i+2] != SQUOTE && tokens[i-2] != SQUOTE {
			log.Println("Single quote not closed")
			return false
		}
	}

	return true
}

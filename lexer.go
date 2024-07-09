package main

import (
	"bufio"
	"io"
	"unicode"
)

type Token int

const (
	EOF = iota
	ILLEGAL
	IDENT
	INT
	SEMI
	COLON
	COMMA
	BOOL
	NULL

	ADD
	SUB
	MUL
	DIV
	ASSIGN

	LCURL
	RCURL
	LPAR
	RPAR
	LBRAC
	RBRAC
	DQUOTE
	SQUOTE
)

var tokens = []string{
	EOF:     "EOF",
	ILLEGAL: "ILLEGAL",
	IDENT:   "IDENT",
	INT:     "INT",
	BOOL:    "BOOL",
	NULL:    "NULL",
	SEMI:    ";",
	COLON:   ":",
	COMMA:   ",",

	ADD:    "+",
	SUB:    "-",
	MUL:    "*",
	DIV:    "/",
	ASSIGN: "=",

	LCURL:  "{",
	RCURL:  "}",
	LPAR:   "(",
	RPAR:   ")",
	LBRAC: "[",
	RBRAC: "]",
	DQUOTE: "\"",
	SQUOTE: "'",
}

func (t Token) toString() string {
	return tokens[t]
}

type Pos struct {
	line int
	col  int
}

type Lexer struct {
	pos    Pos
	reader *bufio.Reader
}

func NewLexer(r io.Reader) *Lexer {
	return &Lexer{
		pos:    Pos{1, 1},
		reader: bufio.NewReader(r),
	}
}

func (l *Lexer) Lex() (Pos, Token, string) {
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return l.pos, EOF, ""
			}

			panic(err)
		}

		l.pos.col++

		switch r {
		case '\n':
			l.resetPosition()
		case ';':
			return l.pos, SEMI, ";"
		case '+':
			return l.pos, ADD, "+"
		case '-':
			return l.pos, SUB, "-"
		case '*':
			return l.pos, MUL, "*"
		case '/':
			return l.pos, DIV, "/"
		case '=':
			return l.pos, ASSIGN, "="
		case '{':
			return l.pos, LCURL, "{"
		case '}':
			return l.pos, RCURL, "}"
		case '(':
			return l.pos, LPAR, "("
		case ')':
			return l.pos, RPAR, ")"
		case '[':
			return l.pos, LBRAC, "["
		case ']':
			return l.pos, RBRAC, "]"
		case '"':
			return l.pos, DQUOTE, "\""
		case '\'':
			return l.pos, SQUOTE, "'"
		case ':':
			return l.pos, COLON, ":"
		case ',':
			return l.pos, COMMA, ","
		default:
			if unicode.IsSpace(r) {
				continue
			}

			if unicode.IsDigit(r) {
				start := l.pos
				l.backup()
				lit := l.lexInt()

				return start, INT, lit
			}

			if unicode.IsLetter(r) {
				start := l.pos
				l.backup()
				lit := l.lexIdent()

				switch lit {
				case "true", "false":
					return start, BOOL, lit
				case "null":
					return start, NULL, lit
				}

				return start, IDENT, lit
			}
		}
	}
}

func (l *Lexer) resetPosition() {
	l.pos.line++
	l.pos.col = 0
}

func (l *Lexer) backup() {
	if err := l.reader.UnreadRune(); err != nil {
		panic(err)
	}

	l.pos.col--
}

func (l *Lexer) lexInt() string {
	var lit string

	for {
		r, _, err := l.reader.ReadRune()
		if err == io.EOF {
			return lit
		}

		l.pos.col++
		if unicode.IsDigit(r) {
			lit += string(r)
		} else {
			l.backup()
			return lit
		}
	}
}

func (l *Lexer) lexIdent() string {
	var lit string

	for {
		r, _, err := l.reader.ReadRune()
		if err == io.EOF {
			return lit
		}

		l.pos.col++
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '-' || unicode.IsSpace(r) {
			lit += string(r)
		} else {
			l.backup()
			return lit
		}
	}
}

func (t Token) isValue() bool {
	return t == INT || t == BOOL || t == NULL || t == IDENT || t == LCURL || t == LBRAC
}

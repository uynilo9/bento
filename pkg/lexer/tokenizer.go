package lexer

import (
	"regexp"

	"github.com/uynilo9/bento/pkg/logger"
)

type (
	regexHandler func (
		tokenizer *tokenizer,
		regex *regexp.Regexp,
	)
	regexPattern struct {
		regex *regexp.Regexp
		handler regexHandler
	}

	tokenizer struct {
		Tokens []Token
		patterns []regexPattern
		position uint
		source string
	}
)

func (tokenizer * tokenizer) at() string {
	return string(tokenizer.source[tokenizer.position])
}

func (tokenizer *tokenizer) atEOF() bool {
	return tokenizer.position >= uint(len(tokenizer.source))
}

func (tokenizer *tokenizer) remainder() string {
	return tokenizer.source[tokenizer.position:]
}

func (tokenizer *tokenizer) push(token Token) {
	tokenizer.Tokens = append(tokenizer.Tokens, token)
}

func (tokenizer *tokenizer) advance(length uint) {
	tokenizer.position += length
}


func Tokenize(source string) []Token {
	tokenizer := newTokenizer(source)
	
	for !tokenizer.atEOF() {
		var matched bool
		for _, pattern := range tokenizer.patterns {
			location := pattern.regex.FindStringIndex(tokenizer.remainder())
			if location != nil && location[0] == 0 {
				pattern.handler(tokenizer, pattern.regex)
				matched = true
				break
			}
		}
		if !matched {
			logger.Fatalf("Unrecognized token `%s` while tokenizing in the Lexer\n", tokenizer.at())
		}
	}
	tokenizer.push(NewToken(EOF, ""))
	return tokenizer.Tokens
}

func defaultHandler(kind TokenKind, value string) regexHandler {
	return func(tokenizer *tokenizer, regex *regexp.Regexp) {
		tokenizer.advance(uint(len(value)))
		tokenizer.push(NewToken(kind, value))
	}
}

func numberHandler(tokenizer *tokenizer, regex *regexp.Regexp) {
	matched := regex.FindString(tokenizer.remainder())
	tokenizer.push(NewToken(NUMBER, matched))
	tokenizer.advance(uint(len(matched)))
}

func newTokenizer(source string) *tokenizer {
	return &tokenizer{
		Tokens: make([]Token, 0),
		patterns: []regexPattern{
			{regexp.MustCompile(`[0-9]+(\.[0-9]+)?`), numberHandler},

			{regexp.MustCompile(`\(`), defaultHandler(LROUND, "(")}, {regexp.MustCompile(`\)`), defaultHandler(RROUND, ")")},
			{regexp.MustCompile(`\[`), defaultHandler(LSQUARE, "[")}, {regexp.MustCompile(`\]`), defaultHandler(RSQUARE, "]")},
			{regexp.MustCompile(`\{`), defaultHandler(LCURLY, "{")}, {regexp.MustCompile(`\}`), defaultHandler(RCURLY, "}")},

			{regexp.MustCompile(`==`), defaultHandler(EQUALTO, "==")}, {regexp.MustCompile(`!=`), defaultHandler(NEQUALTO, "!=")},
			{regexp.MustCompile(`<`), defaultHandler(LESS, "<")}, {regexp.MustCompile(`<=`), defaultHandler(LEQUALTO, "<=")},
			{regexp.MustCompile(`>`), defaultHandler(GREATER, ">")}, {regexp.MustCompile(`>=`), defaultHandler(GEQUALTO, ">=")},
			{regexp.MustCompile(`!`), defaultHandler(NOT, "!")},
			
			{regexp.MustCompile(`\+\+`), defaultHandler(PLUSPLUS, "++")}, {regexp.MustCompile(`\+`), defaultHandler(PLUS, "+")},
			{regexp.MustCompile(`--`), defaultHandler(MINUSMINUS, "--")}, {regexp.MustCompile(`-`), defaultHandler(MINUS, "-")},
			{regexp.MustCompile(`\*\*`), defaultHandler(TIMESTIMES, "**")}, {regexp.MustCompile(`\*`), defaultHandler(TIMES, "*")},
			{regexp.MustCompile(`\/\/`), defaultHandler(SLASHSLASH, "//")}, {regexp.MustCompile(`/`), defaultHandler(DIVIDEDBY, "/")},
			{regexp.MustCompile(`%`), defaultHandler(MODULO, "%")},			

			{regexp.MustCompile(`=`), defaultHandler(EQUALS, "=")},
			{regexp.MustCompile(`\+=`), defaultHandler(PEQUALS, "+=")}, {regexp.MustCompile(`-=`), defaultHandler(MEQUALS, "-=")},
			{regexp.MustCompile(`\*=`), defaultHandler(TEQUALS, "*=")}, {regexp.MustCompile(`\/=`), defaultHandler(DEQUALS, "/=")},

			{regexp.MustCompile(`&&`), defaultHandler(AND, "&&")}, {regexp.MustCompile(`\|\|`), defaultHandler(OR, "||")}, {regexp.MustCompile(`##`), defaultHandler(XOR, "##")},

			{regexp.MustCompile(`,`), defaultHandler(COMMA, ",")},
			{regexp.MustCompile(`\.\.\.`), defaultHandler(DOTDOTDOT, "...")}, {regexp.MustCompile(`\.\.`), defaultHandler(DOTDOT, "..")},
			{regexp.MustCompile(`\.`), defaultHandler(DOT, ".")},
			{regexp.MustCompile(`\?`), defaultHandler(QUESTION, "?")}, {regexp.MustCompile(`:`), defaultHandler(COLON, ":")},
		},
		source: source,
	}
}
package lexer

import (
	"regexp"
	"strconv"

	"github.com/uynilo9/bento/pkg/logger"
)

type (
	position struct {
		index uint
		line uint
		column uint
	}
	
	regexHandler func (tokenizer *tokenizer, regex *regexp.Regexp)
	regexPattern struct {
		regex *regexp.Regexp
		handler regexHandler
	}

	tokenizer struct {
		source string
		position position
		patterns []regexPattern
		Tokens []Token
	}
)

func (tokenizer *tokenizer) at() string {
	return string(tokenizer.source[tokenizer.position.index])
}

func (tokenizer *tokenizer) atEOF() bool {
	return tokenizer.position.index >= uint(len(tokenizer.source))
}

func (tokenizer *tokenizer) atLine() string {
	return strconv.Itoa(int(tokenizer.position.line))
}

func (tokenizer *tokenizer) atColumn() string {
	return strconv.Itoa(int(tokenizer.position.column))
}

func (tokenizer * tokenizer) getWhere() (string, string) {
	return tokenizer.atLine(), tokenizer.atColumn()
}

func (tokenizer *tokenizer) remainder() string {
	return tokenizer.source[tokenizer.position.index:]
}

func (tokenizer *tokenizer) push(token Token) {
	tokenizer.Tokens = append(tokenizer.Tokens, token)
}

func (tokenizer *tokenizer) advanceIndex(length uint) {
	tokenizer.position.index += length
}

func (tokenizer *tokenizer) advanceLine(length uint) {
	tokenizer.position.line += length
	tokenizer.position.column = 1
	tokenizer.advanceIndex(length)
}

func (tokenizer *tokenizer) advanceColumn(length uint) {
	tokenizer.position.column += length
	tokenizer.advanceIndex(length)
}

func Tokenize(source string, file string) []Token {
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
			char := tokenizer.at()
			line, column := tokenizer.getWhere()
			logger.Fatalf("Got a legit token `%s` while tokenizing line %s column %s of the file `%s`\n", char, line, column, file)
		}
	}
	tokenizer.push(NewToken(EOF, ""))
	return tokenizer.Tokens
}

func defaultHandler(kind TokenKind, value string) regexHandler {
	return func(tokenizer *tokenizer, regex *regexp.Regexp) {
		tokenizer.advanceColumn(uint(len(value)))
		tokenizer.push(NewToken(kind, value))
	}
}

func newlineHandler(tokenizer *tokenizer, regex *regexp.Regexp) {
	matched := regex.FindString(tokenizer.remainder())
	tokenizer.push(NewToken(NL, strconv.Itoa(len(matched))))
	tokenizer.advanceLine(uint(len(matched)))
}

func numberHandler(tokenizer *tokenizer, regex *regexp.Regexp) {
	matched := regex.FindString(tokenizer.remainder())
	tokenizer.push(NewToken(NUMBER, matched))
	tokenizer.advanceColumn(uint(len(matched)))
}

func newTokenizer(source string) *tokenizer {
	return &tokenizer{
		source: source,
		position: position{
			index: 0,
			line: 1,
			column: 1,
		},
		patterns: []regexPattern{
			{regexp.MustCompile(`\n+`), newlineHandler},
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
		Tokens: make([]Token, 0),
	}
}
package lexer

import (
	"fmt"
	"slices"
)

type (
	TokenKind uint
	Token struct {
		Kind TokenKind
		Value string
	}
)

const (
	EOF TokenKind = iota
	NL

	NUMBER
	STRING
	IDENTIFIER

	LROUND; RROUND		// ( )
	LSQUARE; RSQUARE	// [ ]
	LCURLY; RCURLY		// { }

	NOT 				// !
	EQUALTO; NEQUALTO	// == !=
	LESS; LEQUALTO		// < <=
	GREATER; GEQUALTO	// > >=

	PLUS; PLUSPLUS		// + ++
	MINUS; MINUSMINUS	// - --
	TIMES; TIMESTIMES	// * **
	DIVIDEDBY; MODULO	// / %

	SLASHSLASH 			// //
	
	EQUALS				// =
	PEQUALS; MEQUALS 	// += -=
	TEQUALS; DEQUALS 	// *= /=
	
	AND; OR; XOR		// && || ##
	
	COMMA; DOT			// , .
	DOTDOT; DOTDOTDOT	// .. ...
	QUESTION; COLON		// ? :

	IMP; EXP
	VAR; CONST
	TYPE; ENUM
	STRUCT; INTERF
	CLASS; MODULE
	NEW
	IN; OF
	IF; ELIF; ELSE
	SWICH; CASE
	FOR; WHILE
	BREAK; CONTIN
	RET
)

var tokenMap = map[TokenKind]string{
	EOF: "eof",
	NL: "nl",

	NUMBER: "number",
	STRING: "sring",
	IDENTIFIER: "identifier",

	LROUND: "lround", RROUND: "rround",
	LSQUARE: "lsquare", RSQUARE: "rsquare",
	LCURLY: "lcurly", RCURLY: "rcurly",

	NOT: "not",
	EQUALTO: "equalto", NEQUALTO: "nequalto",
	LESS: "less", LEQUALTO: "lequalto",
	GREATER: "greater", GEQUALTO: "gequalto",

	PLUS: "plus", PLUSPLUS: "plusplus",
	MINUS: "minus", MINUSMINUS: "minusminus",
	TIMES: "times", TIMESTIMES: "timestimes",
	DIVIDEDBY: "dividedby", MODULO: "modulo",

	SLASHSLASH: "slashslash",

	EQUALS: "equals",
	PEQUALS: "pequals", MEQUALS: "mequals",
	TEQUALS: "tequals", DEQUALS: "dequals",

	AND: "and", OR: "or", XOR: "xor",

	COMMA: "comma", DOT: "dot",
	DOTDOT: "dotdot", DOTDOTDOT: "dotdotdot",
	QUESTION: "question", COLON: "colon",

	IMP: "imp", EXP: "exp",
	VAR: "var", CONST: "const",
	TYPE: "type", ENUM: "enum",
	STRUCT: "struct", INTERF: "interf",
	CLASS: "class", MODULE: "module",
	NEW: "new",
	IN: "in", OF: "of",
	IF: "if", ELIF: "elif", ELSE: "else",
	SWICH: "swich", CASE: "case",
	FOR: "for", WHILE: "while",
	BREAK: "break", CONTIN: "contin",
	RET: "ret",
}

func (token Token) isOneOf(kinds ...TokenKind) bool {
	return slices.Contains(kinds, token.Kind)
}

func (token Token) string() string {
	return tokenMap[token.Kind]
}
 
func (token Token) Debug() string {
	var result string
	if token.isOneOf(NL, NUMBER, STRING, IDENTIFIER) {
		result = fmt.Sprintf("{%s: \"%s\"}\n", token.string(), token.Value)
	} else {
		result = fmt.Sprintf("{%s: \"\"}\n", token.string())
	}
	return result
}

func NewToken(kind TokenKind, value string) Token {
	return Token{kind, value}
}
package tokenizer

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
	SPACE; INDENTATION

	IDENTIFIER
	NUMBER
	STRING

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

	IMP; FROM
	VAR; CONST; PROC
	TYPE; ENUM
	STRUCT; INTERF
	CLASS; MODULE
	NEW
	IN; OF
	IF; ELIF; ELSE
	SWITCH; CASE; DEFAULT
	FOR; WHILE
	BREAK; CONTIN
	RET
)

var tokenMap = map[TokenKind]string{
	EOF: "eof",
	NL: "nl",
	SPACE: "space", INDENTATION: "indentation",

	IDENTIFIER: "identifier",
	NUMBER: "number",
	STRING: "sring",

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

	IMP: "imp", FROM: "from",
	VAR: "var", CONST: "const", PROC: "proc",
	TYPE: "type", ENUM: "enum",
	STRUCT: "struct", INTERF: "interf",
	CLASS: "class", MODULE: "module",
	NEW: "new",
	IN: "in", OF: "of",
	IF: "if", ELIF: "elif", ELSE: "else",
	SWITCH: "switch", CASE: "case", DEFAULT: "default",
	FOR: "for", WHILE: "while",
	BREAK: "break", CONTIN: "contin",
	RET: "ret",
}

var keywordMap = map[string]TokenKind{
	"imp": IMP, "from": FROM,
	"var": VAR, "const": CONST, "proc": PROC,
	"type": TYPE, "enum": ENUM,
	"struct": STRUCT, "interf": INTERF,
	"class": CLASS, "module": MODULE,
	"new": NEW,
	"in": IN, "of": OF,
	"if": IF, "elif": ELIF, "else": ELSE,
	"switch": SWITCH, "case": CASE, "default": DEFAULT,
	"for": FOR, "while": WHILE,
	"break": BREAK, "contin": CONTIN,
	"ret": RET,
}

func (token Token) isOneOf(kinds ...TokenKind) bool {
	return slices.Contains(kinds, token.Kind)
}

func (token Token) string() string {
	return tokenMap[token.Kind]
}
 
func (token Token) Debug() string {
	var result string
	if token.isOneOf(NL, SPACE, INDENTATION) {
		result = fmt.Sprintf("%s (%s)\n", token.string(), token.Value)
	} else if token.isOneOf(IDENTIFIER, NUMBER, STRING) {
		result = fmt.Sprintf("%s (\"%s\")\n", token.string(), token.Value)
	} else {
		result = fmt.Sprintf("%s\n", token.string())
	}
	return result
}

func NewToken(kind TokenKind, value string) Token {
	return Token{kind, value}
}
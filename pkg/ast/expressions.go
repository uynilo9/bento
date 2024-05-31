package ast

import (
	"github.com/uynilo9/bento/pkg/tokenizer"
)

type (
	IdentifierExpression struct {
		Value string
	}

	Uint8Expression struct {
		Value uint8
	}
	Uint16Expression struct {
		Value uint16
	}
	Uint32Expression struct {
		Value uint32
	}
	Uint64Expression struct {
		Value uint64
	}
	// UintExpressions interface {
	// 	expression()
	// }

	Int8Expression struct {
		Value int8
	}
	Int16Expression struct {
		Value int16
	}
	Int32Expression struct {
		Value int32
	}
	Int64Expression struct {
		Value int64
	}
	// IntExpressions interface {
	// 	expression()
	// }

	Float32Expression struct {
		Value float32
	}
	Float64Expression struct {
		Value float64
	}
	// FloatExpressions interface {
	// 	expression()
	// }

	StringExpression struct {
		Value string
	}
)

func (_ IdentifierExpression) expression() {}

func (_ Uint8Expression) expression() {}

func (_ Uint16Expression) expression() {}

func (_ Uint32Expression) expression() {}

func (_ Uint64Expression) expression() {}

func (_ Int8Expression) expression() {}

func (_ Int16Expression) expression() {}

func (_ Int32Expression) expression() {}

func (_ Int64Expression) expression() {}

func (_ Float32Expression) expression() {}

func (_ Float64Expression) expression() {}

func (_ StringExpression) expression() {}

type (
	BinaryExpression struct {
		Left Expression
		Right Expression
		Operator tokenizer.Token
	}
)

func (_ BinaryExpression) expression() {}
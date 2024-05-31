package ast

type (
	Statement interface {
		statement()
	}
	Expression interface {
		expression()
	}
)
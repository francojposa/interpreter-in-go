package ast

import "github.com/francojposa/interpreter-in-go/01/monkey/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// Program will be the root Node of every Monkey program AST
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// LetStatement Node represents a Let Statement
// ex1: `let count = 0;`
//    Token.Literal is "let"
//    Identifier is an Identifier representing "count"
//    Value is an Expression representing "0"
type LetStatement struct {
	Token token.Token // the token.LET token,
	Name  *Identifier // the var/fn name to assign the value of the expression to
	Value Expression  // the expression to be evaluated for assignment to an identifier
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// Identifier Node represents an Identifier Expression
// ex1: `let count = 0;`
//    Token.Literal is "count"
//    Value is "count"
type Identifier struct {
	Token token.Token // the token.IDENT token, with the Literal
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

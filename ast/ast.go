package ast

import "github.com/x-color/monkey/token"

// Node is AST node
type Node interface {
	TokenLiteral() string
}

// Statement is statement's interface
type Statement interface {
	Node
	statementNode()
}

// Expression is expression's interface
type Expression interface {
	Node
	expressionNode()
}

// Program is AST root node
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {

}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token
	Value string
}

func expressionNode() {

}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

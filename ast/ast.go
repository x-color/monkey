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

// Program is root node in AST
type Program struct {
	Statements []Statement
}

// TokenLiteral returns current node's literal
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// LetStatement is 'let' statement node in AST
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {

}

// TokenLiteral returns 'let'
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// ReturnStatement is 'return' statement node in AST
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {

}

// TokenLiteral returns 'return'
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

// Identifier is variable node in AST
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {

}

// TokenLiteral returns variable name
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

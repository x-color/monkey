package ast

import (
	"bytes"

	"github.com/x-color/monkey/token"
)

// Node is AST node
type Node interface {
	TokenLiteral() string
	String() string
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

// String returns all statements
func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
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

// String returns 'let' statement
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
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

// String returns 'return' statement
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")
	out.WriteString(" = ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// ExpressionStatement is expression node in AST
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {

}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

// String returns expression statement
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
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

// String returns variable value
func (i *Identifier) String() string {
	return i.Value
}

// IntegerLiteral is integer literal node in AST
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {

}

// TokenLiteral returns integer literal name
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

// String returns integer literal name
func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

// PrefixExpression is prefix expression node in AST
type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {

}

// TokenLiteral returns value after prefix expression
func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

// String returns prefix expression
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}

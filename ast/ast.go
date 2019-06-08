package ast

import (
	"bytes"
	"strings"

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

// TokenLiteral returns current node's token literal
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// String returns all statements string
func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// LetStatement is 'let' statement node in AST
type LetStatement struct {
	Token token.Token // 'let' token
	Name  *Identifier // Variable token
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
	Token       token.Token // 'return' token
	ReturnValue Expression  // Return expression
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
	Token      token.Token // First token of expression statement
	Expression Expression  // Expression node
}

func (es *ExpressionStatement) statementNode() {

}

// TokenLiteral returns first token of expression statement
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

// BlockStatement is block statement({...}) node in AST
type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (bs *BlockStatement) statementNode() {

}

// TODO: Add comment
func (bs *BlockStatement) TokenLiteral() string {
	return bs.Token.Literal
}

// String returns all statements in block statement
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// Identifier is variable node in AST
type Identifier struct {
	Token token.Token // Variable token
	Value string      // Variable name
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
	Token token.Token // Integer literal token
	Value int64       // Integer literal value
}

func (il *IntegerLiteral) expressionNode() {

}

// TokenLiteral returns integer literal value
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

// String returns integer literal value
func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

// PrefixExpression is prefix expression node in AST
type PrefixExpression struct {
	Token    token.Token // Prefix operator token
	Operator string      // Prefix operator
	Right    Expression  // Expression after prefix operator
}

func (pe *PrefixExpression) expressionNode() {

}

// TokenLiteral returns prefix operator
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

// InfixExpression is infix expression node in AST
type InfixExpression struct {
	Token    token.Token // Infix operator token
	Left     Expression  // Expression before infix operator
	Operator string      // Infix operator
	Right    Expression  // Expression after infix operator
}

func (oe *InfixExpression) expressionNode() {

}

// TokenLiteral returns infix operator
func (oe *InfixExpression) TokenLiteral() string {
	return oe.Token.Literal
}

// String returns infix expression
func (oe *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(oe.Left.String())
	out.WriteString(" " + oe.Operator + " ")
	out.WriteString(oe.Right.String())
	out.WriteString(")")
	return out.String()
}

// Boolean is boolean node in AST
type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode() {

}

// TokenLiteral returns boolean literal ('true' or 'false')
func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}

// String returns boolean literal ('true' or 'false')
func (b *Boolean) String() string {
	return b.Token.Literal
}

// IfExpression is 'if-else' expression node in AST
type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode() {

}

// TokenLiteral returns 'if'
func (ie *IfExpression) TokenLiteral() string {
	return ie.Token.Literal
}

// String returns 'if-else' expression
func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString(" ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}

// FunctionLiteral is function node in AST
type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode() {

}

// TokenLiteral returns 'fn'
func (fl *FunctionLiteral) TokenLiteral() string {
	return fl.Token.Literal
}

// String returns function
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ","))
	out.WriteString(")")
	out.WriteString(fl.Body.String())

	return out.String()
}

// CallExpression is calling function node in AST
type CallExpression struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

func (ce *CallExpression) expressionNode() {

}

// TokenLiteral returns function name
func (ce *CallExpression) TokenLiteral() string {
	return ce.Token.Literal
}

// String returns calling function expression
func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}

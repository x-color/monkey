package token

// TokenType is token type (literal, variable or operator ..)
type TokenType string

// Token has token info
type Token struct {
	Type    TokenType
	Literal string
}

const (
	Illegal = "ILLEGAL"
	Eof     = "EOF"

	// variable, literal
	Ident = "IDENT"
	Int   = "INT"

	// operators
	Assign   = "="
	Plus     = "+"
	Minus    = "-"
	Bang     = "!"
	Asterisk = "*"
	Slash    = "/"
	Lt       = "<"
	Gt       = ">"
	Eq       = "=="
	NotEq    = "!="

	// delimeters
	Comma     = ","
	Semicolon = ";"

	// branckets
	LParen = "("
	RParen = ")"
	LBrace = "{"
	RBrace = "}"

	// keywords
	Function = "FUNCTION"
	Let      = "LET"
	True     = "true"
	False    = "false"
	If       = "if"
	Else     = "else"
	Return   = "return"
)

var keywords = map[string]TokenType{
	"fn":     Function,
	"let":    Let,
	"true":   True,
	"false":  False,
	"if":     If,
	"else":   Else,
	"return": Return,
}

// LookupIdent checks if word is keyword
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return Ident
}

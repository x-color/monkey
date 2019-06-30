package token

// TokenType is token type (literal, variable or operator ..)
type TokenType string

// Token has token info
type Token struct {
	Type    TokenType
	Literal string
}

// Token types
const (
	Illegal = "ILLEGAL" // Illegal
	Eof     = "EOF"     // End of file

	Ident  = "IDENT"  // Variable
	Int    = "INT"    // Integer literal
	String = "STRING" // String literal

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
	Colon     = ":"

	// branckets
	LParen   = "("
	RParen   = ")"
	LBrace   = "{"
	RBrace   = "}"
	LBracket = "["
	RBracket = "]"

	// keywords
	Function = "FUNCTION"
	Let      = "LET"
	True     = "TRUE"
	False    = "FALSE"
	If       = "IF"
	While    = "While"
	Else     = "ELSE"
	Return   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     Function,
	"let":    Let,
	"true":   True,
	"false":  False,
	"if":     If,
	"while":  While,
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

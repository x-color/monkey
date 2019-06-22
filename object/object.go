package object

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"strings"

	"github.com/x-color/monkey/ast"
)

// ObjectType is object type (int, bool, null)
type ObjectType string

// BuiltinFunction is builtin function
type BuiltinFunction func(args ...Object) Object

// Object types
const (
	IntegerObj     = "INTEGER"
	StringObj      = "STRING"
	BooleanObj     = "BOOLEAN"
	NullObj        = "NULL"
	ReturnValueObj = "RETURN_VALUE"
	ErrorObj       = "ERROR"
	FunctionObj    = "FUNCTION"
	BuiltinObj     = "BUILTIN"
	ArrayObj       = "ARRAY"
	HashObj        = "HASH"
)

// Object is object interface
type Object interface {
	Type() ObjectType
	Inspect() string
}

// Builtin is builtin function object
type Builtin struct {
	Fn BuiltinFunction
}

// Type returns 'BUILTIN'
func (b *Builtin) Type() ObjectType {
	return BuiltinObj
}

// Inspect returns 'builtin function'
func (b *Builtin) Inspect() string {
	return "builtin function"
}

// Integer is integer object
type Integer struct {
	Value int64
}

// Type returns 'INTEGER'
func (i *Integer) Type() ObjectType {
	return IntegerObj
}

// Inspect returns value of integer object ('5', '99', ...)
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

// String is string object
type String struct {
	Value string
}

// Type returns 'STRING'
func (s *String) Type() ObjectType {
	return StringObj
}

// Inspect returns value of string object ('Hello', 'World', ...)
func (s *String) Inspect() string {
	return s.Value
}

// Boolean is boolean object
type Boolean struct {
	Value bool
}

// Type returns 'BOOLEAN'
func (b *Boolean) Type() ObjectType {
	return BooleanObj
}

// Inspect returns value of boolean object ('true' or 'false')
func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

// Null is null object
type Null struct {
}

// Type returns 'NULL'
func (n *Null) Type() ObjectType {
	return NullObj
}

// Inspect returns null value
func (n *Null) Inspect() string {
	return "null"
}

// ReturnValue is returned value object by 'return' statement
type ReturnValue struct {
	Value Object
}

// Type returns 'RETURN_VALUE'
func (rv *ReturnValue) Type() ObjectType {
	return ReturnValueObj
}

// Inspect returns value of returned value by 'return' statement
func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

// Error is error object
type Error struct {
	Message string
}

// Type returns 'ERROR'
func (e *Error) Type() ObjectType {
	return ErrorObj
}

// Inspect returns error message (e.g. 'ERROR: <Error Message>')
func (e *Error) Inspect() string {
	return "ERROR: " + e.Message
}

// Function is function object
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

// Type returns 'FUNCTION'
func (f *Function) Type() ObjectType {
	return FunctionObj
}

// Inspect returns definition of function
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n")

	return out.String()
}

// Array is array object
type Array struct {
	Elements []Object
}

// Type returns 'ARRAY'
func (ao *Array) Type() ObjectType {
	return ArrayObj
}

// Inspect returns array
func (ao *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ","))
	out.WriteString("]")

	return out.String()
}

// Hashable is hashable object's interface
type Hashable interface {
	HashKey() HashKey
}

// HashKey is key of associative array
type HashKey struct {
	Type  ObjectType
	Value uint64
}

// HashKey returns key for boolean object
func (b *Boolean) HashKey() HashKey {
	var value uint64

	if b.Value {
		value = 1
	} else {
		value = 0
	}

	return HashKey{Type: b.Type(), Value: value}
}

// HashKey returns key for integer object
func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

// HashKey returns key for string object
func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))
	return HashKey{Type: s.Type(), Value: h.Sum64()}
}

// HashPair is pair of key and value of associative array
type HashPair struct {
	Key   Object
	Value Object
}

// Hash is associative array object
type Hash struct {
	Pairs map[HashKey]HashPair
}

// Type returns 'HASH'
func (h *Hash) Type() ObjectType {
	return HashObj
}

// Inspect returns associative array
func (h *Hash) Inspect() string {
	var out bytes.Buffer

	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s",
			pair.Key.Inspect(), pair.Value.Inspect()))
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

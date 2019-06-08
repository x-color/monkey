package object

import "fmt"

// ObjectType is object type (int, bool, null)
type ObjectType string

// Object types
const (
	IntegerObj = "INTEGER"
	BooleanObj = "BOOLEAN"
	NullObj    = "NULL"
	ReturnValueObj = "RETURN_VALUE"
	ErrorObj = "ERROR"
)

// Object is object interface
type Object interface {
	Type() ObjectType
	Inspect() string
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
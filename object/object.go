package object

import "fmt"

// ObjectType is object type (int, bool, null)
type ObjectType string

// Object types
const (
	IntegerObj = "INTEGER"
	BooleanObj = "BOOLEAN"
	NullObj    = "NULL"
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

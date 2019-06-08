package evaluator

import (
	"github.com/x-color/monkey/ast"
	"github.com/x-color/monkey/object"
)

// Constant boolean objects
var (
	True = &object.Boolean{Value: true}
	False = &object.Boolean{Value: false}
)

// Eval evaluates node of AST and retruns evaluated node
func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalStatements(node.Statements)

	case *ast.ExpressionStatement:
		return Eval(node.Expression)

	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)
	}
	return nil
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object
	for _, statement := range stmts {
		result = Eval(statement)
	}
	return result
}

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return True
	}
	return False
}
package evaluator

import (
	"TroInterpreter/ast"
	"TroInterpreter/object"
)

var (
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
	NULL  = &object.Null{}
)

// 求值
func Eval(node ast.Node) object.Object {
	switch node := node.(type) {

	//分析程序
	case *ast.Program:
		return evalStatements(node.Statements)

		//分析表达式
	case *ast.ExpressionStatement:
		return Eval(node.Expression)

		//分析整数
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

		//分析布尔值
	case *ast.Boolean:
		return bool2BoolObject(node.Value)
	}

	return nil
}

// 由于布尔值只有两种，所以直接引用
func bool2BoolObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

// 求值语句
func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range stmts {
		result = Eval(statement)
	}

	return result
}

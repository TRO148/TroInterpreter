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

		//分析前缀表达式
	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)
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

// 求值前缀表达式
func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	default:
		return NULL
	}
}

// 布尔值
func evalBangOperatorExpression(right object.Object) object.Object {
	//操作数在为负或为NULL可以明确表示出来，就用TRUE
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

// 负数
func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	//先检查操作数是不是整数，不是就返回NULL
	if right.Type() != object.INTEGER_OBJ {
		return NULL
	}

	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}

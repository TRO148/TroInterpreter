package evaluator

import (
	"TroInterpreter/ast"
	"TroInterpreter/object"
)

// Eval 求值
func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	//分析程序
	case *ast.Program:
		return evalProgram(node, env)

		//分析表达式
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)

		//分析标识符
	case *ast.Identifier:
		return evalIdentifier(node, env)

		//分析整数
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

		//分析布尔值
	case *ast.Boolean:
		return bool2BoolObject(node.Value)

		//分析前缀表达式
	case *ast.PrefixExpression:
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)

		//分析中缀表达式
	case *ast.InfixExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalInfixExpression(node.Operator, left, right)

		//分析if
	case *ast.IfExpression:
		return evalIfExpression(node, env)

		//分析block
	case *ast.BlockStatement:
		return evalBlockStatements(node, env)

		//分析return
	case *ast.ReturnStatement:
		val := Eval(node.ReturnValue, env)
		if isError(val) {
			return val
		}
		return &object.ReturnValue{Value: val}

		//分析let
	case *ast.LetStatement:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		env.Set(node.Name.Value, val)

	}

	return nil
}

// 求值程序
func evalProgram(program *ast.Program, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range program.Statements {
		result = Eval(statement, env)

		if returnValue, ok := result.(*object.ReturnValue); ok {
			return returnValue.Value
		}
		if errorValue, ok := result.(*object.Error); ok {
			return errorValue
		}
	}

	return result
}

// 求值块语句
func evalBlockStatements(block *ast.BlockStatement, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range block.Statements {
		result = Eval(statement, env)

		if result != nil {
			if result.Type() == object.RETRUN_VALUE_OBJ || result.Type() == object.ERROR_OBJ {
				return result
			}
		}
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
		return newError("错误操作符: %s%s", operator, right.Type())
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
		return newError("错误操作符: -%s", right.Type())
	}

	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}

// 求值中缀表达式
func evalInfixExpression(operator string, left, right object.Object) object.Object {
	//如果都是整数，就用evalIntegerInfixExpression求值
	if left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ {
		return evalIntegerInfixExpression(operator, left, right)
	}

	if operator == "==" {
		//通过都是一个对象，来进行比较，实现true 与 false比较
		return bool2BoolObject(left == right)
	} else if operator == "!=" {
		return bool2BoolObject(left != right)
	}

	if left.Type() != right.Type() {
		//如果两个操作数类型不同，就报错
		return newError("类型不匹配: %s %s %s", left.Type(), operator, right.Type())
	}

	//都不是就报错
	return newError("错误操作符: %s %s %s", left.Type(), operator, right.Type())
}

// 整数中缀运算
func evalIntegerInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "/":
		return &object.Integer{Value: leftVal / rightVal}
	case "<":
		return bool2BoolObject(leftVal < rightVal)
	case ">":
		return bool2BoolObject(leftVal > rightVal)
	case "==":
		return bool2BoolObject(leftVal == rightVal)
	case "!=":
		return bool2BoolObject(leftVal != rightVal)
	default:
		return newError("错误操作符: %s %s %s", left.Type(), operator, right.Type())
	}
}

// 求值if语句
func evalIfExpression(ie *ast.IfExpression, env *object.Environment) object.Object {
	condition := Eval(ie.Condition, env)
	if isError(condition) {
		return condition
	}

	if isTruthy(condition) {
		return Eval(ie.Consequence, env)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative, env)
	} else {
		return NULL
	}
}

// 分析标识符
func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	val, ok := env.Get(node.Value)
	if !ok {
		return newError("无法找到标识符: " + node.Value)
	}

	return val
}

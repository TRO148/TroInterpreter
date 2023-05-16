package evaluator

import "TroInterpreter/object"

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("参数数量错误，期望=1，实际=%d", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("参数类型错误，期望=string，实际=%s", arg.Type())
			}
		},
	},
}

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
	"help": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) >= 2 {
				return newError("参数数量错误，期望<=1，实际=%d", len(args))
			}

			if len(args) == 1 {
				switch arg := args[0].Inspect(); arg {
				case "let":
					return &object.String{
						Value: "let语句用于声明变量，格式为：let 标识符 = 表达式",
					}
				case "return":
					return &object.String{
						Value: "return语句用于返回值，格式为：return 表达式",
					}
				}
			}

			return &object.String{
				Value: "tro使用手册:\n" +
					"本语言分为语句和标识符两大类\n" +
					"语句现在有let与return\n" +
					"表达式有基本类型整型、字符串、函数、布尔值，if与前缀运算符、中缀运算符\n" +
					`help参数可以使用："let","return"，以获取更多信息`,
			}
		},
	},
}

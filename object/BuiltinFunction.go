package object

type BuiltInFunction func(args ...Object) Object

type Builtin struct {
	Fn BuiltInFunction
}

func (b *Builtin) Type() TypeObject {
	return BUILTIN_OBJ
}

func (b *Builtin) Inspect() string {
	return "builtin function"
}

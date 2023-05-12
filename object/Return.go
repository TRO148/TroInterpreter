package object

// 返回值
type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() TypeObject {
	return RETRUN_VALUE_OBJ
}
func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

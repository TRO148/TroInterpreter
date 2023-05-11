package object

import "fmt"

// 整数
type Integer struct {
	Value int64
}

func (i *Integer) Type() TypeObject {
	return INTEGER_OBJ
}
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

package object

import "fmt"

// 布尔值
type Boolean struct {
	Value bool
}

func (b *Boolean) Type() TypeObject {
	return BOOLEAN_OBJ
}
func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

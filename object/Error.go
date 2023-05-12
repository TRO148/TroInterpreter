package object

// 错误
type Error struct {
	Message string
}

func (e *Error) Type() TypeObject {
	return ERROR_OBJ
}
func (e *Error) Inspect() string {
	return "ERROR: " + e.Message
}

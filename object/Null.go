package object

type Null struct{}

func (n *Null) Type() TypeObject {
	return NULL_OBJ
}
func (n *Null) Inspect() string {
	return "null"
}

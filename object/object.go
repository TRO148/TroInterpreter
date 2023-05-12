package object

type TypeObject string

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETRUN_VALUE_OBJ = "RETURN_VALUE"
)

type Object interface {
	Type() TypeObject
	Inspect() string
}

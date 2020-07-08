package object

import "fmt"

type ObjetctType string

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
)

type Object interface {
	Type() ObjetctType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string   { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjetctType { return INTEGER_OBJ }

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjetctType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string   { return fmt.Sprintf("%t", b.Value) }

type Null struct{}

func (n *Null) Type() ObjetctType { return NULL_OBJ }
func (n *Null) Inspect() string   { return "null" }

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjetctType {
	return RETURN_VALUE_OBJ
}

func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

type Error struct {
	Message string
}

func (e *Error) Type() ObjetctType { return ERROR_OBJ }
func (e *Error) Inspect() string   { return "ERROR: " + e.Message }

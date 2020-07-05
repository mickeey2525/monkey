package object

import "fmt"

type ObjetctType string

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ    = "NULL"
)

type Object interface {
	Type() ObjetctType
	Insepect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Insepect() string  { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjetctType { return INTEGER_OBJ }

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjetctType { return BOOLEAN_OBJ }
func (b *Boolean) Insepect() string  { return fmt.Sprintf("%t", b.Value) }

type Null struct{}

func (n *Null) Type() ObjetctType { return NULL_OBJ }
func (n *Null) Insepect() string  { return "null" }

package object

import (
	"bytes"
	"fmt"
	"github.com/mickeey2525/monkey/ast"
	"strings"
)

type ObjetctType string

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
	BUILTIN_OBJ      = "BUILTIN"
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

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

func NewEncloseEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

type Environment struct {
	store map[string]Object
	outer *Environment
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjetctType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

type String struct {
	Value string
}

func (s *String) Type() ObjetctType { return STRING_OBJ }
func (s *String) Inspect() string   { return s.Value }

type BuiltinFunction func(args ...Object) Object

type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() ObjetctType { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string   { return "builtin function" }

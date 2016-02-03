package parg

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Action int

const (
	Store Action = iota
	StoreConst
	StoreTrue
	StoreFalse
	Append
	AppendConst
	Count
	Help
	Version
)

type argument struct {
	Name            string
	ActionType      Action
	ArgNum          string
	ConstVal        interface{}
	DefaultVal      interface{}
	Kind            reflect.Kind
	PossibleChoices []interface{}
	IsPositional    bool
	IsRequired      bool
	HelpText        string
	MetaVarText     []string
	DestName        string
}

func (a *argument) Action(action Action) *argument {
	a.ActionType = action
	return a
}

func (a *argument) Bool() *argument {
	return a.Type(reflect.Bool)
}

func (a *argument) Choices(choices []interface{}) *argument {
	for _, choice := range choices {
		if a.Kind != reflect.Invalid && reflect.ValueOf(choice).Kind() != a.Kind {
			panic(fmt.Errorf("Choice: '%v' must be of type: '%s'", choice, a.Kind.String()))
		}
		choices = append(choices, choice)
	}
	return a
}

func (a *argument) Const(value interface{}) *argument {
	if a.Kind != reflect.Invalid && reflect.ValueOf(value).Kind() != a.Kind {
		panic(fmt.Errorf("Constant value: '%v' must be of type: '%s'", value, a.Kind.String()))
	}
	a.ConstVal = value
	return a
}

func (a *argument) Default(value interface{}) *argument {
	if a.Kind != reflect.Invalid && reflect.ValueOf(value).Kind() != a.Kind {
		panic(fmt.Errorf("Constant value: '%v' must be of type: '%s'", value, a.Kind.String()))
	}
	a.DefaultVal = value
	return a
}

func (a *argument) Dest(name string) *argument {
	a.DestName = name
	return a
}

func (a *argument) Float() *argument {
	return a.Type(reflect.Float64)
}

func (a *argument) Float32() *argument {
	return a.Type(reflect.Float32)
}

func (a *argument) GetUsage() string {
	var argUsage []string

	if a.IsPositional == false {
		argUsage = append(argUsage, "[", "-")
	}

	if len(a.Name) > 1 && a.IsPositional == false {
		argUsage = append(argUsage, "-")
	}

	argUsage = append(argUsage, strings.ToLower(a.Name))

	var nargs []string

	if strings.ContainsAny(a.ArgNum, "?*+rR") == false {
		count := 0
		max, err := strconv.Atoi(a.ArgNum)
		if err != nil {
			panic(err)
		}

		metaLen := len(a.MetaVarText)
		for count < max {
			meta := ""
			if count >= metaLen {
				meta = a.MetaVarText[metaLen-1]
			} else {
				meta = a.MetaVarText[count]
			}
			nargs = append(nargs, strings.ToUpper(meta))
			count++
		}
		if len(nargs) > 0 {
			argUsage = append(argUsage, " ", join(" ", nargs...))
		}
	} else {
		switch a.ArgNum {
		case "?":
			argUsage = append(
				argUsage,
				" [",
				strings.ToUpper(a.MetaVarText[0]),
				"]",
			)
		case "+":
			fallthrough
		case "*":
			first := a.MetaVarText[0]
			second := first

			if len(a.MetaVarText) > 1 {
				second = a.MetaVarText[1]
			}

			before := ""
			after := ""
			if a.ArgNum == "*" {
				before = "["
				after = "]"
			}

			argUsage = append(
				argUsage,
				" ",
				before,
				strings.ToUpper(first),
				" [",
				strings.ToUpper(second),
				" ...]",
				after,
			)
		case "r":
			fallthrough
		case "R":
			argUsage = append(
				argUsage,
				" ",
				"...",
			)
		}
	}

	if a.IsPositional == false {
		argUsage = append(argUsage, "]")
	}

	return join("", argUsage...)
}

func (a *argument) Help(text string) *argument {
	a.HelpText = text
	return a
}

func (a *argument) Int64() *argument {
	return a.Type(reflect.Int64)
}

func (a *argument) Int32() *argument {
	return a.Type(reflect.Int32)
}

func (a *argument) MetaVar(meta string, metaSlice ...string) *argument {
	s := []string{meta}
	for _, text := range metaSlice {
		s = append(s, text)
	}

	a.MetaVarText = s
	return a
}

func (a *argument) Nargs(nargs string) *argument {
	allowed_chars := []string{"?", "*", "+", "r", "R"}
	for _, char := range allowed_chars {
		if nargs == char {
			a.ArgNum = char
			return a
		}
	}

	_, err := strconv.Atoi(nargs)
	if err != nil {
		panic(fmt.Errorf("Invalid nargs amount/character: '%s'", nargs))
	}

	a.ArgNum = nargs
	return a
}

func (a *argument) NotPositional() *argument {
	a.IsPositional = false
	return a
}

func (a *argument) NotRequired() *argument {
	a.IsRequired = false
	return a
}

func (a *argument) Positional() *argument {
	a.IsPositional = true
	return a
}

func (a *argument) Required() *argument {
	a.IsRequired = true
	return a
}

func (a *argument) String() *argument {
	return a.Type(reflect.String)
}

func (a *argument) Uint64() *argument {
	return a.Type(reflect.Uint64)
}

func (a *argument) Uint32() *argument {
	return a.Type(reflect.Uint32)
}

func (a *argument) Type(kind reflect.Kind) *argument {
	a.Kind = kind
	return a
}

func Argument(name string, help string) *argument {
	a := argument{ActionType: StoreTrue, DestName: name, MetaVarText: []string{name}, Name: name, IsPositional: false, ArgNum: "0", HelpText: help}
	return &a
}

func Flag(name, help string) *argument {
	a := argument{
		ActionType: StoreTrue,
		ArgNum:     "0",
		DefaultVal: false,
		DestName:   name,
		HelpText:   help,
		Name:       name,
	}
	return &a
}

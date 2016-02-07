package parg

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func NewFlag(name, help string) *Flag {
	f := Flag{
		ArgNum:        "0",
		ConstVal:      nil,
		DefaultVal:    false,
		DesiredAction: StoreTrue,
		DestName:      name,
		HelpText:      help,
		IsRequired:    false,
		MetaVarText:   []string{name},
		PublicName:    name,
		RequiredKind:  reflect.Bool,
	}
	return &f
}

type Flag struct {
	ArgNum          string
	ConstVal        interface{}
	DefaultVal      interface{}
	DesiredAction   Action
	DestName        string
	HelpText        string
	IsRequired      bool
	MetaVarText     []string
	PossibleChoices []interface{}
	PublicName      string
	RequiredKind    reflect.Kind
}

func (f *Flag) Action(action Action) *Flag {
	f.DesiredAction = action
	return f
}

func (f *Flag) Bool() *Flag {
	return f.Kind(reflect.Bool)
}

func (f *Flag) Choices(choices []interface{}) *Flag {
	for _, choice := range choices {
		if f.RequiredKind != reflect.Invalid && reflect.ValueOf(choice).Kind() != f.RequiredKind {
			panic(fmt.Errorf("Choice: '%v' must be of type: '%s'", choice, f.RequiredKind.String()))
		}
		choices = append(choices, choice)
	}
	return f
}

func (f *Flag) Const(value interface{}) *Flag {
	if f.RequiredKind != reflect.Invalid && reflect.ValueOf(value).Kind() != f.RequiredKind {
		panic(fmt.Errorf("Constant value: '%v' must be of type: '%s'", value, f.RequiredKind.String()))
	}
	f.ConstVal = value
	return f
}

func (f *Flag) Default(value interface{}) *Flag {
	if f.RequiredKind != reflect.Invalid && reflect.ValueOf(value).Kind() != f.RequiredKind {
		panic(fmt.Errorf("Constant value: '%v' must be of type: '%s'", value, f.RequiredKind.String()))
	}
	f.DefaultVal = value
	return f
}

func (f *Flag) Dest(name string) *Flag {
	f.DestName = name
	return f
}

func (f *Flag) Float() *Flag {
	return f.Kind(reflect.Float64)
}

func (f *Flag) Float32() *Flag {
	return f.Kind(reflect.Float32)
}

func (f *Flag) DisplayName() string {
	var prefix string

	if len(f.PublicName) > 1 {
		prefix = "-"
	} else if len(f.PublicName) == 1 {
		prefix = "--"
	}
	return join("", prefix, strings.ToLower(f.PublicName))
}

func (f *Flag) GetUsage() string {
	var usage []string

	usage = append(usage, "[", f.DisplayName())

	var nargs []string

	if strings.ContainsAny(f.ArgNum, "?*+") == false {
		count := 0
		max, err := strconv.Atoi(f.ArgNum)
		if err != nil {
			panic(err)
		}

		metaLen := len(f.MetaVarText)
		if metaLen == 0 {
			f.MetaVarText = []string{f.PublicName}
			metaLen = 1
		}

		for count < max {
			meta := ""
			if count >= metaLen {
				meta = f.MetaVarText[metaLen-1]
			} else {
				meta = f.MetaVarText[count]
			}
			nargs = append(nargs, strings.ToUpper(meta))
			count++
		}
		if len(nargs) > 0 {
			usage = append(usage, " ", join(" ", nargs...))
		}
	} else {
		switch f.ArgNum {
		case "?":
			usage = append(
				usage,
				" [",
				strings.ToUpper(f.MetaVarText[0]),
				"]",
			)
		case "+":
			fallthrough
		case "*":
			first := f.PublicName
			if len(f.MetaVarText) > 0 {
				first = f.MetaVarText[0]
			}
			second := first

			if len(f.MetaVarText) > 1 {
				second = f.MetaVarText[1]
			}

			before := ""
			after := ""
			if f.ArgNum == "*" {
				before = "["
				after = "]"
			}

			usage = append(
				usage,
				" ",
				before,
				strings.ToUpper(first),
				" [",
				strings.ToUpper(second),
				" ...]",
				after,
			)
		}
	}

	usage = append(usage, "]")

	return join("", usage...)
}

func (f *Flag) Help(text string) *Flag {
	f.HelpText = text
	return f
}

func (f *Flag) Int64() *Flag {
	return f.Kind(reflect.Int64)
}

func (f *Flag) Int32() *Flag {
	return f.Kind(reflect.Int32)
}

func (f *Flag) MetaVar(meta string, metaSlice ...string) *Flag {
	s := []string{meta}
	for _, text := range metaSlice {
		s = append(s, text)
	}

	f.MetaVarText = s
	return f
}

func (f *Flag) Nargs(nargs string) *Flag {
	// TODO: Allow "r"/"R" for remainder args
	allowed_chars := []string{"?", "*", "+"}
	for _, char := range allowed_chars {
		if nargs == char {
			f.ArgNum = char
			return f
		}
	}

	_, err := strconv.Atoi(nargs)
	if err != nil {
		panic(fmt.Errorf("Invalid nargs amount/character: '%s'", nargs))
	}

	f.ArgNum = nargs
	return f
}

func (f *Flag) NotRequired() *Flag {
	f.IsRequired = false
	return f
}

func (f *Flag) Required() *Flag {
	f.IsRequired = true
	return f
}

func (f *Flag) String() *Flag {
	return f.Kind(reflect.String)
}

func (f *Flag) Uint64() *Flag {
	return f.Kind(reflect.Uint64)
}

func (f *Flag) Uint32() *Flag {
	return f.Kind(reflect.Uint32)
}

func (f *Flag) Kind(kind reflect.Kind) *Flag {
	f.RequiredKind = kind
	return f
}

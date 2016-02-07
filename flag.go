package parg

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// NewFlag returns a pointer to a new Flag instance, setting the flag's destination
// name, public name and metavar text to the provided name, and the help text to the
// provided help string.
func NewFlag(name, help string) *Flag {
	f := Flag{
		ArgNum:        "0",
		ConstVal:      nil,
		DefaultVal:    nil,
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

// Flag contains the necessary attributes for representing a parsable flag.
type Flag struct {
	ArgNum          string
	ConstVal        interface{}
	DefaultVal      interface{}
	DesiredAction   Action
	DestName        string
	HelpText        string
	IsRequired      bool
	MetaVarText     []string
	PossibleChoices []interface{} // Currently unused. TODO: implement.
	PublicName      string
	RequiredKind    reflect.Kind
}

// Action sets the flag's action to the provided action function.
func (f *Flag) Action(action Action) *Flag {
	f.DesiredAction = action
	return f
}

// Bool sets the flag's required type as a boolean.
func (f *Flag) Bool() *Flag {
	return f.Kind(reflect.Bool)
}

// Choices appends the provided slice as acceptable arguments for the flag.
func (f *Flag) Choices(choices []interface{}) *Flag {
	for _, choice := range choices {
		if f.RequiredKind != reflect.Invalid && reflect.ValueOf(choice).Kind() != f.RequiredKind {
			panic(fmt.Errorf("Choice: '%v' must be of type: '%s'", choice, f.RequiredKind.String()))
		}
		choices = append(choices, choice)
	}
	return f
}

// Const sets the flag's constant value to the provided interface. A flag's constant value
// is only used for certain actions. By default, the constant value is `nil`.
func (f *Flag) Const(value interface{}) *Flag {
	f.ConstVal = value
	return f
}

// Default sets the flag's default value. A flag's default value is only used for
// certain actions. By default, the default value is `nil`.
func (f *Flag) Default(value interface{}) *Flag {
	f.DefaultVal = value
	return f
}

// Dest sets a flag's destination name. This is used as the key for storing the flag's
// values within the parser.
func (f *Flag) Dest(name string) *Flag {
	f.DestName = name
	return f
}

// Float sets the flag's required type as a float.
func (f *Flag) Float() *Flag {
	return f.Kind(reflect.Float64)
}

// Float32 sets the flag's required type as a float32.
func (f *Flag) Float32() *Flag {
	return f.Kind(reflect.Float32)
}

// DisplayName returns the flag's public name, prefixed with the appropriate number
// of hyphen-minus characters.
func (f *Flag) DisplayName() string {
	var prefix string

	if len(f.PublicName) > 1 {
		prefix = "-"
	} else if len(f.PublicName) == 1 {
		prefix = "--"
	}
	return join("", prefix, strings.ToLower(f.PublicName))
}

// GetUsage returns the usage text for the flag. This includes proper formatting
// of the flag's display name & parameters. For parameters: by default, parameters
// will be the flag's public name. This can be overridden by modifying the MetaVars
// slice for the flag.
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

// Help sets the flag's help/usage text.
func (f *Flag) Help(text string) *Flag {
	f.HelpText = text
	return f
}

// Int64 sets the flag's required type as a int64.
func (f *Flag) Int64() *Flag {
	return f.Kind(reflect.Int64)
}

// Int32 sets the flag's required type as a int32.
func (f *Flag) Int32() *Flag {
	return f.Kind(reflect.Int32)
}

// MetaVar sets the flag's metavar text to the provided string. Additional
// metavar strings can be provided, and will be used for flags with more than
// expected argument.
func (f *Flag) MetaVar(meta string, metaSlice ...string) *Flag {
	s := []string{meta}
	for _, text := range metaSlice {
		s = append(s, text)
	}

	f.MetaVarText = s
	return f
}

// Nargs sets the flag's number of expected arguments. Integers represent
// the absolute number of arguments to be expected. The `?` character represents
// an expection of zero or one arguments. The `*` character represents an expectation
// of any number or arguments. The `+` character represents an expectation of one
// or more arguments.
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

// NotRequired prevents the flag from being required to be present when parsing
// arguments.
func (f *Flag) NotRequired() *Flag {
	f.IsRequired = false
	return f
}

// Required enables the flag to required to be present when parsing arguments.
func (f *Flag) Required() *Flag {
	f.IsRequired = true
	return f
}

// String sets the flag's required type as a string.
func (f *Flag) String() *Flag {
	return f.Kind(reflect.String)
}

// Uint64 sets the flag's required type as a uint64.
func (f *Flag) Uint64() *Flag {
	return f.Kind(reflect.Uint64)
}

// Uint32 sets the flag's required type as a uint32.
func (f *Flag) Uint32() *Flag {
	return f.Kind(reflect.Uint32)
}

// Kind sets the flag's required kind to the provided reflect.Kind constant.
func (f *Flag) Kind(kind reflect.Kind) *Flag {
	f.RequiredKind = kind
	return f
}

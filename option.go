package argparse

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func NewFlag(names, dest, help string) *Option {
	opt := NewOption(names, dest, help)
	opt.Nargs("0").Action(StoreTrue).Default("false")

	return opt
}

func NewArg(names, dest, help string) *Option {
	return NewOption(names, dest, help).Nargs("1").Action(Store)
}

// ValidateChoice returns an error if the provided interface value
// does not exists as valid choice for the provided flag.
func ValidateChoice(f Option, arg string) error {
	if len(f.ValidChoices) == 0 {
		return nil
	}

	for _, c := range f.ValidChoices {
		if arg == c {
			return nil
		}
	}

	return InvalidChoiceErr{f, arg}
}

// ValidateType attempt to type-convert the string argument to the flag's desired
// type. It will return an error if the provided interface value does not
// satisfy the Option's expected Reflect.Kind type.
func ValidateType(f Option, arg string) error {
	switch f.ExpectedType {
	case reflect.Invalid, reflect.String:
		return nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if _, err := strconv.Atoi(arg); err == nil {
			return nil
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if _, err := strconv.ParseUint(arg, 10, 0); err == nil {
			return nil
		}
	case reflect.Float32:
		if _, err := strconv.ParseFloat(arg, 32); err == nil {
			return nil
		}
	case reflect.Float64:
		if _, err := strconv.ParseFloat(arg, 64); err == nil {
			return nil
		}
	case reflect.Bool:
		if _, err := strconv.ParseBool(arg); err == nil {
			return nil
		}
	}
	return InvalidTypeErr{f, arg}
}

// NewOption instantiates a new Option pointer, initializing it as a boolean
// flag. Multiple names should be delimited by a space; names should not
// contain the prefix character.
func NewOption(names, dest, help string) *Option {
	f := Option{
		ArgNum:        "0",
		ConstVal:      "",
		DefaultVal:    "",
		DesiredAction: StoreTrue,
		DestName:      dest,
		HelpText:      help,
		MetaVarText:   []string{},
		PublicNames:   strings.Split(names, " "),
		ValidChoices:  []string{},
	}
	return &f
}

// Option contains the necessary attributes for representing a parsable option.
type Option struct {
	ArgNum        string
	ConstVal      string
	DefaultVal    string
	DesiredAction Action
	DestName      string
	ExpectedType  reflect.Kind
	HelpText      string
	IsRequired    bool
	IsPositional  bool
	MetaVarText   []string
	PublicNames   []string
	ValidChoices  []string
}

// Action sets the option's action to the provided action function.
func (f *Option) Action(action Action) *Option {
	f.DesiredAction = action
	return f
}

// Choices appends the provided slice as acceptable arguments for the option.
func (f *Option) Choices(choices ...string) *Option {
	f.ValidChoices = []string{}
	for _, choice := range choices {
		f.ValidChoices = append(f.ValidChoices, choice)
	}
	return f
}

// Const sets the option's constant value to the provided interface. A option's constant value
// is only used for certain actions. By default, the constant value is `nil`.
func (f *Option) Const(value string) *Option {
	f.ConstVal = value
	return f
}

// Default sets the option's default value. A option's default value is only used for
// certain actions. By default, the default value is `nil`.
func (f *Option) Default(value string) *Option {
	f.DefaultVal = value
	return f
}

// Dest sets a option's destination name. This is used as the key for storing the option's
// values within the parser.
func (f *Option) Dest(name string) *Option {
	f.DestName = name
	return f
}

// DisplayName returns the option's public name, prefixed with the appropriate number
// of hyphen-minus characters.
func (f *Option) DisplayName() string {
	getDisplayName := func(name string) string {
		var prefix string

		if f.IsPositional == false {
			if len(name) == 1 {
				prefix = "-"
			} else if len(name) > 1 {
				prefix = "--"
			}
		}

		return join("", prefix, strings.ToLower(name))
	}

	var names []string
	for _, name := range f.PublicNames {
		names = append(names, getDisplayName(name))
	}

	return strings.Join(names, ", ")
}

// GetChoices returns a string-representation of the valid chocies for the
// current Option.
func (f *Option) GetChoices() string {
	if len(f.ValidChoices) == 0 {
		return ""
	}
	var choices []string
	for _, i := range f.ValidChoices {
		choices = append(choices, fmt.Sprintf("%v", i))
	}
	return join("", "{", strings.Join(choices, ","), "}")
}

// GetUsage returns the usage text for the option. This includes proper formatting
// of the option's display name & parameters. For parameters: by default, parameters
// will be the option's public name. This can be overridden by modifying the MetaVars
// slice for the option.
func (f *Option) GetUsage() string {
	var usage []string

	isRequired := f.IsRequired
	if isRequired == false {
		usage = append(usage, "[")
	}

	if len(f.PublicNames) == 1 {
		usage = append(usage, f.DisplayName())
	} else {
		pNames := f.PublicNames
		f.PublicNames = []string{f.PublicNames[0]}
		usage = append(usage, f.DisplayName())
		f.PublicNames = pNames
	}

	var nargs []string
	choices := f.GetChoices()
	if len(choices) == 0 && len(f.MetaVarText) == 0 {
		f.MetaVarText = []string{f.DestName}
	} else if len(f.MetaVarText) == 0 {
		f.MetaVarText = []string{choices}
	}

	if strings.ContainsAny(f.ArgNum, "?*+rR") == false {
		count := 0
		max, err := strconv.Atoi(f.ArgNum)
		if err != nil {
			panic(err)
		}

		metaLen := len(f.MetaVarText)

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
		case "r":
			fallthrough
		case "R":
			usage = append(
				usage,
				" ",
				" ...",
			)
		case "+":
			fallthrough
		case "*":
			first := f.DestName
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

	if isRequired == false {
		usage = append(usage, "]")
	}

	return join("", usage...)
}

// Help sets the option's help/usage text.
func (f *Option) Help(text string) *Option {
	f.HelpText = text
	return f
}

func (f *Option) IsPublicName(name string) bool {
	for _, opName := range f.PublicNames {
		if name == opName {
			return true
		}
	}
	return false
}

// MetaVar sets the option's metavar text to the provided string. Additional
// metavar strings can be provided, and will be used for options with more than
// expected argument.
func (f *Option) MetaVar(meta string, metaSlice ...string) *Option {
	s := []string{meta}
	for _, text := range metaSlice {
		s = append(s, text)
	}

	f.MetaVarText = s
	return f
}

// Nargs sets the option's number of expected arguments. Integers represent
// the absolute number of arguments to be expected. The `?` character represents
// an expection of zero or one arguments. The `*` character represents an expectation
// of any number or arguments. The `+` character represents an expectation of one
// or more arguments.
func (f *Option) Nargs(nargs string) *Option {
	// TODO: Allow "r"/"R" for remainder args
	allowedChars := []string{"?", "*", "+", "r", "R"}
	for _, char := range allowedChars {
		if nargs == char {
			f.ArgNum = char
			return f
		}
	}

	_, err := strconv.Atoi(nargs)
	if err != nil {
		panic(fmt.Errorf("Invalid nargs: '%s' Must be an int or one of: '?*+'", nargs))
	}

	f.ArgNum = nargs
	return f
}

// NotRequired prevents the option from being required to be present when parsing
// arguments.
func (f *Option) NotRequired() *Option {
	f.IsRequired = false
	return f
}

// NotPositional disables a option from being positionally interpretted.
func (f *Option) NotPositional() *Option {
	f.IsPositional = false
	return f
}

// Positional enables a option to be positionally interpretted.
func (f *Option) Positional() *Option {
	f.IsPositional = true
	return f
}

// Required enables the option to required to be present when parsing arguments.
func (f *Option) Required() *Option {
	f.IsRequired = true
	return f
}

// String outputs a string-serialized version of the Option.
func (f *Option) String() string {
	return join(" ", f.GetUsage(), f.HelpText)
}

// Type sets the expected reflect.Kind type an option will accept.
func (f *Option) Type(kind reflect.Kind) *Option {
	invalidKinds := []reflect.Kind{
		reflect.Uintptr,
		reflect.Complex64,
		reflect.Complex128,
		reflect.Array,
		reflect.Chan,
		reflect.Func,
		reflect.Interface,
		reflect.Map,
		reflect.Ptr,
		reflect.Slice,
		reflect.Struct,
		reflect.UnsafePointer,
	}

	for _, bad := range invalidKinds {
		if kind == bad {
			panic(fmt.Sprintf("Cannot use kind: '%s' as a valid type", kind.String()))
		}
	}

	f.ExpectedType = kind
	return f
}

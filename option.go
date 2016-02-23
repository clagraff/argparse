package parg

import (
	"fmt"
	"strconv"
	"strings"
)

// NewOption returns a pointer to a new Option instance, setting the option's destination
// name, public name and metavar text to the provided name, and the help text to the
// provided help string.
func NewOption(name, help string) *Option {
	f := Option{
		ArgNum:        "0",
		ConstVal:      nil,
		DefaultVal:    nil,
		DesiredAction: StoreTrue,
		DestName:      name,
		HelpText:      help,
		MetaVarText:   []string{name},
		PublicName:    name,
	}
	return &f
}

// Option contains the necessary attributes for representing a parsable option.
type Option struct {
	ArgNum          string
	ConstVal        interface{}
	DefaultVal      interface{}
	DesiredAction   Action
	DestName        string
	HelpText        string
	IsRequired      bool
	IsPositional    bool
	MetaVarText     []string
	PossibleChoices []interface{} // Currently unused. TODO: implement.
	PublicName      string
}

// Action sets the option's action to the provided action function.
func (f *Option) Action(action Action) *Option {
	f.DesiredAction = action
	return f
}

// Choices appends the provided slice as acceptable arguments for the option.
func (f *Option) Choices(choices []interface{}) *Option {
	f.PossibleChoices = []interface{}{}
	for _, choice := range choices {
		f.PossibleChoices = append(f.PossibleChoices, choice)
	}
	return f
}

// Const sets the option's constant value to the provided interface. A option's constant value
// is only used for certain actions. By default, the constant value is `nil`.
func (f *Option) Const(value interface{}) *Option {
	f.ConstVal = value
	return f
}

// Default sets the option's default value. A option's default value is only used for
// certain actions. By default, the default value is `nil`.
func (f *Option) Default(value interface{}) *Option {
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
	var prefix string

	if f.IsPositional == false {
		if len(f.PublicName) == 1 {
			prefix = "-"
		} else if len(f.PublicName) > 1 {
			prefix = "--"
		}
	}

	return join("", prefix, strings.ToLower(f.PublicName))
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

	usage = append(usage, f.DisplayName())

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
	allowedChars := []string{"?", "*", "+"}
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

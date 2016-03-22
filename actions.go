package argparse

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Action is type to represent a callable function which will operate on a parser,
// a option, and an array of argument strings.
type Action func(*Parser, *Option, ...string) ([]string, error)

// Store will attempt to store the appropriate number of arguments for the option,
// (if any), into the parser. Remaining arguments & any errors are returned.
func Store(p *Parser, f *Option, args ...string) ([]string, error) {
	// If we are not expecting any arguments, panic!
	if f.ArgNum == "0" {
		panic(fmt.Sprintf("option '%s' must expect at least one argument", f.DisplayName()))
	} else if f.ArgNum == "?" {
		if len(args) > 0 {
			if err := ValidateChoice(*f, args[0]); err != nil {
				return args, err
			} else if err := ValidateType(*f, args[0]); err != nil {
				return args, err
			}
			p.Namespace.Set(f.DestName, args[0])
			return args[1:], nil
		}
	} else if strings.ContainsAny(f.ArgNum, "*+rR") == true {
		if f.ArgNum == "+" && len(args) == 0 {
			return args, TooFewArgsErr{*f}
		}
		var values []string
		for len(args) > 0 {
			if err := ValidateChoice(*f, args[0]); err != nil {
				return args, err
			} else if err := ValidateType(*f, args[0]); err != nil {
				return args, err
			}
			values = append(values, args[0])
			args = args[1:]
		}

		p.Namespace.Set(f.DestName, values)
		return args, nil
	} else if regexp.MustCompile(`^[1-9]+$`).MatchString(f.ArgNum) == true {
		num, _ := strconv.Atoi(f.ArgNum)
		if len(args) < num {
			return args, TooFewArgsErr{*f}
		}

		if num > 1 {
			var values []string
			for _, v := range args[0:num] {
				if err := ValidateChoice(*f, v); err != nil {
					return args, err
				} else if err := ValidateType(*f, v); err != nil {
					return args, err
				}
				values = append(values, v)
			}
			p.Namespace.Set(f.DestName, values)
			if num > len(args) {
				args = []string{}
			} else {
				args = args[num:]
			}
		} else {
			if err := ValidateChoice(*f, args[0]); err != nil {
				return args, err
			} else if err := ValidateType(*f, args[0]); err != nil {
				return args, err
			}
			p.Namespace.Set(f.DestName, args[0])
			if len(args) > 1 {
				args = args[1:]
			} else {
				args = []string{}
			}
		}
	}

	return args, nil
}

// StoreConst stores the option's constant value into the parser. Provided
// arguments remain unmodified.
func StoreConst(p *Parser, f *Option, args ...string) ([]string, error) {
	if f.ArgNum != "0" {
		panic(fmt.Sprintf("option '%s' cannot expect any arguments.", f.DisplayName()))
	}
	p.Namespace.Set(f.DestName, f.ConstVal)

	return args, nil
}

// StoreFalse stores a boolean `false` into the parser. Provided arguments remain
// unmodified.
func StoreFalse(p *Parser, f *Option, args ...string) ([]string, error) {
	if f.ArgNum != "0" {
		panic(fmt.Sprintf("option '%s' cannot expect any arguments.", f.DisplayName()))
	}
	p.Namespace.Set(f.DestName, "false")

	return args, nil
}

// StoreTrue stores a boolean `true` into the parser. Provided arguments remain unmodified.
func StoreTrue(p *Parser, f *Option, args ...string) ([]string, error) {
	if f.ArgNum != "0" {
		panic(fmt.Sprintf("option '%s' cannot expect any arguments.", f.DisplayName()))
	}
	p.Namespace.Set(f.DestName, "true")

	return args, nil
}

// Append retrives the appropriate number of argumnents for the current option, (if any),
// and appends them individually into the parser. Remaining arguments and errors are returned.
func Append(p *Parser, f *Option, args ...string) ([]string, error) {
	appendValue := func(p *Parser, f *Option, value interface{}) error {
		if p.Namespace.KeyExists(f.DestName) == false {
			p.Namespace.Set(f.DestName, make([]string, 0))
		}
		slice, err := p.Namespace.Try(f.DestName)
		if err != nil {
			return err
		}
		slice = append(slice.([]string), value.(string))
		p.Namespace.Set(f.DestName, slice)
		return nil
	}

	if regexp.MustCompile(`^[1-9]+$`).MatchString(f.ArgNum) == true {
		num, _ := strconv.Atoi(f.ArgNum)
		if len(args) < num {
			return args, TooFewArgsErr{*f}
		}

		count := 0
		for count < num {
			if err := ValidateChoice(*f, args[0]); err != nil {
				return args, err
			} else if err := ValidateType(*f, args[0]); err != nil {
				return args, err
			}
			appendValue(p, f, args[0])
			args = args[1:]
			count++
		}
		return args, nil
	} else if f.ArgNum == "0" {
		appendValue(p, f, f.DefaultVal)
		return args, nil
	} else if f.ArgNum == "?" {
		if len(args) > 0 {
			if err := ValidateChoice(*f, args[0]); err != nil {
				return args, err
			} else if err := ValidateType(*f, args[0]); err != nil {
				return args, err
			}
			appendValue(p, f, args[0])
			args = args[1:]
		} else {
			appendValue(p, f, f.DefaultVal)
		}
	} else if strings.ContainsAny(f.ArgNum, "*+rR") == true {
		if f.ArgNum == "+" && len(args) == 0 {
			return args, MissingOneOrMoreArgsErr{*f}
		}

		for len(args) > 0 {
			if err := ValidateChoice(*f, args[0]); err != nil {
				return args, err
			} else if err := ValidateType(*f, args[0]); err != nil {
				return args, err
			}
			appendValue(p, f, args[0])
			args = args[1:]
		}

		return args, nil
	}

	return args, nil
}

// AppendConst appends the option's constant value into the parser. Provided arguments
// remain unmodified.
func AppendConst(p *Parser, f *Option, args ...string) ([]string, error) {
	if f.ArgNum != "0" {
		panic(fmt.Sprintf("option '%s' cannot expect any arguments.", f.DisplayName()))
	}

	if p.Namespace.KeyExists(f.DestName) == false {
		p.Namespace.Set(f.DestName, make([]string, 0))
	}
	slice, err := p.Namespace.Try(f.DestName)
	if err != nil {
		return nil, err
	}
	slice = append(slice.([]string), f.ConstVal)
	p.Namespace.Set(f.DestName, slice)
	return args, nil
}

// ShowHelp calls the parser's ShowHelp function to output parser usage information
// and help information for each option to stdout. Provided arguments remain unchanged.
func ShowHelp(p *Parser, f *Option, args ...string) ([]string, error) {
	p.ShowHelp()
	return args, ShowHelpErr{}
}

// ShowVersion calls the parser's ShowVersion function to output parser/program
// version information. Provided arguments remain unchanged.
func ShowVersion(p *Parser, f *Option, args ...string) ([]string, error) {
	p.ShowVersion()
	return args, nil
}

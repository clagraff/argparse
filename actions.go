package parg

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// Action is type to represent a callable function which will operate on a parser,
// a flag, and an array of argument strings.
type Action func(*Parser, *Flag, ...string) ([]string, error)

// Store will attempt to store the appropriate number of arguments for the flag,
// (if any), into the parser. Remaining arguments & any errors are returned.
func Store(p *Parser, f *Flag, args ...string) ([]string, error) {
	// If we are not expecting any arguments, panic!
	if f.ArgNum == "0" {
		panic(fmt.Sprintf("flag '%s' must expect at least one argument", f.DisplayName()))
	} else if f.ArgNum == "?" {
		if len(args) > 0 {
			p.Values[f.DestName] = args[0]
			return args[1:], nil
		}
	} else if strings.ContainsAny(f.ArgNum, "*+") == true {
		if f.ArgNum == "+" && len(args) == 0 {
			return args, fmt.Errorf("flag '%s' expects at least one argument", f.DisplayName())
		}

		var values []interface{}
		for len(args) > 0 {
			values = append(values, args[0])
			args = args[1:]
		}

		p.Values[f.DestName] = values
		return args, nil
	} else if regexp.MustCompile(`^[1-9]+$`).MatchString(f.ArgNum) == true {
		num, _ := strconv.Atoi(f.ArgNum)
		if len(args) < num {
			return args, fmt.Errorf("flag '%s' is expecting %d argument(s) but was provided %d", f.DisplayName(), num, len(args))
		}

		if num > 1 {
			var values []string
			values = append(values, args[0:num]...)
			p.Values[f.DestName] = values
		} else {
			p.Values[f.DestName] = args[0]
		}
	}

	return args, nil
}

// StoreConst stores the flag's constant value into the parser. Provided
// arguments remain unmodified.
func StoreConst(p *Parser, f *Flag, args ...string) ([]string, error) {
	if f.ArgNum != "0" {
		panic(fmt.Sprintf("flag '%s' cannot expect any arguments.", f.DisplayName()))
	}
	p.Values[f.DestName] = f.ConstVal

	return args, nil
}

// StoreFalse stores a boolean `false` into the parser. Provided arguments remain
// unmodified.
func StoreFalse(p *Parser, f *Flag, args ...string) ([]string, error) {
	if f.ArgNum != "0" {
		panic(fmt.Sprintf("flag '%s' cannot expect any arguments.", f.DisplayName()))
	}
	p.Values[f.DestName] = false

	return args, nil
}

// StoreTrue stores a boolean `true` into the parser. Provided arguments remain unmodified.
func StoreTrue(p *Parser, f *Flag, args ...string) ([]string, error) {
	if f.ArgNum != "0" {
		panic(fmt.Sprintf("flag '%s' cannot expect any arguments.", f.DisplayName()))
	}
	p.Values[f.DestName] = true

	return args, nil
}

// Append retrives the appropriate number of argumnents for the current flag, (if any),
// and appends them individually into the parser. Remaining arguments and errors are returned.
func Append(p *Parser, f *Flag, args ...string) ([]string, error) {
	appendValue := func(p *Parser, f *Flag, value interface{}) {
		if p.Values[f.DestName] == nil || reflect.ValueOf(p.Values[f.DestName]).Kind() != reflect.Slice {
			p.Values[f.DestName] = make([]interface{}, 0)
		}
		p.Values[f.DestName] = append(p.Values[f.DestName].([]interface{}), value)
	}

	if regexp.MustCompile(`^[1-9]+$`).MatchString(f.ArgNum) == true {
		num, _ := strconv.Atoi(f.ArgNum)
		if len(args) < num {
			return args, fmt.Errorf("flag '%s' is expecting %d argument(s) but was provided %d", f.DisplayName(), num, len(args))
		}

		count := 0
		for count < num {
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
			appendValue(p, f, args[0])
			args = args[1:]
		} else {
			appendValue(p, f, f.DefaultVal)
		}
	} else {
		if f.ArgNum == "+" && len(args) == 0 {
			return args, fmt.Errorf("flag '%s' expects at least one argument", f.DisplayName())
		}

		for len(args) > 0 {
			appendValue(p, f, args[0])
			args = args[1:]
		}

		return args, nil
	}

	return args, nil
}

// AppendConst appends the flag's constant value into the parser. Provided arguments
// remain unmodified.
func AppendConst(p *Parser, f *Flag, args ...string) ([]string, error) {
	if f.ArgNum != "0" {
		panic(fmt.Sprintf("flag '%s' cannot expect any arguments.", f.DisplayName()))
	}

	if p.Values[f.DestName] == nil || reflect.ValueOf(p.Values[f.DestName]).Kind() != reflect.Slice {
		p.Values[f.DestName] = make([]interface{}, 0)
	}
	p.Values[f.DestName] = append(p.Values[f.DestName].([]interface{}), f.ConstVal)
	return args, nil
}

// ShowHelp calls the parser's ShowHelp function to output parser usage information
// and help information for each flag to stdout. Provided arguments remain unchanged.
func ShowHelp(p *Parser, f *Flag, args ...string) ([]string, error) {
	p.ShowHelp()
	return args, nil
}

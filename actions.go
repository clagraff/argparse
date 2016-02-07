package parg

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type Action func(*parser, *Flag, ...string) ([]string, error)

func Store(p *parser, f *Flag, args ...string) ([]string, error) {
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

		var values []string
		values = append(values, args[0:num]...)
		p.Values[f.DestName] = values
	}

	return args, nil
}

func StoreConst(p *parser, f *Flag, args ...string) ([]string, error) {
	if f.ArgNum != "0" {
		panic(fmt.Sprintf("flag '%s' cannot expect any arguments.", f.DisplayName()))
	}
	p.Values[f.DestName] = f.ConstVal

	return args, nil
}

func StoreFalse(p *parser, f *Flag, args ...string) ([]string, error) {
	if f.ArgNum != "0" {
		panic(fmt.Sprintf("flag '%s' cannot expect any arguments.", f.DisplayName()))
	}
	p.Values[f.DestName] = false

	return args, nil
}

func StoreTrue(p *parser, f *Flag, args ...string) ([]string, error) {
	if f.ArgNum != "0" {
		panic(fmt.Sprintf("flag '%s' cannot expect any arguments.", f.DisplayName()))
	}
	p.Values[f.DestName] = true

	return args, nil
}

func Append(p *parser, f *Flag, args ...string) ([]string, error) {
	appendValue := func(p *parser, f *Flag, value interface{}) {
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

func AppendConst(p *parser, f *Flag, args ...string) ([]string, error) {
	if f.ArgNum != "0" {
		panic(fmt.Sprintf("flag '%s' cannot expect any arguments.", f.DisplayName()))
	}

	if p.Values[f.DestName] == nil || reflect.ValueOf(p.Values[f.DestName]).Kind() != reflect.Slice {
		p.Values[f.DestName] = make([]interface{}, 0)
	}
	p.Values[f.DestName] = append(p.Values[f.DestName].([]interface{}), f.ConstVal)
	return args, nil
}

func ShowHelp(p *parser, f *Flag, args ...string) ([]string, error) {
	p.ShowHelp()
	return args, nil
}

// Package parg provides functionallity to emulate Python's argparse for
// setting-up and parsing a program's flags & arguments.
package parg

import (
	"fmt"
	"os"
	"strings"
)

// parser contains program-level settings and information, stores flags,
// and values collected upon parsing.
type parser struct {
	ProgramName string
	AllowAbbrev bool
	Flags       []*Flag
	UsageText   string
	Values      map[string]interface{}
}

// AddHelp adds a new flag to output usage information for the current parser
// and each of its flags.
func (p *parser) AddHelp() *parser {
	helpFlag := NewFlag("help", "Display usage information").Action(ShowHelp)
	shortHelpFlag := NewFlag("h", "Display usage information").Action(ShowHelp).Dest("help")

	p.Flags = append(p.Flags, helpFlag, shortHelpFlag)
	return p
}

// AddFlag appends the provided flag to the current parser.
func (p *parser) AddFlag(f *Flag) *parser {
	p.Flags = append(p.Flags, f)
	return p
}

// GetFlag retrieves the first flag with a public name matching the specified
// name, or will otherwise return an error.
func (p *parser) GetFlag(name string) (*Flag, error) {
	for _, flag := range p.Flags {
		if flag.PublicName == name {
			return flag, nil
		}
	}

	return nil, fmt.Errorf("No argument named: '%s'", name)
}

// GetHelp returns a string containing the parser's description text,
// and the usage information for each flag currently incorperated within
// the parser.
func (p *parser) GetHelp() string {
	// Get screen width to determine max line lengths later.
	screenWidth := getScreenWidth()

	var positional []*Flag
	var notPositional []*Flag
	var usage []string

	header := []string{"usage:", p.ProgramName}
	headerIndent := len(join(" ", header...))
	headerLen := headerIndent

	var notPosArgs []string
	var posArgs []string
	longest := 0

	for _, arg := range p.Flags {
		//if arg.IsPositional == false {
		notPositional = append(notPositional, arg)
		name := arg.GetUsage()
		if len(name) > longest {
			longest = len(name)
		}

		argUsg := arg.GetUsage()
		notPosArgs = append(notPosArgs, arg.GetUsage())
		headerLen = headerLen + len(argUsg)
		if headerLen+len(argUsg) > screenWidth {
			headerLen = headerIndent
			notPosArgs = append(notPosArgs, join("", "\n", spacer(headerIndent)))
		}
		/*} else {
			positional = append(positional, arg)
			name := arg.GetUsage()
			if len(name) > longest {
				longest = len(name)
			}

			argUsg := arg.GetUsage()
			posArgs = append(posArgs, arg.GetUsage())
			headerLen = headerLen + len(argUsg)
			if headerLen+len(argUsg) > screenWidth {
				headerLen = headerIndent
				posArgs = append(posArgs, join("", "\n", spacer(headerIndent)))
			}
		}*/
	}

	longest = longest + 4

	header = append(header, notPosArgs...)
	header = append(header, posArgs...)

	usage = append(usage, join(" ", header...), "\n")

	if len(p.UsageText) > 0 {
		usage = append(usage, "\n", p.UsageText, "\n")
	}

	if len(positional) > 0 {
		usage = append(usage, "\n", "positional arguments:", "\n")
		var names []string
		var help []string

		for _, arg := range positional {
			names = append(names, arg.GetUsage())
			help = append(help, arg.HelpText)
		}

		var lines []string
		for i, name := range names {
			lines = append(lines, "  ", name)
			lines = append(lines, spacer(longest-len(name)-2))
			if longest > screenWidth {
				lines = append(lines, "\n", spacer(longest))
			}
			lines = append(lines, help[i], "\n")
		}
		usage = append(usage, lines...)
	}

	if len(notPositional) > 0 {
		usage = append(usage, "\n", "optional arguments:", "\n")
		var names []string
		var help []string

		for _, arg := range notPositional {
			name := arg.GetUsage()

			names = append(names, name[1:len(name)-1])
			help = append(help, arg.HelpText)
		}

		var lines []string
		for i, name := range names {
			lines = append(lines, "  ", name)
			lines = append(lines, spacer(longest-len(name)-2))
			if longest > screenWidth {
				lines = append(lines, "\n", spacer(longest))
			}

			helpLines := wordWrap(help[i], screenWidth-longest)
			lines = append(lines, helpLines[0], "\n")
			if len(helpLines) > 1 {
				for _, helpLine := range helpLines[1:len(helpLines)] {
					lines = append(lines, spacer(longest), helpLine, "\n")
				}
			}
		}
		usage = append(usage, lines...)
	}

	return join("", usage...)
}

// Parser accepts a slice of strings as flags and arguments to be parsed. The
// parser will call each encountered flag's action. Unexpected flags will
// cause an error. All errors are returned.
func (p *parser) Parse(allArgs ...string) error {
	p.Values = make(map[string]interface{})
	for _, flag := range p.Flags {
		p.Values[flag.DestName] = flag.DefaultVal
	}

	flagNames, args := extractFlags(allArgs...)
	for _, flagName := range flagNames {
		var flag *Flag

		for _, f := range p.Flags {
			if flagName == f.PublicName {
				flag = f
				break
			}
		}

		if flag == nil {
			return fmt.Errorf("flag '%s' is not a valid flag.", flagName)
		}

		_, err := flag.DesiredAction(p, flag, args...)
		if err != nil {
			return err
		}
	}

	return nil
	/*
		for _, opt := range p.Arguments {
			if opt.IsPositional {
				positionals = append(positionals, opt)
			} else {
				flags = append(flags, opt)
			}
			p.Values[opt.Name] = opt.DefaultVal
		}

		count := 0
		max := len(args)
		for count < max {
			arg := args[count]
			if arg == "--" {
				count = count + 2
				continue
			}
			if isFlagFormat(arg) {
				flagName := getFlagName(arg)
				for _, flag := range flags {
					if flag.Name == flagName {
						if count+1 < max {
							_, err := flag.ActionType(p, flag, args[count+1:]...)
							if err != nil {
								fmt.Println(err.Error())
								break
							}
						} else {
							_, err := flag.ActionType(p, flag, "")
							if err != nil {
								fmt.Println(err.Error())
								break
							}
						}
					}
				}
			}
			count++
		}*/
}

// Path will set the parser's program name to the program name specified by the
// provided path.
func (p *parser) Path(progPath string) *parser {
	paths := strings.Split(progPath, string(os.PathSeparator))
	return p.Prog(paths[len(paths)-1])
}

// Prog sets the name of the parser directly.
func (p *parser) Prog(name string) *parser {
	p.ProgramName = name
	return p
}

// Usage sets the provide string as the usage/description text for the parser.
func (p *parser) Usage(usage string) *parser {
	p.UsageText = usage
	return p
}

// ShowHelp outputs to stdout the parser's generated help text.
func (p *parser) ShowHelp() *parser {
	fmt.Println(p.GetHelp())

	return p
}

// NewParser returns an instantiated pointer to a new parser instance, with
// a description matching the provided string.
func NewParser(desc string) *parser {
	p := parser{UsageText: desc}
	if len(os.Args) >= 1 {
		p.Path(os.Args[0])
	}
	return &p
}

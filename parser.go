package parg

import (
	"fmt"
	"os"
	"strings"
)

type parser struct {
	ProgramName     string
	UsageText       string
	DescriptionText string
	PrefixChar      string
	AddHelpArg      bool
	AllowAbbrev     bool
	Flags           []*Flag
	Values          map[string]interface{}
}

func (p *parser) AddHelp() *parser {
	return p.Help(true)
}

func (p *parser) AddArg(f *Flag) *parser {
	p.Flags = append(p.Flags, f)
	return p
}

func (p *parser) Desc(text string) *parser {
	p.DescriptionText = text
	return p
}

func (p *parser) GetArg(name string) (*Flag, error) {
	for _, flag := range p.Flags {
		if flag.PublicName == name {
			return flag, nil
		}
	}

	return nil, fmt.Errorf("No argument named: '%s'", name)
}

func (p *parser) GetHelp() string {
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

func (p *parser) Help(enable bool) *parser {
	p.AddHelpArg = enable
	return p
}

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

func (p *parser) Path(progPath string) *parser {
	paths := strings.Split(progPath, string(os.PathSeparator))
	return p.Prog(paths[len(paths)-1])
}

func (p *parser) Prog(name string) *parser {
	p.ProgramName = name
	return p
}

func (p *parser) RemoveHelp() *parser {
	return p.Help(false)
}

func (p *parser) Usage(usage string) *parser {
	p.UsageText = usage
	return p
}

func (p *parser) ShowHelp() *parser {
	fmt.Println(p.GetHelp())

	return p
}

func Parser(desc string) *parser {
	p := parser{AddHelpArg: true, AllowAbbrev: true, PrefixChar: "-"}
	if len(os.Args) >= 1 {
		p.Path(os.Args[0])
	}
	return &p
}

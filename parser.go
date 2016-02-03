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
	Arguments       []*argument
}

func (p *parser) AddHelp() *parser {
	return p.Help(true)
}

func (p *parser) AddArg(a *argument) *parser {
	p.Arguments = append(p.Arguments, a)
	return p
}

func (p *parser) Char(char string) *parser {
	p.PrefixChar = char
	return p
}

func (p *parser) Desc(text string) *parser {
	p.DescriptionText = text
	return p
}

func (p *parser) Help(enable bool) *parser {
	p.AddHelpArg = enable
	return p
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

func (p *parser) ShowHelp() {
	screenWidth := getScreenWidth()

	var positional []*argument
	var notPositional []*argument
	var usage []string

	header := []string{"usage:", p.ProgramName}
	headerIndent := len(join(" ", header...))
	headerLen := headerIndent

	var notPosArgs []string
	var posArgs []string
	longest := 0

	for _, arg := range p.Arguments {
		if arg.IsPositional == false {
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
		} else {
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
		}
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

	fmt.Println(join("", usage...))
}

func Parser(desc string) *parser {
	p := parser{AddHelpArg: true, AllowAbbrev: true, PrefixChar: "-"}
	if len(os.Args) >= 1 {
		p.Path(os.Args[0])
	}
	return &p
}

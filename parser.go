package parg

import (
	"bytes"
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
	var positional []*argument
	var notPositional []*argument
	var usage []string

	header := []string{"usage:", p.ProgramName}
	headerIndent := len(join(" ", header...))
	headerLen := headerIndent

	var notPosArgs []string
	var posArgs []string

	for _, arg := range p.Arguments {
		if arg.IsPositional == false {
			notPositional = append(notPositional, arg)

			argUsg := arg.GetUsage()
			notPosArgs = append(notPosArgs, arg.GetUsage())
			headerLen = headerLen + len(argUsg)
			if headerLen+len(argUsg) > 90 {
				var spacer bytes.Buffer
				count := 0
				for count < headerIndent {
					spacer.WriteString(" ")
					count++
				}
				headerLen = headerIndent
				notPosArgs = append(notPosArgs, join("", "\n", spacer.String()))
			}
		} else {
			positional = append(positional, arg)

			argUsg := arg.GetUsage()
			posArgs = append(posArgs, arg.GetUsage())
			headerLen = headerLen + len(argUsg)
			if headerLen+len(argUsg) > 90 {
				headerLen = headerIndent
				posArgs = append(posArgs, join("", "\n", spacer(headerIndent)))
			}
		}
	}

	header = append(header, notPosArgs...)
	header = append(header, posArgs...)

	usage = append(usage, join(" ", header...))

	if len(positional) > 0 {
		usage = append(usage, "\n", "positional arguments:", "\n")
		var names []string
		var help []string

		longest := 0

		for _, arg := range positional {
			names = append(names, arg.Name)
			help = append(help, arg.HelpText)
			if len(arg.Name) > longest {
				longest = len(arg.Name)
			}
		}
		longest = longest + 4

		var lines []string
		for i, name := range names {
			lines = append(lines, "  ", name)
			lines = append(lines, spacer(longest-len(name)-2))
			if longest > 80 {
				lines = append(lines, "\n", spacer(longest))
			}
			lines = append(lines, help[i])
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

package parg

import (
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

func Parser(desc string) *parser {
	p := parser{AddHelpArg: true, AllowAbbrev: true}
	if len(os.Args) >= 1 {
		p.Path(os.Args[0])
	}
	return &p
}

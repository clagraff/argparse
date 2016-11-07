package argparse

import (
	"fmt"
	"os"
	"strings"
)

type SubParser struct {
	Parser *Parser
	Name   string
}

// Parser contains program-level settings and information, stores options,
// and values collected upon parsing.
type Parser struct {
	Callback    func(*Parser, *Namespace, []string, error)
	ProgramName string
	AllowAbbrev bool
	Options     []*Option
	Parsers     []SubParser
	UsageText   string
	VersionDesc string
	Namespace   *Namespace
}

// AddHelp adds a new option to output usage information for the current parser
// and each of its options.
func (p *Parser) AddHelp() *Parser {
	helpOption := NewOption("h help", "help", "Show program help").Action(ShowHelp)

	p.Options = append(p.Options, helpOption)
	return p
}

// AddVersion adds a new option to the program version.
func (p *Parser) AddVersion() *Parser {
	versionOption := NewOption("v version", "version", "Show program version").Action(ShowVersion)

	p.Options = append(p.Options, versionOption)
	return p
}

// AddOption appends the provided option to the current parser.
func (p *Parser) AddOption(f *Option) *Parser {
	p.Options = append(p.Options, f)
	return p
}

// AddOptions appends the provided options to the current parser.
func (p *Parser) AddOptions(opts ...*Option) *Parser {
	for _, opt := range opts {
		p.Options = append(p.Options, opt)
	}
	return p
}

// AddParser appends the provided parse to the current parser as an available command.
func (p *Parser) AddParser(name string, parser *Parser) *Parser {
	if p.Parsers == nil {
		p.Parsers = make([]SubParser, 0)
	}
	p.Parsers = append(p.Parsers, SubParser{Name: name, Parser: parser})
	return p
}

// GetOption retrieves the first option with a public name matching the specified
// name, or will otherwise return an error.
func (p *Parser) GetOption(name string) (*Option, error) {
	if len(name) <= 0 {
		return nil, InvalidFlagNameErr{name}
	}
	for _, option := range p.Options {
		if option.IsPublicName(name) == true {
			return option, nil
		}
	}

	return nil, InvalidFlagNameErr{name}
}

// GetParser retrieves the desired sub-parser from the current parser, or returns
// an error if the desired parser does not exist.
func (p Parser) GetParser(name string) (*Parser, error) {
	if len(name) <= 0 {
		return nil, InvalidParserNameErr{name}
	}

	for _, subP := range p.Parsers {
		if subP.Name == name {
			return subP.Parser, nil
		}
	}

	return nil, InvalidParserNameErr{name}
}

// GetHelp returns a string containing the parser's description text,
// and the usage information for each option currently incorperated within
// the parser.
func (p *Parser) GetHelp() string {
	// Get screen width to determine max line lengths later.
	screenWidth := getScreenWidth()

	var commands []string
	var commandStr string
	var notPositional []*Option
	var positional []*Option
	var usage []string

	header := []string{"usage:", p.ProgramName}
	headerIndent := len(join(" ", header...))
	headerLen := headerIndent

	var notPosArgs []string
	var posArgs []string
	longest := 0

	for _, arg := range p.Options {
		if arg.IsPositional == false {
			notPositional = append(notPositional, arg)
		} else {
			positional = append(positional, arg)
		}
	}

	for _, arg := range notPositional {
		displayName := arg.DisplayName()
		if len(displayName) > longest {
			longest = len(displayName)
		}

		argUsg := arg.GetUsage()
		notPosArgs = append(notPosArgs, arg.GetUsage())
		headerLen = headerLen + len(argUsg)
		if headerLen+len(argUsg) > screenWidth {
			headerLen = headerIndent
			notPosArgs = append(notPosArgs, join("", "\n", spacer(headerIndent)))
		}
	}

	if len(p.Parsers) > 0 {
		for _, subP := range p.Parsers {
			commands = append(commands, subP.Name)
		}
		commandStr = join("", "{", join(",", commands...), "}")
		if len(commandStr) > longest {
			longest = len(commandStr)
		}

		posArgs = append(posArgs, commandStr)
		headerLen = headerLen + len(commandStr)
		if headerLen+len(commandStr) > screenWidth {
			headerLen = headerIndent
			posArgs = append(posArgs, join("", "\n", spacer(headerIndent)))
		}
	}

	for _, arg := range positional {
		displayName := arg.DisplayName()
		if len(displayName) > longest {
			longest = len(displayName)
		}

		argUsg := arg.GetUsage()
		posArgs = append(posArgs, arg.GetUsage())
		headerLen = headerLen + len(argUsg)
		if headerLen+len(argUsg) > screenWidth {
			headerLen = headerIndent
			posArgs = append(posArgs, join("", "\n", spacer(headerIndent)))
		}
	}

	longest = longest + 4

	header = append(header, notPosArgs...)
	header = append(header, posArgs...)

	usage = append(usage, join(" ", header...), "\n")

	if len(p.UsageText) > 0 {
		usage = append(usage, "\n", p.UsageText, "\n")
	}

	if len(positional) > 0 || len(commandStr) > 0 {
		usage = append(usage, "\n", "positional arguments:", "\n")
		var names []string
		var help []string

		if len(commandStr) > 0 {
			names = append(names, commandStr)
			help = append(help, "commands")
		}

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

	if len(notPositional) > 0 {
		usage = append(usage, "\n", "optional arguments:", "\n")
		var names []string
		var help []string

		for _, arg := range notPositional {
			names = append(names, arg.DisplayName())
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

// GetVersion will return the version text for the current parser.
func (p *Parser) GetVersion() string {
	return p.ProgramName + " version " + p.VersionDesc
}

// Parser accepts a slice of strings as options and arguments to be parsed. The
// parser will call each encountered option's action. Unexpected options will
// cause an error. All errors are returned.
func (p *Parser) Parse(allArgs ...string) {
	if p.Namespace == nil {
		p.Namespace = NewNamespace()
	}
	requiredOptions := make(map[string]*Option)
	remainderOptions := make(map[string]*Option)
	var err error

	var optionListing []*Option

	for _, option := range p.Options {
		if option.IsRequired == true {
			requiredOptions[option.DestName] = option
		}
		p.Namespace.Set(option.DestName, option.DefaultVal)
		if strings.ToLower(option.ArgNum) == "r" {
			remainderOptions[option.DisplayName()] = option
		} else {
			optionListing = append(optionListing, option)
		}
	}

	if len(p.Parsers) > 0 {
		var usedParser bool

		for _, subParser := range p.Parsers {
			name := subParser.Name
			parser := subParser.Parser
			if len(allArgs) <= 0 {
				p.Callback(p, p.Namespace, nil, MissingParserErr{p.Parsers})
				return
			}
			if allArgs[0] == name {
				if len(allArgs) > 0 {
					allArgs = allArgs[1:len(allArgs)]
				} else {
					allArgs = []string{}
				}

				parser.Parse(allArgs...)
				return
			}
		}

		if usedParser != true {
			p.Callback(p, p.Namespace, nil, MissingParserErr{p.Parsers})
		}
	}

	optionNames, args := extractOptions(allArgs...)
	for _, optionName := range optionNames {
		var option *Option

		for _, f := range p.Options {
			if f.IsPositional == true {
				continue
			}
			if f.IsPublicName(optionName) == true {
				if _, ok := requiredOptions[f.DisplayName()]; ok {
					delete(requiredOptions, f.DisplayName())
				} else if _, ok := remainderOptions[f.DisplayName()]; ok {
					continue
				}
				option = f
				break
			}
		}

		if option == nil {
			p.Callback(p, p.Namespace, args, InvalidOptionErr{optionName})
		}

		args, err = option.DesiredAction(p, option, args...)
		if err != nil {
			p.Callback(p, p.Namespace, args, err)
		}
	}

	if len(args) > 0 {
		for _, opt := range remainderOptions {
			if _, ok := requiredOptions[opt.DisplayName()]; ok {
				delete(requiredOptions, opt.DisplayName())
			}
			_, err := opt.DesiredAction(p, opt, args...)
			if err != nil {
				p.Callback(p, p.Namespace, args, err)
			}
		}
	}

	for _, f := range p.Options {
		if f.IsPositional == false {
			continue
		}
		if _, ok := requiredOptions[f.DestName]; ok {
			delete(requiredOptions, f.DestName)
		}
		args, err = f.DesiredAction(p, f, args...)
		if err != nil {
			p.Callback(p, p.Namespace, args, err)
		}
	}

	if len(requiredOptions) != 0 {
		for _, option := range requiredOptions {
			p.Callback(p, p.Namespace, args, MissingOptionErr{option.DisplayName()})
		}
	}
	p.Callback(p, p.Namespace, args, nil)
}

// Path will set the parser's program name to the program name specified by the
// provided path.
func (p *Parser) Path(progPath string) *Parser {
	paths := strings.Split(progPath, string(os.PathSeparator))
	return p.Prog(paths[len(paths)-1])
}

// Prog sets the name of the parser directly.
func (p *Parser) Prog(name string) *Parser {
	p.ProgramName = name
	return p
}

// Usage sets the provide string as the usage/description text for the parser.
func (p *Parser) Usage(usage string) *Parser {
	p.UsageText = usage
	return p
}

// ShowHelp outputs to stdout the parser's generated help text.
func (p *Parser) ShowHelp() *Parser {
	fmt.Println(p.GetHelp())

	return p
}

// ShowVersion outputs to stdout the parser's generated versioning text.
func (p *Parser) ShowVersion() *Parser {
	fmt.Println(p.GetVersion())

	return p
}

// Version sets the provide string as the version text for the parser.
func (p *Parser) Version(version string) *Parser {
	p.VersionDesc = version
	return p
}

// NewParser returns an instantiated pointer to a new parser instance, with
// a description matching the provided string.
func NewParser(desc string, callback func(*Parser, *Namespace, []string, error)) *Parser {
	p := Parser{UsageText: desc}
	p.Namespace = NewNamespace()
	p.Parsers = make([]SubParser, 0)
	if len(os.Args) >= 1 {
		p.Path(os.Args[0])
	}
	p.Callback = callback
	return &p
}

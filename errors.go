package argparse

import (
	"fmt"
	"strings"
)

// InvalidChoiceErr indicates that an argument is not among the valid choices
// for the option.
type InvalidChoiceErr struct {
	opt Option
	arg string
}

// Error will return a string error message for the InvalidChoiceErr
func (err InvalidChoiceErr) Error() string {
	msg := "%s: invalid choice \"%s\" (choose from: %s)"
	return fmt.Sprintf(msg, err.opt.DisplayName(), err.arg, strings.Join(err.opt.ValidChoices, ", "))

}

// InvalidParserNameErr indicates that a Command name has already been assigned and cannot be re-assigned.
type InvalidParserNameErr struct {
	name string
}

// Error will return a string error message for the InvalidParserNameErr
func (err InvalidParserNameErr) Error() string {
	msg := "invalid command name \"%s\""
	return fmt.Sprintf(msg, err.name)

}

// InvalidFlagNameErr indicates that an argument with the provided public name
// not exist.
type InvalidFlagNameErr struct {
	name string
}

// Error will return a string error message for the InvalidFlagNameErr
func (err InvalidFlagNameErr) Error() string {
	msg := "invalid flag name \"%s\""
	return fmt.Sprintf(msg, err.name)

}

// InvalidOptionErr indicates that an option is invalid.
type InvalidOptionErr struct {
	name string
}

// Error will return a string error message for the InvalidFlagNameErr
func (err InvalidOptionErr) Error() string {
	msg := "invalid option \"%s\""
	return fmt.Sprintf(msg, err.name)

}

// InvalidTypeErr indicates that an argument cannot be casted the the option's
// expected type.
type InvalidTypeErr struct {
	opt Option
	arg string
}

// Error will return a string error message for the InvalidTypeErr
func (err InvalidTypeErr) Error() string {
	msg := "%s: invalid %s value: \"%s\""
	return fmt.Sprintf(msg, err.opt.DisplayName(), err.opt.ExpectedType.String(), err.arg)
}

// MissingEnvVarErr indicates that an environmental variable could not be found
// with the provided variable name.
type MissingEnvVarErr struct {
	varName string
}

// Error will return a string error message for the MissingEnvVarErr.
func (err MissingEnvVarErr) Error() string {
	msg := "missing environmental variable \"%s\""
	return fmt.Sprintf(msg, err.varName)
}

// ShowHelpErr indicates that the program was instructed to show it's help text.
type ShowHelpErr struct{}

func (err ShowHelpErr) Error() string { return "" }

// ShowVersionErr indicates that the program was instructed to show it's versioning text.
type ShowVersionErr struct{}

func (err ShowVersionErr) Error() string { return "" }

// TooFewArgsErr indicated that not enough arguments were provided for the option.
type TooFewArgsErr struct {
	opt Option
}

// Error will return a string error message for the TooFewArgsErr
func (err TooFewArgsErr) Error() string {
	msg := "%s: too few arguments"
	return fmt.Sprintf(msg, err.opt.DisplayName())
}

// MissingOneOrMoreArgsErr indicated that not enough arguments were provided,
// when one or more arguments were expected, for the option.
type MissingOneOrMoreArgsErr struct {
	opt Option
}

// Error will return a string error message for the TooFewArgsErr
func (err MissingOneOrMoreArgsErr) Error() string {
	msg := "%s: at least one argument requireds"
	return fmt.Sprintf(msg, err.opt.DisplayName())
}

// MissingParserErr indicated that commands were available, but none were used.
type MissingParserErr struct {
	Parsers []SubParser
}

// Error will return a string error message for the MissingParserErr
func (err MissingParserErr) Error() string {
	var names []string
	for _, subP := range err.Parsers {
		names = append(names, subP.Name)
	}
	msg := "must use an available command: %s"
	return fmt.Sprintf(msg, join("", "{", join(",", names...), "}"))
}

// MissingOptionErr indicated that an option was required but is missing.
type MissingOptionErr struct {
	name string
}

// Error will return a string error message for the MissingOptionErr
func (err MissingOptionErr) Error() string {
	msg := "option \"%s\" required"
	return fmt.Sprintf(msg, err.name)
}

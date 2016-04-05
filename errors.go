package argparse

import (
	"fmt"
	"strings"
)

// ShowHelpErr indicates that the program was instructed to show it's help text.
type ShowHelpErr struct{}

func (err ShowHelpErr) Error() string { return "" }

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

// MissingOptionErr indicated that an option was required but is missing.
type MissingOptionErr struct {
	name string
}

// Error will return a string error message for the MissingOptionErr
func (err MissingOptionErr) Error() string {
	msg := "option \"%s\" required"
	return fmt.Sprintf(msg, err.name)
}

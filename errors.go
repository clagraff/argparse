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

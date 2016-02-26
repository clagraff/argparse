package argparse

import (
	"fmt"
	"strings"
)

type InvalidChoiceErr struct {
	opt Option
	arg string
}

func (err InvalidChoiceErr) Error() string {
	msg := "%s: invalid choice \"%s\" (choose from: %s)"
	return fmt.Sprintf(msg, err.opt.DisplayName(), err.arg, strings.Join(err.opt.ValidChoices, ", "))

}

type InvalidTypeErr struct {
	opt Option
	arg string
}

func (err InvalidTypeErr) Error() string {
	msg := "%s: invalid %s value: \"%s\""
	return fmt.Sprintf(msg, err.opt.DisplayName(), err.opt.ExpectedType.String(), err.arg)
}

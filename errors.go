package argparse

import (
	"fmt"
	"strings"
)

func InvalidChoiceErr(opt Option, arg string) error {
	msg := "%s: invalid choice \"%s\" (choose from: %s)"
	return fmt.Errorf(msg, opt.DisplayName(), arg, strings.Join(opt.ValidChoices, ", "))
}

func InvalidTypeErr(opt Option, arg string) error {
	msg := "%s: invalid type \"%s\" (must be: %s)"
	return fmt.Errorf(msg, opt.DisplayName(), arg, opt.ExpectedType.String())
}

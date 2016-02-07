package parg

import (
	"strings"
	"testing" //import go package for testing related functionality
)

// TestExtractFlags_NoArgs tests to ensure that when no arguments are provided,
// no flags or arguments are returned by extractFlags.
func TestExtractFlags_NoArgs(t *testing.T) {
	var noArgs []string
	flags, args := extractFlags(noArgs...)

	if len(flags) != 0 {
		t.Error("No flags should have been extracted")
	}

	if len(args) != 0 {
		t.Error("No arguments should have been extracted")
	}
}

// TestExtractFlags_OnlyFlags tests to ensure that if only flag arguments are
// provided, that the same number of flag arguments are returned & with no
// additional arguments by extractFlags.
func TestExtractFlags_OnlyFlags(t *testing.T) {
	onlyFlagArgs := []string{"-f", "--foobar"}
	flags, args := extractFlags(onlyFlagArgs...)

	if len(flags) != len(onlyFlagArgs) {
		t.Errorf(
			"%d number of flags expected, but only %d were extracted",
			len(onlyFlagArgs),
			len(flags),
		)
	}

	if len(args) != 0 {
		t.Error("No arguments should have been extracted")
	}
}

// TestExtractFlags_MutliShortFlags tests to ensure that multiple short-flags
// residing beside each other are properly recognized and extract individually,
// and no other arguments are returned.
func TestExtractFlags_MultiShortFlags(t *testing.T) {
	shortFlags := []string{"a", "b", "c"}
	shortFlagArgs := []string{"-" + strings.Join(shortFlags, "")}

	flags, args := extractFlags(shortFlagArgs...)

	if len(flags) != len(shortFlags) {
		t.Errorf(
			"%d number of flags expected, but only %d were extracted",
			len(shortFlags),
			len(flags),
		)
	}

	if len(args) != 0 {
		t.Error("No arguments should have been extracted")
	}
}

// TestExtractFlags_OnlyArgs tests to ensure that if only passive arguments are
// provided, that the same number of passive arguments are returned & with no
// flags extracted by extractFlags.
func TestExtractFlags_OnlyArgs(t *testing.T) {
	onlyArgs := []string{"arg1", "arg2", "arg3", "arg4"}
	flags, args := extractFlags(onlyArgs...)

	if len(args) != len(onlyArgs) {
		t.Errorf(
			"%d number of passive argumentss expected, but only %d were extracted",
			len(onlyArgs),
			len(args),
		)
	}

	if len(flags) != 0 {
		t.Error("No flags should have been extracted")
	}
}

// TestExtractFlags tests to ensure that the expected number of flags & passive
// arguments are extracted by extractFlags.
func TestExtractFlags(t *testing.T) {
	allArgs := []string{"-f", "foobar", "--fizzbuzz", "four", "five", "-irtusc"}
	numFlags := 8
	numArgs := 3

	flags, args := extractFlags(allArgs...)

	if len(flags) != numFlags {
		t.Errorf(
			"%d number of flags expected, but only %d were extracted",
			numFlags,
			len(flags),
		)
	}

	if len(args) != numArgs {
		t.Errorf(
			"%d number of passive argumentss expected, but only %d were extracted",
			numArgs,
			len(args),
		)
	}
}

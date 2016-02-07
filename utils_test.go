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

// TestGetScreenWidth tests to ensure that a positive, non-zero integer value is returned
// to represent the width of the current screen.
func TestGetScreenWidth(t *testing.T) {
	// I am not really sure the best way to test this.
	// TODO: make a better test!

	width := getScreenWidth()
	if width <= 0 {
		t.Error("Retrieved screen width should be a positive, non-zero integer")
	}
}

// TestJoin tests to ensure that a variety of string slices can be joined in the
// correct, expected manner.
func TestJoin(t *testing.T) {
	testStrings := [][]string{
		[]string{"one", "two"},
		[]string{""},
		[]string{"three", "four", "five"},
	}

	expectedStrings := []string{
		"one two",
		"",
		"three four five",
	}

	for i, test := range testStrings {
		actual := join(" ", test...)
		expected := expectedStrings[i]

		if actual != expected {
			t.Errorf(
				"Expected: '%s' but received: '%s'",
				expected,
				actual,
			)
		}
	}

	actual := join("-*-", "abc", "def")
	expected := "abc-*-def"
	if actual != expected {
		t.Errorf(
			"Expected: '%s' but received: '%s'",
			expected,
			actual,
		)
	}
}

// TestSpacer tests to make sure the proper length strings are returned, as expected.
func TestSpacer(t *testing.T) {
	intTests := []int{-1000, -100, -10, -1, 0, 1, 10, 100, 1000}

	for _, test := range intTests {
		actual := spacer(test)
		expectedLen := test
		if expectedLen < 0 {
			expectedLen = 0
		}

		if len(actual) != expectedLen {
			if len(actual) != 0 {
				t.Errorf("Expected string of length: %d but received: '%s'", expectedLen, actual)
			}
		}
	}
}

package argparse

import (
	"os"
	"strings"
	"testing" //import go package for testing related functionality
)

// TestExtractOptions_NoArgs tests to ensure that when no arguments are provided,
// no options or arguments are returned by extractOptions.
func TestExtractOptions_NoArgs(t *testing.T) {
	var noArgs []string
	options, args := extractOptions(noArgs...)

	if len(options) != 0 {
		t.Error("No options should have been extracted")
	}

	if len(args) != 0 {
		t.Error("No arguments should have been extracted")
	}
}

// TestExtractOptions_OnlyOptions tests to ensure that if only option arguments are
// provided, that the same number of option arguments are returned & with no
// additional arguments by extractOptions.
func TestExtractOptions_OnlyOptions(t *testing.T) {
	onlyOptionArgs := []string{"-f", "--foobar"}
	options, args := extractOptions(onlyOptionArgs...)

	if len(options) != len(onlyOptionArgs) {
		t.Errorf(
			"%d number of options expected, but only %d were extracted",
			len(onlyOptionArgs),
			len(options),
		)
	}

	if len(args) != 0 {
		t.Error("No arguments should have been extracted")
	}
}

// TestExtractOptions_MutliShortOptions tests to ensure that multiple short-options
// residing beside each other are properly recognized and extract individually,
// and no other arguments are returned.
func TestExtractOptions_MultiShortOptions(t *testing.T) {
	shortOptions := []string{"a", "b", "c"}
	shortOptionArgs := []string{"-" + strings.Join(shortOptions, "")}

	options, args := extractOptions(shortOptionArgs...)

	if len(options) != len(shortOptions) {
		t.Errorf(
			"%d number of options expected, but only %d were extracted",
			len(shortOptions),
			len(options),
		)
	}

	if len(args) != 0 {
		t.Error("No arguments should have been extracted")
	}
}

// TestExtractOptions_OnlyArgs tests to ensure that if only passive arguments are
// provided, that the same number of passive arguments are returned & with no
// options extracted by extractOptions.
func TestExtractOptions_OnlyArgs(t *testing.T) {
	onlyArgs := []string{"arg1", "arg2", "arg3", "arg4"}
	options, args := extractOptions(onlyArgs...)

	if len(args) != len(onlyArgs) {
		t.Errorf(
			"%d number of passive argumentss expected, but only %d were extracted",
			len(onlyArgs),
			len(args),
		)
	}

	if len(options) != 0 {
		t.Error("No options should have been extracted")
	}
}

// TestExtractOptions tests to ensure that the expected number of options & passive
// arguments are extracted by extractOptions.
func TestExtractOptions(t *testing.T) {
	allArgs := []string{"-f", "foobar", "--fizzbuzz", "four", "five", "-irtusc"}
	numOptions := 8
	numArgs := 3

	options, args := extractOptions(allArgs...)

	if len(options) != numOptions {
		t.Errorf(
			"%d number of options expected, but only %d were extracted",
			numOptions,
			len(options),
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

// TestGetEnvVar tests to ensure that we can grab the value of a specified
// environmental variable if it exists, or otherwise error out.
func TestGetEnvVar(t *testing.T) {
	key := "FAKE_EXIST_ENV_VAR_1337"
	val := "value chosen at random"
	if err := os.Setenv(key, val); err != nil {
		t.Fatal(err)
	}

	if str, err := getEnvVar("$" + key); err != nil {
		t.Errorf(
			"no error was expected, but an error was returned: %s",
			err.Error(),
		)
	} else {
		if str != val {
			t.Errorf(
				"\"%s\" expected value does not match returned value \"%s\"",
				val,
				str,
			)
		}
	}

	key = "FAKE_DOES_NOT_EXIST"
	if _, err := getEnvVar("$" + key); err == nil {
		t.Error("An error was expected, but no errors were returned")
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

// TestIsEnvVarFormat is a table-test to ensure that isEnvVarFormat will
// return the expected result for a given slice of inputs representing
// possible environmental names.
func TestIsEnvVarFormat(t *testing.T) {
	table := map[string]bool{
		"":                          false,
		"$":                         false,
		"$0":                        false,
		"0":                         false,
		"_":                         false,
		"abc":                       false,
		"DEF":                       false,
		"VAR$":                      false,
		"$TEST":                     true,
		"$abc":                      true,
		"$FOO_BAR":                  true,
		"$_":                        true,
		"$LEET_1337":                true,
		"$z28":                      true,
		"$one_222_THREE":            true,
		"$MANY_______und3rscor3s__": true,
	}

	for testStr, expectedBool := range table {
		result := isEnvVarFormat(testStr)
		if result != expectedBool {
			t.Errorf(
				"envVarFormat expected the string \"%s\" to return: %t, but got %t instead",
				testStr,
				expectedBool,
				result,
			)
		}
	}
}

// TestJoin tests to ensure that a variety of string slices can be joined in the
// correct, expected manner.
func TestJoin(t *testing.T) {
	testStrings := [][]string{
		{"one", "two"},
		{""},
		{"three", "four", "five"},
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

// TestWordWrap tests to ensure strings will be broken into the appropriate
// word-length limited slice of strings.
func TestWordWrap(t *testing.T) {
	oneLine := "This text is below the limit."
	threeLines := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vestibulum dolor justo, tempor quis"

	if len(wordWrap(oneLine, 35)) != 1 {
		t.Error("wordWrap did not return a slice of lenth 1")
	}

	if len(wordWrap(threeLines, 35)) != 3 {
		t.Error("wordWrap did not return a slice of length 3")
	}
}

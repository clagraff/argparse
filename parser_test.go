package parg

import (
	"bufio"
	"os"
	"testing"
) //import go package for testing related functionality

// TestParserAddHelp tests the AddHelp method to ensure two help flags
// are appended to the parser, a short flag & a long flag.
func TestParserAddHelp(t *testing.T) {
	p := Parser{}
	if len(p.Flags) != 0 {
		t.Error("Parser should not contain any flags")
	}

	p.AddHelp()
	if len(p.Flags) != 2 {
		t.Errorf("Expected 2 flags, but only has %d", len(p.Flags))
	}
}

// TestParserAddFlag tests the AddFlah method to ensure that flags can be appended
// to a parser.
func TestParserAddFlag(t *testing.T) {
	p := Parser{}

	if len(p.Flags) != 0 {
		t.Error("Parser should not contain any flags")
	}

	f1 := Flag{}
	f2 := Flag{}

	p.AddFlag(&f1)
	if len(p.Flags) != 1 {
		t.Error("Parser should contain one flag")
	}

	p.AddFlag(&f2)
	if len(p.Flags) != 2 {
		t.Error("Parser should contain two flags")
	}
}

// TestParserGetFlag_InvalidFlag tests retreival of an error and nil for a flag
// from a Parser instance by specifying an incorrect PublicName attribute.
func TestParserGetFlag_InvalidFlag(t *testing.T) {
	f1 := NewFlag("first", "this is the first flag")
	f2 := NewFlag("second", "this is the second flag")
	f3 := NewFlag("three", "this is the third flag")

	p := Parser{}
	p.AddFlag(f1).AddFlag(f2).AddFlag(f3)

	if len(p.Flags) != 3 {
		t.Error("Parser should contain three flags")
	}

	f, err := p.GetFlag("twenty")
	if err == nil {
		t.Errorf("An error was expected but did not occur.")
	}

	if f != nil {
		t.Error("The retrived flag was not nil")
	}
}

// TestParserGetFlag_ValidFlag tests retreival of flags from a Parser instance by specifying
// their PublicName attribute. A valid flag PublicName will be used to retrieve a flag.
func TestParserGetFlag_ValidFlag(t *testing.T) {
	f1 := NewFlag("first", "this is the first flag")
	f2 := NewFlag("second", "this is the second flag")
	f3 := NewFlag("three", "this is the third flag")

	p := Parser{}
	p.AddFlag(f1).AddFlag(f2).AddFlag(f3)

	if len(p.Flags) != 3 {
		t.Error("Parser should contain three flags")
	}

	f, err := p.GetFlag("second")
	if err != nil {
		t.Errorf("An unexpected error occurred: %s", err.Error())
	}

	if f == nil {
		t.Error("The retrived flag cannot be nil")
	} else if f.PublicName != "second" {
		t.Errorf("Expected flag name 'second', but retrieved name: %s", f.PublicName)
	}
}

// TestParserGetHelp tests the GetHelp method to ensure that the parser will
// return a help-string containing usage information and flag-dependent
// help text.
func TestParserGetHelp(t *testing.T) {
	p := NewParser("this is a description of the program")

	if len(p.GetHelp()) <= 0 {
		t.Errorf("A non-empty string was expected but was not received")
	}

	// TODO: implement a better test for the Parser.GetHelp() method.
}

// TestParserParse tests the Parse method to ensure that arguments provided to
// the parser are properly parsed and the necessary actions for all flags are
// executed.
func TestParserParse(t *testing.T) {
	// TODO: create an actual test.
}

// TestParserPath tests the Path method to ensure that providing a filepath will
// result in updating the parser's program name.
func TestParserPath(t *testing.T) {
	path := "/usr/local/bin/my_prog"
	expected := "my_prog"

	p := Parser{}
	if len(p.ProgramName) != 0 {
		t.Errorf("Parser program name should be an empty string, but is: %s", p.ProgramName)
	}

	p.Path(path)
	if p.ProgramName != expected {
		t.Errorf(
			"The parser's ProgramName '%s' does not match the expected name: '%s'",
			p.ProgramName,
			expected,
		)
	}
}

// TestParserProg tests the Prog method to ensure that providing a program name
// will result in updating the parser's program name string.
func TestParserProg(t *testing.T) {
	name := "awesome_go_prog"

	p := Parser{}
	if len(p.ProgramName) != 0 {
		t.Errorf("Parser program name should be an empty string, but is: %s", p.ProgramName)
	}

	p.Prog(name)
	if p.ProgramName != name {
		t.Errorf(
			"The parser's ProgramName '%s' does not match the expected name: '%s'",
			p.ProgramName,
			name,
		)
	}
}

// TestParserUsage tests the Usage method to ensure that providing a usage string
// will result in updating the parser's usage string.
func TestParserUsage(t *testing.T) {
	usage := "do stuff to accomplish things"

	p := Parser{}
	if len(p.UsageText) != 0 {
		t.Errorf("Parser usage text should be an empty string, but is: %s", p.UsageText)
	}

	p.Usage(usage)
	if p.UsageText != usage {
		t.Errorf(
			"The parser's UsageText '%s' does not match the expected usage: '%s'",
			p.UsageText,
			usage,
		)
	}
}

// TestParserShowHelp tests the ShowHelp method to ensure the parser will print
// the tet returned by GetHelp to stdout.
func TestParserShowHelp(t *testing.T) {
	oldStdout := os.Stdout
	_, writeFile, err := os.Pipe()
	if err != nil {
		t.Error(err.Error())
	}
	os.Stdout = writeFile

	p := NewParser("this is an awesome program which does awesome things")
	p.ShowHelp()

	writeFile.Close()
	os.Stdout = oldStdout

	var lines []string
	scanner := bufio.NewScanner(writeFile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if len(lines) < 0 {
		t.Error("Expected help text but none was received")
	}
}

// TestNewParser tests to ensure a new parser with a populated description
// is returned using the NewParser function.
func TestNewParser(t *testing.T) {
	desc := "this program does things."
	p := NewParser(desc)

	if p == nil {
		t.Error("The parser pointer cannot be null")
	}

	if p.UsageText != desc {
		t.Error("The parser's usage text: '%s' does not match the expected description: '%s'", p.UsageText, desc)
	}
}

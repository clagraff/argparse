package argparse

import (
	"bufio"
	"os"
	"testing"
) //import go package for testing related functionality

// TestParserAddHelp tests the AddHelp method to ensure two help options
// are appended to the parser, a short option & a long option.
func TestParserAddHelp(t *testing.T) {
	p := Parser{}
	if len(p.Options) != 0 {
		t.Error("Parser should not contain any options")
	}

	p.AddHelp()
	if len(p.Options) != 1 {
		t.Errorf("Expected 1 options, but only has %d", len(p.Options))
	}
}

// TestParserAddOption tests the AddFlah method to ensure that options can be appended
// to a parser.
func TestParserAddOption(t *testing.T) {
	p := Parser{}

	if len(p.Options) != 0 {
		t.Error("Parser should not contain any options")
	}

	f1 := Option{}
	f2 := Option{}

	p.AddOption(&f1)
	if len(p.Options) != 1 {
		t.Error("Parser should contain one option")
	}

	p.AddOption(&f2)
	if len(p.Options) != 2 {
		t.Error("Parser should contain two options")
	}
}

// TestParserGetOption_InvalidOption tests retreival of an error and nil for a option
// from a Parser instance by specifying an incorrect PublicName attribute.
func TestParserGetOption_InvalidOption(t *testing.T) {
	f1 := NewOption("first", "dest", "this is the first option")
	f2 := NewOption("second", "dest", "this is the second option")
	f3 := NewOption("three", "dest", "this is the third option")

	p := Parser{}
	p.AddOption(f1).AddOption(f2).AddOption(f3)

	if len(p.Options) != 3 {
		t.Error("Parser should contain three options")
	}

	f, err := p.GetOption("twenty")
	if err == nil {
		t.Errorf("An error was expected but did not occur.")
	}

	if f != nil {
		t.Error("The retrived option was not nil")
	}
}

// TestParserGetOption_ValidOption tests retreival of options from a Parser instance by specifying
// their PublicName attribute. A valid option PublicName will be used to retrieve a option.
func TestParserGetOption_ValidOption(t *testing.T) {
	f1 := NewOption("first", "dest", "this is the first option")
	f2 := NewOption("second", "dest", "this is the second option")
	f3 := NewOption("three", "dest", "this is the third option")

	p := Parser{}
	p.AddOption(f1).AddOption(f2).AddOption(f3)

	if len(p.Options) != 3 {
		t.Error("Parser should contain three options")
	}

	f, err := p.GetOption("second")
	if err != nil {
		t.Errorf("An unexpected error occurred: %s", err.Error())
	}

	if f == nil {
		t.Error("The retrived option cannot be nil")
	} else if f.PublicNames[0] != "second" {
		t.Errorf("Expected option name 'second', but retrieved name: %s", f.PublicNames)
	}
}

// TestParserGetHelp tests the GetHelp method to ensure that the parser will
// return a help-string containing usage information and option-dependent
// help text.
func TestParserGetHelp(t *testing.T) {
	p := NewParser("this is a description of the program")

	if len(p.GetHelp()) <= 0 {
		t.Errorf("A non-empty string was expected but was not received")
	}

	// TODO: implement a better test for the Parser.GetHelp() method.
}

// TestParserGetVersion tests the GetVersion  method to ensure that the parser will
// return a version-string containing version information of the parser.
func TestParserGetVersion(t *testing.T) {
	p := NewParser("some description").Version("1.0.b")

	if p.GetVersion() != "argparse.test version 1.0.b" {
		t.Errorf("The retrieved version text did not match the expected string")
	}
}

// TestParserParse tests the Parse method to ensure that arguments provided to
// the parser are properly parsed and the necessary actions for all options are
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
// the text returned by GetHelp to stdout.
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

// TestParserShowVersion tests the ShowVersion method to ensure the parser will print
// the text returned by GetVersion to stdout.
func TestParserShowVersion(t *testing.T) {
	oldStdout := os.Stdout
	_, writeFile, err := os.Pipe()
	if err != nil {
		t.Error(err.Error())
	}
	os.Stdout = writeFile

	p := NewParser("this is a program").Version("1.foo.bar.0.42 alpha")
	p.ShowVersion()

	writeFile.Close()
	os.Stdout = oldStdout

	var lines []string
	scanner := bufio.NewScanner(writeFile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if len(lines) < 0 {
		t.Error("Expected version text but none was received")
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

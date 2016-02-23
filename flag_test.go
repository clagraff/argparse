package parg

import "testing"

//import go package for testing related functionality

// TestNewFlag tests the creation of a new flag, populated with defaults
// and appropriate name and description as provided.
func TestNewFlag(t *testing.T) {
	name := "flag1"
	desc := "my flag"
	f := NewFlag(name, desc)

	if f.DefaultVal != nil {
		t.Error("Default value should be nil")
	}

	if f.ConstVal != nil {
		t.Error("Constant value should be nil")
	}

	if f.DestName != name {
		t.Error("DestName value should match expected name")
	}

	if f.PublicName != name {
		t.Error("PublicName value should match expected name")
	}

	if len(f.MetaVarText) != 1 {
		t.Error("MetaVarText does not contain only one element")
	} else if f.MetaVarText[0] != name {
		t.Error("MetaVarText[0] does not match the expected name")
	}
}

// TestFlagAction tests the Action method to ensure a flag's DesiredAction
// will become the provided action.
func TestFlagAction(t *testing.T) {
	testAction := func(p *Parser, f *Flag, args ...string) ([]string, error) { return nil, nil }
	f := Flag{}

	f.Action(testAction)
	if f.DesiredAction == nil {
		t.Error("Flag action was not properly set")
	}
}

// TestFlagChoices tests that the Choices method will set a Flags avaliable choices
// to the provided []interface{}.
func TestFlagChoices(t *testing.T) {
	f := Flag{}
	choices := []interface{}{"foobar", true, 12}

	if len(f.PossibleChoices) != 0 {
		t.Error("Flag should not contain any choices yet")
	}

	f.Choices(choices)
	if len(f.PossibleChoices) != len(choices) {
		t.Errorf("Flag contains %d choices, but is expected a total of %d", len(f.PossibleChoices), len(choices))
	}
}

// TestFlagConst tests that a flag's ConstValue is updated to the provided value
// via the Const method.
func TestFlagConst(t *testing.T) {
	f := Flag{}

	if f.ConstVal != nil {
		t.Error("Flag ConstVal should be nil upon initialization")
	}

	expected := "some value"
	f.Const(expected)

	if f.ConstVal != expected {
		t.Errorf("Flag ConstVal is '%v', but was expected to be: '%s'", f.ConstVal, expected)
	}
}

// TestFlagDefault tests that a flag's DefaultValue is updated to the provided value
// via the Default method.
func TestFlagDefault(t *testing.T) {
	f := Flag{}

	if f.DefaultVal != nil {
		t.Error("Flag DefaultVal should be nil upon initialization")
	}

	expected := "some value"
	f.Default(expected)

	if f.DefaultVal != expected {
		t.Errorf("Flag DefaultVal is '%v', but was expected to be: '%s'", f.DefaultVal, expected)
	}
}

// TestFlagDisplayName tests the retrival of a flag's display name, with an
// appropriate number of preceeding hypens, via the DisplayName method.
func TestFlagDisplayName(t *testing.T) {
	f := Flag{}

	expected := ""
	name := f.DisplayName()
	if name != expected {
		t.Errorf("DisplayName '%s' does not match the expected: '%s'", name, expected)
	}

	f.PublicName = "f"
	expected = "-f"
	name = f.DisplayName()
	if name != expected {
		t.Errorf("DisplayName '%s' does not match the expected: '%s'", name, expected)
	}

	f.PublicName = "foobar"
	expected = "--foobar"
	name = f.DisplayName()
	if name != expected {
		t.Errorf("DisplayName '%s' does not match the expected: '%s'", name, expected)
	}

	f.PublicName = "f"
	expected = "f"
	name = f.Positional().DisplayName()
	if name != expected {
		t.Errorf("DisplayName '%s' does not match the expected: '%s'", name, expected)
	}

	f.PublicName = "foobar"
	expected = "foobar"
	name = f.Positional().DisplayName()
	if name != expected {
		t.Errorf("DisplayName '%s' does not match the expected: '%s'", name, expected)
	}
}

// TestFlagGetUsage tests the retrival of a flag's usage string via the GetUsage method.
func TestFlagGetUsage(t *testing.T) {
	f := NewFlag("foobar", "Activate a foobar boolean")

	if len(f.GetUsage()) <= 0 {
		t.Error("Flag's returned usage should not be empty")
	}
}

// TestFlagHelp tests that a flag's HelpText is updated to the provided value
// via the Help method.
func TestFlagHelp(t *testing.T) {
	f := Flag{}

	if f.HelpText != "" {
		t.Error("Flag HelpText should be empty string upon initialization")
	}

	expected := "this is some help text"
	f.Help(expected)

	if f.HelpText != expected {
		t.Errorf("Flag HelpText is '%v', but was expected to be: '%s'", f.HelpText, expected)
	}
}

// TestFlagMetaVar tests that a flag's MetaVarText is updated to the provided value
// via the MetaVar method.
func TestFlagMetaVar(t *testing.T) {
	f := Flag{}

	if len(f.MetaVarText) != 0 {
		t.Error("Flag MetaVarText should be empty string slice upon initialization")
	}

	expected := []string{"foo", "bar", "fizz", "buzz"}
	f.MetaVar(expected[0], expected[1:]...)

	if len(f.MetaVarText) != len(expected) {
		t.Errorf("Flag MetaVarText is '%v', but was expected to be: '%v'", f.MetaVarText, expected)
	}
}

// TestFlagNargs tests that a flag's ArgNum is updated to the provided value
// via the Nargs method.
func TestFlagNargs(t *testing.T) {
	f := Flag{}

	if len(f.ArgNum) != 0 {
		t.Error("Flag ArgNum should be empty string  upon initialization")
	}

	chars := []string{"*", "+", "?", "0", "5", "10"}
	for _, c := range chars {
		expected := c
		f.Nargs(expected)

		if f.ArgNum != expected {
			t.Errorf("Flag ArgNum is '%s', but was expected to be: '%s'", f.ArgNum, expected)
		}
	}
}

// TestFlagNotRequired tests that a flag's IsRequired boolean is updated to
// become 'false' when the Required method is called.
func TestFlagNotRequired(t *testing.T) {
	f := Flag{}

	if f.IsRequired != false {
		t.Error("Flag IsRequired should be false upon initialization")
	}

	f.IsRequired = true

	expected := false
	f.NotRequired()

	if f.IsRequired != expected {
		t.Errorf("Flag IsRequired is '%t', but was expected to be: '%t'", f.IsRequired, expected)
	}
}

// TestFlagNotPositional tests that a flag's IsPositional boolean is updated to
// become 'false' when the NotPositional method is called.
func TestFlagNotPositional(t *testing.T) {
	f := Flag{}

	if f.IsPositional != false {
		t.Error("Flag IsPositional should be false upon initialization")
	}

	f.IsPositional = true

	expected := false
	f.NotPositional()

	if f.IsPositional != expected {
		t.Errorf("Flag IsPositional is '%t', but was expected to be: '%t'", f.IsPositional, expected)
	}
}

// TestFlagPositional tests that a flag's IsPositional boolean is updated to
// become 'true' when the Positional method is called.
func TestFlagPositional(t *testing.T) {
	f := Flag{}

	if f.IsPositional != false {
		t.Error("Flag IsPositional should be false upon initialization")
	}

	expected := true
	f.Positional()

	if f.IsPositional != expected {
		t.Errorf("Flag IsPositional is '%t', but was expected to be: '%t'", f.IsPositional, expected)
	}
}

// TestFlagRequired tests that a flag's IsRequired boolean is updated to
// become 'true' when the Required method is called.
func TestFlagRequired(t *testing.T) {
	f := Flag{}

	if f.IsRequired != false {
		t.Error("Flag IsRequired should be false upon initialization")
	}

	expected := true
	f.Required()

	if f.IsRequired != expected {
		t.Errorf("Flag IsRequired is '%t', but was expected to be: '%t'", f.IsRequired, expected)
	}
}

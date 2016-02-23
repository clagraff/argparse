package argparse

import "testing"

// TestIsValidChoice ensures that when an option has no choices, or when it
// does and a valid choice is provided, the function returns true. Otherwise,
// it is expected to return false.
func TestIsValidChoice(t *testing.T) {
	expected := 3.14159

	f := NewOption("a", "a", "no choices")
	if IsValidChoice(*f, expected) != true {
		t.Error("Was expecting a true boolean as the returned value")
	}

	f = NewOption("b", "b", "includes valid choice")
	f.ValidChoices = []interface{}{"foobar", 42, false, 3.14159, nil}
	if IsValidChoice(*f, expected) != true {
		t.Error("Was expecting a true boolean as the returned value")
	}

	f = NewOption("c", "c", "does not valid choice")
	f.ValidChoices = []interface{}{"fizzbuzz", 666, true, nil}
	if IsValidChoice(*f, expected) != false {
		t.Error("Was expecting a false boolean as the returned value")
	}
}

// TestNewOption tests the creation of a new option, populated with defaults
// and appropriate name and description as provided.
func TestNewOption(t *testing.T) {
	name := "option1"
	dest := "opt dest"
	desc := "my option"
	f := NewOption(name, dest, desc)

	if f.DefaultVal != nil {
		t.Error("Default value should be nil")
	}

	if f.ConstVal != nil {
		t.Error("Constant value should be nil")
	}

	if f.DestName != dest {
		t.Error("DestName value should match expected name")
	}

	if f.PublicNames[0] != name {
		t.Error("PublicName value should match expected name")
	}

	if len(f.MetaVarText) != 1 {
		t.Error("MetaVarText does not contain only one element")
	} else if f.MetaVarText[0] != name {
		t.Error("MetaVarText[0] does not match the expected name")
	}
}

// TestOptionAction tests the Action method to ensure a option's DesiredAction
// will become the provided action.
func TestOptionAction(t *testing.T) {
	testAction := func(p *Parser, f *Option, args ...string) ([]string, error) { return nil, nil }
	f := Option{}

	f.Action(testAction)
	if f.DesiredAction == nil {
		t.Error("Option action was not properly set")
	}
}

// TestOptionChoices tests that the Choices method will set a Options avaliable choices
// to the provided []interface{}.
func TestOptionChoices(t *testing.T) {
	f := Option{}
	choices := []interface{}{"foobar", true, 12}

	if len(f.ValidChoices) != 0 {
		t.Error("Option should not contain any choices yet")
	}

	f.Choices(choices)
	if len(f.ValidChoices) != len(choices) {
		t.Errorf("Option contains %d choices, but is expected a total of %d", len(f.ValidChoices), len(choices))
	}
}

// TestOptionConst tests that a option's ConstValue is updated to the provided value
// via the Const method.
func TestOptionConst(t *testing.T) {
	f := Option{}

	if f.ConstVal != nil {
		t.Error("Option ConstVal should be nil upon initialization")
	}

	expected := "some value"
	f.Const(expected)

	if f.ConstVal != expected {
		t.Errorf("Option ConstVal is '%v', but was expected to be: '%s'", f.ConstVal, expected)
	}
}

// TestOptionDefault tests that a option's DefaultValue is updated to the provided value
// via the Default method.
func TestOptionDefault(t *testing.T) {
	f := Option{}

	if f.DefaultVal != nil {
		t.Error("Option DefaultVal should be nil upon initialization")
	}

	expected := "some value"
	f.Default(expected)

	if f.DefaultVal != expected {
		t.Errorf("Option DefaultVal is '%v', but was expected to be: '%s'", f.DefaultVal, expected)
	}
}

// TestOptionDisplayName tests the retrival of a option's display name, with an
// appropriate number of preceeding hypens, via the DisplayName method.
func TestOptionDisplayName(t *testing.T) {
	f := Option{}

	expected := ""
	name := f.DisplayName()
	if name != expected {
		t.Errorf("DisplayName '%s' does not match the expected: '%s'", name, expected)
	}

	f.PublicNames = []string{"f"}
	expected = "-f"
	name = f.DisplayName()
	if name != expected {
		t.Errorf("DisplayName '%s' does not match the expected: '%s'", name, expected)
	}

	f.PublicNames = []string{"foobar"}
	expected = "--foobar"
	name = f.DisplayName()
	if name != expected {
		t.Errorf("DisplayName '%s' does not match the expected: '%s'", name, expected)
	}

	f.PublicNames = []string{"f"}
	expected = "f"
	name = f.Positional().DisplayName()
	if name != expected {
		t.Errorf("DisplayName '%s' does not match the expected: '%s'", name, expected)
	}

	f.PublicNames = []string{"foobar"}
	expected = "foobar"
	name = f.Positional().DisplayName()
	if name != expected {
		t.Errorf("DisplayName '%s' does not match the expected: '%s'", name, expected)
	}
}

// TestOptionGetUsage tests the retrival of a option's usage string via the GetUsage method.
func TestOptionGetUsage(t *testing.T) {
	f := NewOption("foobar", "foobar dest", "Activate a foobar boolean")

	if len(f.GetUsage()) <= 0 {
		t.Error("Option's returned usage should not be empty")
	}
}

// TestOptionHelp tests that a option's HelpText is updated to the provided value
// via the Help method.
func TestOptionHelp(t *testing.T) {
	f := Option{}

	if f.HelpText != "" {
		t.Error("Option HelpText should be empty string upon initialization")
	}

	expected := "this is some help text"
	f.Help(expected)

	if f.HelpText != expected {
		t.Errorf("Option HelpText is '%v', but was expected to be: '%s'", f.HelpText, expected)
	}
}

// TestOptionMetaVar tests that a option's MetaVarText is updated to the provided value
// via the MetaVar method.
func TestOptionMetaVar(t *testing.T) {
	f := Option{}

	if len(f.MetaVarText) != 0 {
		t.Error("Option MetaVarText should be empty string slice upon initialization")
	}

	expected := []string{"foo", "bar", "fizz", "buzz"}
	f.MetaVar(expected[0], expected[1:]...)

	if len(f.MetaVarText) != len(expected) {
		t.Errorf("Option MetaVarText is '%v', but was expected to be: '%v'", f.MetaVarText, expected)
	}
}

// TestOptionNargs tests that a option's ArgNum is updated to the provided value
// via the Nargs method.
func TestOptionNargs(t *testing.T) {
	f := Option{}

	if len(f.ArgNum) != 0 {
		t.Error("Option ArgNum should be empty string  upon initialization")
	}

	chars := []string{"*", "+", "?", "0", "5", "10"}
	for _, c := range chars {
		expected := c
		f.Nargs(expected)

		if f.ArgNum != expected {
			t.Errorf("Option ArgNum is '%s', but was expected to be: '%s'", f.ArgNum, expected)
		}
	}
}

// TestOptionNotRequired tests that a option's IsRequired boolean is updated to
// become 'false' when the Required method is called.
func TestOptionNotRequired(t *testing.T) {
	f := Option{}

	if f.IsRequired != false {
		t.Error("Option IsRequired should be false upon initialization")
	}

	f.IsRequired = true

	expected := false
	f.NotRequired()

	if f.IsRequired != expected {
		t.Errorf("Option IsRequired is '%t', but was expected to be: '%t'", f.IsRequired, expected)
	}
}

// TestOptionNotPositional tests that a option's IsPositional boolean is updated to
// become 'false' when the NotPositional method is called.
func TestOptionNotPositional(t *testing.T) {
	f := Option{}

	if f.IsPositional != false {
		t.Error("Option IsPositional should be false upon initialization")
	}

	f.IsPositional = true

	expected := false
	f.NotPositional()

	if f.IsPositional != expected {
		t.Errorf("Option IsPositional is '%t', but was expected to be: '%t'", f.IsPositional, expected)
	}
}

// TestOptionPositional tests that a option's IsPositional boolean is updated to
// become 'true' when the Positional method is called.
func TestOptionPositional(t *testing.T) {
	f := Option{}

	if f.IsPositional != false {
		t.Error("Option IsPositional should be false upon initialization")
	}

	expected := true
	f.Positional()

	if f.IsPositional != expected {
		t.Errorf("Option IsPositional is '%t', but was expected to be: '%t'", f.IsPositional, expected)
	}
}

// TestOptionRequired tests that a option's IsRequired boolean is updated to
// become 'true' when the Required method is called.
func TestOptionRequired(t *testing.T) {
	f := Option{}

	if f.IsRequired != false {
		t.Error("Option IsRequired should be false upon initialization")
	}

	expected := true
	f.Required()

	if f.IsRequired != expected {
		t.Errorf("Option IsRequired is '%t', but was expected to be: '%t'", f.IsRequired, expected)
	}
}

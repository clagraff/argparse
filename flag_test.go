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

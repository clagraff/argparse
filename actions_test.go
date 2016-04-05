package argparse

import "testing"

// TestStore_OneNargs tests the Store Action will store the expected value and
// return the appropriate args & error when operating upon a option with
// one expected argument.
func TestStore_OneNargs(t *testing.T) {
	p := NewParser("parser")
	f := NewOption("option", "option", "option").Nargs("1")
	args := []string{"foobar"}

	args, err := Store(p, f, args...)

	if len(args) != 0 {
		t.Error("args should be empty")
	}

	if err != nil {
		t.Error("An error was not expected")
	}

	if p.Namespace.Mapping[f.DestName] != "foobar" {
		t.Error("Action did not store correct value in parser")
	}

	args = []string{}
	_, err = Store(p, f, args...)
	if err == nil {
		t.Error("An error was expected but did not occurr")
	}
}

// TestStore_ThreeNargs tests the Store Action will store the expected value and
// return the appropriate args & error when operating upon a option with
// three  expected arguments.
func TestStore_ThreeNargs(t *testing.T) {
	p := NewParser("parser")
	f := NewOption("option", "option", "option").Nargs("3")
	args := []string{"foo", "bar", "fizzbuzz"}

	args, err := Store(p, f, args...)

	if len(args) != 0 {
		t.Error("args should be empty")
	}

	if err != nil {
		t.Error("An error was not expected")
	}

	if len(p.Namespace.Mapping[f.DestName].([]string)) != 3 {
		t.Error("Action did not store correct number of values in parser")
	}

	args = []string{}
	_, err = Store(p, f, args...)
	if err == nil {
		t.Error("An error was expected but did not occurr")
	}
}

// TestStore_AnyNargs tests the Store Action will store the expected value and
// return the appropriate args & error when operating upon a option with
// any number of expected arguments.
func TestStore_AnyNargs(t *testing.T) {
	p := NewParser("parser")
	f := NewOption("option", "option", "option").Nargs("*")
	args := []string{"foo", "bar", "fizz", "buzz", "hello", "world"}

	args, err := Store(p, f, args...)

	if len(args) != 0 {
		t.Error("args should be empty")
	}

	if err != nil {
		t.Error("An error was not expected")
	}

	if len(p.Namespace.Mapping[f.DestName].([]string)) != 6 {
		t.Error("Action did not store correct number of values in parser")
	}

	args = []string{}
	_, err = Store(p, f, args...)
	if err != nil {
		t.Error("An error was not expected")
	}
}

// TestStore_LeastOneNargs tests the Store Action will store the expected value and
// return the appropriate args & error when operating upon a option with
// at least one number of expected arguments.
func TestStore_LeastOneNargs(t *testing.T) {
	p := NewParser("parser")
	f := NewOption("option", "option", "option").Nargs("+")
	args := []string{"foo", "fizz", "world"}

	args, err := Store(p, f, args...)

	if len(args) != 0 {
		t.Error("args should be empty")
	}

	if err != nil {
		t.Error("An error was not expected")
	}

	if len(p.Namespace.Mapping[f.DestName].([]string)) != 3 {
		t.Error("Action did not store correct number of values in parser")
	}

	args = []string{}
	_, err = Store(p, f, args...)
	if err == nil {
		t.Error("An error was expected but not returned")
	}
}

// TestStoreConst tests the StoreConst Action will store a option's ConstValue.
func TestStoreConst(t *testing.T) {
	p := NewParser("parser")
	f := NewOption("option", "option", "option").Nargs("0").Const("hello world")
	args := []string{}

	args, err := StoreConst(p, f, args...)

	if len(args) != 0 {
		t.Error("args should be empty")
	}

	if err != nil {
		t.Error("An error was not expected")
	}

	if p.Namespace.Mapping[f.DestName] != f.ConstVal {
		t.Error("Action did not store the correct ConstValue in the parser")
	}
}

// TestStoreFalse tests the StoreFalse Action will store a false boolean.
func TestStoreFalse(t *testing.T) {
	p := NewParser("parser")
	f := NewOption("option", "option", "option").Nargs("0")
	args := []string{}

	args, err := StoreFalse(p, f, args...)

	if len(args) != 0 {
		t.Error("args should be empty")
	}

	if err != nil {
		t.Error("An error was not expected")
	}

	if p.Namespace.Mapping[f.DestName] != "false" {
		t.Error("Action did not store the correct boolean value in the parser")
	}
}

// TestStoreTrue tests the StoreFalse Action will store a true boolean.
func TestStoreTrue(t *testing.T) {
	p := NewParser("parser")
	f := NewOption("option", "option", "option").Nargs("0")
	args := []string{}

	args, err := StoreTrue(p, f, args...)

	if len(args) != 0 {
		t.Error("args should be empty")
	}

	if err != nil {
		t.Error("An error was not expected")
	}

	if p.Namespace.Mapping[f.DestName] != "true" {
		t.Error("Action did not store the correct boolean value in the parser")
	}
}

// TestAppend_OneNargs tests the Append Action will store the expected value and
// return the appropriate args & error when operating upon a option with
// one expected argument.
func TestAppend_OneNargs(t *testing.T) {
	p := NewParser("parser")
	f := NewOption("option", "option", "option").Nargs("1")
	args := []string{"foobar"}

	args, err := Append(p, f, args...)

	if len(args) != 0 {
		t.Error("args should be empty")
	}

	if err != nil {
		t.Error("An error was not expected")
	}

	if len(p.Namespace.Mapping[f.DestName].([]string)) != 1 {
		t.Error("Action did not store correct value in parser")
	}

	args = []string{}
	_, err = Append(p, f, args...)
	if err == nil {
		t.Error("An error was expected but did not occurr")
	}
}

// TestAppend_ThreeNargs tests the Store Action will store the expected value and
// return the appropriate args & error when operating upon a option with
// three  expected arguments.
func TestAppend_ThreeNargs(t *testing.T) {
	p := NewParser("parser")
	f := NewOption("option", "option", "option").Nargs("3")
	args := []string{"foo", "bar", "fizzbuzz"}

	args, err := Append(p, f, args...)

	if len(args) != 0 {
		t.Error("args should be empty")
	}

	if err != nil {
		t.Error("An error was not expected")
	}

	if len(p.Namespace.Mapping[f.DestName].([]string)) != 3 {
		t.Error("Action did not store correct number of values in parser")
	}

	args = []string{}
	_, err = Append(p, f, args...)
	if err == nil {
		t.Error("An error was expected but did not occurr")
	}
}

// TestStore_AnyNargs tests the Store Action will store the expected value and
// return the appropriate args & error when operating upon a option with
// any number of expected arguments.
func TestAppend_AnyNargs(t *testing.T) {
	p := NewParser("parser")
	f := NewOption("option", "option", "option").Nargs("*")
	args := []string{"foo", "bar", "fizz", "buzz", "hello", "world"}

	args, err := Append(p, f, args...)

	if len(args) != 0 {
		t.Error("args should be empty")
	}

	if err != nil {
		t.Error("An error was not expected")
	}

	if len(p.Namespace.Mapping[f.DestName].([]string)) != 6 {
		t.Error("Action did not store correct number of values in parser")
	}

	args = []string{}
	_, err = Append(p, f, args...)
	if err != nil {
		t.Error("An error was not expected")
	}
}

// TestStore_LeastOneNargs tests the Store Action will store the expected value and
// return the appropriate args & error when operating upon a option with
// at least one number of expected arguments.
func TestAppend_LeastOneNargs(t *testing.T) {
	p := NewParser("parser")
	f := NewOption("option", "option", "option").Nargs("+")
	args := []string{"foo", "fizz", "world"}

	args, err := Append(p, f, args...)

	if len(args) != 0 {
		t.Error("args should be empty")
	}

	if err != nil {
		t.Error("An error was not expected")
	}

	if len(p.Namespace.Mapping[f.DestName].([]string)) != 3 {
		t.Error("Action did not store correct number of values in parser")
	}

	args = []string{}
	_, err = Append(p, f, args...)
	if err == nil {
		t.Error("An error was expected but not returned")
	}
}

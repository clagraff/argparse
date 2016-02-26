package argparse

import "testing"

// TestInvalidChoiceErr asserts that the returned error is not nil and contains
// an error message.
func TestInvalidChoiceErr(t *testing.T) {
	err := InvalidChoiceErr(Option{}, "test")
	if err == nil {
		t.Error("Error cannot be nil")
	}
	if len(err.Error()) == 0 {
		t.Error("Error message cannot be empty")
	}
}

// TestInvalidTypeErr asserts that the returned error is not nil and contains
// an error message.
func TestInvalidTypeErr(t *testing.T) {
	err := InvalidTypeErr(Option{}, "test")
	if err == nil {
		t.Error("Error cannot be nil")
	}
	if len(err.Error()) == 0 {
		t.Error("Error message cannot be empty")
	}
}

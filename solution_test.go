package solution

import (
	"testing"
)

func TestGetMessage(t *testing.T) {
	correct := "Hello 🗺️!"
	emoji := GetMessage()
	if emoji != correct {
		t.Errorf("Sum was incorrect, got: %v, want: %v.", emoji, correct)
	}
}

package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Umakant"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Umakant")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Umakant") = %q, %v, want match for %#q`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

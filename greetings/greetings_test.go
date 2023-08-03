package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	greeting := Hello(name)
	if !want.MatchString(greeting.Message) || greeting.Err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, greeting.Message, greeting.Err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	greeting := Hello("")
	if greeting.Message != "" || greeting.Err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, greeting.Message, greeting.Err)
	}
}

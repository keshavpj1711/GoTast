package main

import (
	"testing"
)

func TestRepeat(t *testing.T)  {
	repeat_output := Repeat("a")
	expected := "aaaa"

	if repeat_output != expected {
		t.Errorf("Got %q expected %q", repeat_output, expected)
	}
}
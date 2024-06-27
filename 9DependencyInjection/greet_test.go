package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Keshav")

	// The Buffer type from the bytes package implements the Writer interface,
	// because it has the method
	// Write(p []byte) (n int, err error).

	got := buffer.String()
	want := "Hello Keshav"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

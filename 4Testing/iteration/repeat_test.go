package main

import (
	"fmt"
	"strings"
	"testing"
)

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 4)
	}
}

func ExampleRepeat()  {
	repeat := Repeat("k", 4)
	fmt.Println(repeat)
	// Output: kkkk
}

func TestRepeat(t *testing.T)  {
	repeat_output := Repeat("a", 5)
	expected := strings.Repeat("a", 5)

	if repeat_output != expected {
		t.Errorf("Got %q expected %q", repeat_output, expected)
	}
}
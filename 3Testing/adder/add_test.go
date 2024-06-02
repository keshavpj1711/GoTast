package main

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestAdd(t *testing.T)  {
	got := Add(3, 4)
	want := 7

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
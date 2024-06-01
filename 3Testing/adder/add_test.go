package main

import "testing"

func TestAdd(t *testing.T)  {
	got := Add(3, 4)
	want := 7

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}